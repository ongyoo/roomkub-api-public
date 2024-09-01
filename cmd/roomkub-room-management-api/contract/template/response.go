package contractTemplate

import "time"

type BaseContractTemplateItemResponse[T any] struct {
	Item T `json:"contract_template_item"`
}

type BaseContractTemplateItemsResponse[T any] struct {
	ContractList T `json:"contract_template_list"`
}

type CreateContractTemplateResponse struct {
	ContractID *string `json:"contract_template_id"`
}

type ContractTemplateResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	IsPublisher bool      `json:"is_publisher"` // สามารถให้คนอื่นใช้ฟอร์มได้
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
