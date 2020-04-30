package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	logger "github.com/fravega/go-logger/v2"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)
)

type SupportAPI struct {
	logger logger.Logger
	//	sellerRepository   *sourceInfra.SellerRepository
	/jwtDecoder         *JwtDecoder
	restrictAccessByTC bool
}

func CreateSupportAPI(jwtDecoder *JwtDecoder,
	restrictAccessByTC bool,
) *SupportAPI {
	return &SupportAPI{
		logger:             logger.GetDefaultLogger(),
		jwtDecoder:         jwtDecoder,
		restrictAccessByTC: restrictAccessByTC,
	}
}

func (s *SupportAPI) readBody(body interface{}, r *rest.Request) error {
	read, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error while reading body")
		return err
	}
	unmarshalErr := json.Unmarshal(read, &body)
	if unmarshalErr != nil {
		s.logger.WithFields(logger.Fields{"error": unmarshalErr}).Error("Error while unmarshal body")
		return unmarshalErr
	}
	return nil
}

func (s *SupportAPI) getUser(r *rest.Request) (*model.User, error) {
	user, err := s.getUserWithConditionalTermsVerification(r, true)
	if err != nil {
		return nil, err
	}
	if hasPermission(user.GetPermissions(), PermissionTermsSign) {
		return nil, &BadRequestError{code: termsSingPending, message: "Terminos y condiciones no firmados"}
	}

	return user, err

}

func (s *SupportAPI) getUserWithConditionalTermsVerification(r *rest.Request, checkTerms bool) (*model.User, error) {
	token := r.Header.Get("X-FVG-TOKEN-CORS")
	if token == "" {
		return nil, &BadRequestError{code: tokenRequired, message: "Token is required"}
	}

	username, sellerId, permissions, err := s.jwtDecoder.GetTokenAttributes(token)
	if err != nil {
		return nil, &BadRequestError{code: invalidToken, message: err.Error()}
	}
	impersonatedSellerId := r.Header.Get("X-FVG-SELLER-ID")
	if impersonatedSellerId != "" && hasPermission(permissions, PermissionImpersonate) {
		sellerId = impersonatedSellerId
		logger.GetDefaultLogger().Debugf("Seller impersonated to %v by %f", sellerId, username)
	}

	logger.GetDefaultLogger().Debugf("User %v", username)

	return model.NewUser(username, token, permissions), nil
}

func hasPermission(permissions []string, permission string) bool {
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func (s *SupportAPI) HasPermissions(u *model.User, permissions []string) bool {
	has := true
	for _, p := range permissions {
		has = has && hasPermission(u.GetPermissions(), p)
	}
	return has
}

func (s *SupportAPI) getPage(r *rest.Request, limit bool) (int, int) {
	pageParam := r.URL.Query().Get("page")
	sizeParam := r.URL.Query().Get("size")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 0 {
		page = 0
	}
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size < 0 || size > 500 {
		if limit {
			size = 500
		} else {

			if size < 0 {
				size = -size
			}
		}
	}
	return page, size
}

type BadRequestError struct {
	code    string
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

const (
	tokenRequired          = "token_required"
	sellerIdRequired       = "seller_id_required"
	invalidToken           = "invalid_token"
	validationCode         = "validation"
	itemAlreadyExistsCode  = "item_already_exists"
	categoryNotFoundCode   = "category_not_found"
	brandNotFoundCode      = "brand_not_found"
	jobNotFoundCode        = "job_not_found"
	logisticsNotFoundCode  = "logistics_invalid"
	internalServerCode     = "internal_server_error"
	unauthorrizedCode      = "unauthorized"
	termsSingPending       = "terms_sing_pending"
	termsSingRejected      = "terms_sing_rejected"
	termsSingAlreadySigned = "terms_already_signed"

	PermissionDefault       = "product.list"
	PermissionItemAdd       = "product.add"
	PermissionItemUpdate    = "product.update"
	PermissionPricingUpdate = "pricing.update"
	PermissionStockUpdate   = "stock.update"
	PermissionImpersonate   = "impersonate.all"
	PermissionTerms         = "terms.view"
	PermissionTermsSign     = "terms.write"
)

func (s *SupportAPI) writeError(err error, w rest.ResponseWriter) {
	switch err.(type) {
	case *BadRequestError:
		badRequest := err.(*BadRequestError)
		failed(http.StatusBadRequest, badRequest.code, err, w)
		break
	case *model.ValidationsError:
		writeValidationsErrors(err.(*model.ValidationsError), w)
		break
	case *model.ItemAlreadyExistsError:
		failed(http.StatusConflict, itemAlreadyExistsCode, err, w)
		break
	case *model.CategoryNotFoundError:
		failed(http.StatusNotFound, categoryNotFoundCode, err, w)
		break
	case *model.BrandNotFoundError:
		failed(http.StatusNotFound, brandNotFoundCode, err, w)
		break
	case *model.JobNotFoundError:
		failed(http.StatusNotFound, jobNotFoundCode, err, w)
		break
	case *model.LogisticsNotFoundError:
		failed(http.StatusConflict, logisticsNotFoundCode, err, w)
	case *model.ForbiddenError:
		failed(http.StatusUnauthorized, unauthorrizedCode, err, w)
		break
	case *model.SignAlreadyExistsError:
		failed(http.StatusConflict, termsSingAlreadySigned, err, w)
		break
	default:
		failed(http.StatusInternalServerError, internalServerCode, err, w)
		break
	}
}

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type errorResponse struct {
	Code   string      `json:"code"`
	Failed interface{} `json:"error"`
}

func writeValidationsErrors(v *model.ValidationsError, w rest.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	validations := make([]validationError, 0)
	for k, v := range v.Validations() {
		validations = append(validations, validationError{
			Field:   k,
			Message: v,
		})
	}
	response := errorResponse{
		Code:   validationCode,
		Failed: validations,
	}
	w.WriteJson(response)
}

func failed(responseCode int, code string, v error, w rest.ResponseWriter) {
	w.WriteHeader(responseCode)
	response := errorResponse{
		Code:   code,
		Failed: v.Error(),
	}
	w.WriteJson(response)
}
