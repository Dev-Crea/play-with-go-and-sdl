package models

type ContractResponse struct{}

type ContractResponseMeta struct {
	Total int32
	Page  int32
	Limit int32
}
type ContractResponseData struct {
	// Data Contract
	Meta ContractResponseMeta
}
