package module

import (
	"gitlab.com/fravega-it/adn/ipos/medios-de-pago/backend-medios-de-pago/domain/actions"
	"gitlab.com/fravega-it/adn/ipos/medios-de-pago/backend-medios-de-pago/infrastructure/api"
	"gitlab.com/fravega-it/adn/ipos/medios-de-pago/backend-medios-de-pago/infrastructure/server"
	"gitlab.com/fravega-it/adn/ipos/medios-de-pago/backend-medios-de-pago/infrastructure/server/controllers"
)

func MakeControllers(
	newBin actions.SetBin,
	deleteBin actions.DeleteBin,
	updateBin actions.UpdateBinById,
	getAllBins actions.GetAllBins,
	findBinById actions.GetBinById,
	findBinByNumber actions.GetBinByNumber,

	newPaymentMethodBin actions.SetPaymentMethodBin,
	deletePaymentMethodBin actions.DeletePaymentMethodBin,
	updatePaymentMethodBin actions.UpdatePaymentMethodBinById,
	getAllPaymentMethodBin actions.GetAllPaymentMethodBins,
	findPaymentMethodBinById actions.GetPaymentMethodBinById,

	newIssuingBankBin actions.SetIssuingBankBin,
	updateIssuingBankBin actions.UpdateIssuingBankBinById,
	deleteIssuingBankBin actions.DeleteIssuingBankBin,
	getAllIssuingBankBin actions.GetAllIssuingBankBins,
	findIssuingBankBubById actions.GetIssuingBankBinById,

	newPaymentMethod actions.SetPaymentMethod,
	updatePaymentMethod actions.UpdatePaymentMethodById,
	deletePaymentMethod actions.DeletePaymentMethod,
	getAllPaymentMethod actions.GetAllPaymentMethod,
	findPaymentMethodById actions.GetPaymentMethodById,

	newIssuingBank actions.SetIssuingBank,
	updateIssuingBank actions.UpdateIssuingBankById,
	deleteIssuingBank actions.DeleteIssuingBank,
	getAllIssuingBank actions.GetAllIssuingBank,
	findIssuingBankById actions.GetIssuingBankById,

	/*	sellerRepository *sourceInfra.SellerRepository,
		jwtDecoder *action.JwtDecoder,
		newItem actions.NewItem,
		updateItem actions.UpdateItem,
		updatePrice actions.UpdatePrice,
		updateStock actions.UpdateStock,
		findById actions.GetItemById,
		findBySellerId actions.FindItems,
		newJob actions.NewJob,
		getJobs actions.GetAllJobs,
		getJob actions.GetJob,
		getTasks actions.GetAllTasks,
		getSellerById actions.GetSellerById,
		getAllBrands actions.GetAllBrands,
		activateItem actions.ActivateItem,
		requestApprovalItem actions.RequestApprovalItem,
		getAllSellers actions.GetAllSellers,
		termsVerifier actions.CheckTermsSigned,
		termsSigner actions.SignTerms,
		restrictAccessByTC bool,*/
) []server.Controller {
	/*	supportAPI := api.CreateSupportAPI(sellerRepository, jwtDecoder, termsVerifier, restrictAccessByTC)
		// Create Controllers
		itemController := controllers.CreateItemController(createItemAPI(supportAPI, newItem, updateItem, updatePrice, updateStock, findById, findBySellerId, activateItem, requestApprovalItem))
		taskController := controllers.CreateJobController(createJobAPI(supportAPI, newJob, getJobs, getJob, getTasks))
		sellerController := controllers.CreateSellerController(createSellerAPI(supportAPI, getSellerById, getAllBrands, getAllSellers))
		termsController := controllers.CreateTermsController(createTermsAPI(supportAPI, termsVerifier, termsSigner))*/

	binController := controllers.CreateBinController(createBinAPI(newBin, deleteBin, updateBin, getAllBins, findBinById, findBinByNumber))
	paymentMethodBinController := controllers.CreatePaymentMethodBinController(createPaymentMethodBinAPI(newPaymentMethodBin, deletePaymentMethodBin, updatePaymentMethodBin, getAllPaymentMethodBin, findPaymentMethodBinById))
	issuingBankBinController := controllers.CreateIssuingBankBinController(createIssuingBankBinAPI(newIssuingBankBin, deleteIssuingBankBin, updateIssuingBankBin, getAllIssuingBankBin, findIssuingBankBubById))
	paymentMethodController := controllers.CreatePaymentMethodController(createPaymentMethodAPI(newPaymentMethod, deletePaymentMethod, updatePaymentMethod, getAllPaymentMethod, findPaymentMethodById))
	issuingBankController := controllers.CreateIssuingBankController(createIssuingBankAPI(newIssuingBank, deleteIssuingBank, updateIssuingBank, getAllIssuingBank, findIssuingBankById))

	return []server.Controller{binController, paymentMethodBinController, issuingBankBinController, paymentMethodController, issuingBankController}
}

func createBinAPI(newBin actions.SetBin, deleteBin actions.DeleteBin, updateBin actions.UpdateBinById, getAllBins actions.GetAllBins, findBinById actions.GetBinById, findBinByNumber actions.GetBinByNumber) *api.BinAPI {
	return api.CreateBinAPI(newBin, deleteBin, updateBin, getAllBins, findBinById, findBinByNumber)
}

func createPaymentMethodBinAPI(newPaymentMethodBin actions.SetPaymentMethodBin, deletePaymentMethodBin actions.DeletePaymentMethodBin, updatePaymentMethodBin actions.UpdatePaymentMethodBinById, getAllPaymentMethodBin actions.GetAllPaymentMethodBins, findPaymentMethodBinById actions.GetPaymentMethodBinById) *api.PaymentMethodBinAPI {
	return api.CreatePaymentMethodBinAPI(newPaymentMethodBin, deletePaymentMethodBin, updatePaymentMethodBin, getAllPaymentMethodBin, findPaymentMethodBinById)
}

func createIssuingBankBinAPI(newIssuingBankBin actions.SetIssuingBankBin, deleteIssuingBankBin actions.DeleteIssuingBankBin, updateIssuingBankBin actions.UpdateIssuingBankBinById, getAllIssuingBankBin actions.GetAllIssuingBankBins, findIssuingBankBubById actions.GetIssuingBankBinById) *api.IssuingBankBinAPI {
	return api.CreateIssuingBankBinAPI(newIssuingBankBin, deleteIssuingBankBin, updateIssuingBankBin, getAllIssuingBankBin, findIssuingBankBubById)
}

func createPaymentMethodAPI(newPaymentMethod actions.SetPaymentMethod, deletePaymentMethod actions.DeletePaymentMethod, updatePaymentMethod actions.UpdatePaymentMethodById, getAllPaymentMethod actions.GetAllPaymentMethod, findPaymentMethodById actions.GetPaymentMethodById) *api.PaymentMethodAPI {
	return api.CreatePaymentMethodAPI(newPaymentMethod, deletePaymentMethod, updatePaymentMethod, getAllPaymentMethod, findPaymentMethodById)
}

func createIssuingBankAPI(newIssuingBank actions.SetIssuingBank, deleteIssuingBank actions.DeleteIssuingBank, updateIssuingBank actions.UpdateIssuingBankById, getAllIssuingBank actions.GetAllIssuingBank, findIssuingBankById actions.GetIssuingBankById) *api.IssuingBankAPI {
	return api.CreateIssuingBankAPI(newIssuingBank, deleteIssuingBank, updateIssuingBank, getAllIssuingBank, findIssuingBankById)
}

/*
func createItemAPI(support *api.SupportAPI, newItem actions.NewItem, updateItem actions.UpdateItem, updatePrice actions.UpdatePrice, updateStock actions.UpdateStock,
	findById actions.GetItemById, findBySellerId actions.FindItems, activateItem actions.ActivateItem, requestApprovalItem actions.RequestApprovalItem) *api.ItemAPI {
	return api.CreateItemAPI(support, newItem, updateItem, updatePrice, updateStock, findById, findBySellerId, activateItem, requestApprovalItem)
}

func createJobAPI(support *api.SupportAPI, newJob actions.NewJob, getJobs actions.GetAllJobs, getJob actions.GetJob, getTasks actions.GetAllTasks) *api.JobAPI {
	return api.CreateJobAPI(support, newJob, getJobs, getJob, getTasks)
}

func createSellerAPI(support *api.SupportAPI, getSellerById actions.GetSellerById, getAllBrands actions.GetAllBrands, getAllSellers actions.GetAllSellers) *api.SellerAPI {
	return api.CreateSellerAPI(support, getSellerById, getAllBrands, getAllSellers)
}

func createTermsAPI(support *api.SupportAPI, termsVerifier actions.CheckTermsSigned, termsSigner actions.SignTerms) *api.TermsAPI {
	return api.CreateTermsAPI(support, termsVerifier, termsSigner)
}
*/
