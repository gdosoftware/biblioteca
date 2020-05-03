package model

import "fmt"

type ItemAlreadyExistsError struct {
	sellerID string
	refID    string
}

func (e *ItemAlreadyExistsError) Error() string {
	return fmt.Sprintf("Item existente con refId: [%v] and sellerId: [%v]", e.refID, e.sellerID)
}

func NewItemAlreadyExistsError(sellerID string, refID string) *ItemAlreadyExistsError {
	return &ItemAlreadyExistsError{sellerID: sellerID, refID: refID}
}

type ValidationsError struct {
	validations map[string]string
}

func (e *ValidationsError) Error() string {
	return fmt.Sprintf("Tiene [%v] validaciones fallidas", len(e.validations))
}

func (e *ValidationsError) Validations() map[string]string {
	return e.validations
}

func NewValidationsError(validations map[string]string) *ValidationsError {
	return &ValidationsError{validations: validations}
}

type CategoryNotFoundError struct {
	categoryId string
}

func NewCategoryNotFoundError(categoryId string) *CategoryNotFoundError {
	return &CategoryNotFoundError{categoryId: categoryId}
}

func (e *CategoryNotFoundError) Error() string {
	return fmt.Sprintf("Categoria no encontrada [%v]", e.categoryId)
}

type BrandNotFoundError struct {
	brandId string
}

func NewBrandNotFoundError(brandId string) *BrandNotFoundError {
	return &BrandNotFoundError{brandId: brandId}
}

func (e *BrandNotFoundError) Error() string {
	return fmt.Sprintf("Marca no encontrada [%v]", e.brandId)
}

type ForbiddenError struct {
	message string
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{message: message}
}

func (e *ForbiddenError) Error() string {
	return e.message
}

type ItemNotFoundError struct {
	itemId string
}

func NewItemNotFoundError(itemId string) *ItemNotFoundError {
	return &ItemNotFoundError{itemId: itemId}
}

func (e *ItemNotFoundError) Error() string {
	return fmt.Sprintf("Item no encontrado [%v]", e.itemId)
}

type JobNotFoundError struct {
	jobId string
}

func NewJobNotFoundError(jobId string) *JobNotFoundError {
	return &JobNotFoundError{jobId: jobId}
}

func (e *JobNotFoundError) Error() string {
	return fmt.Sprintf("Job no encontrado [%v]", e.jobId)
}

type LogisticsNotFoundError struct {
	message string
}

func NewLogisticsNotFoundError(message string) *LogisticsNotFoundError {
	return &LogisticsNotFoundError{message: message}
}

func (e *LogisticsNotFoundError) Error() string {
	return e.message
}

func NewSignAlreadyExistsError(sellerID string, termsID string) *SignAlreadyExistsError {
	return &SignAlreadyExistsError{sellerID: sellerID, termsID: termsID}
}

type SignAlreadyExistsError struct {
	sellerID string
	termsID  string
}

func (SignAlreadyExistsError) Error() string {
	return "Contrato ya firmado"
}
