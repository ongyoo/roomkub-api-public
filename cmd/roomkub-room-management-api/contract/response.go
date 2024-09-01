package roomContract

type BaseRoomContractItemResponse[T any] struct {
	Customer T `json:"contract_item"`
}

type BaseRoomContractItemsResponse[T any] struct {
	ContractList T `json:"contract_list"`
}

type CreateRoomContractResponse struct {
	ContractID *string `json:"contract_id"`
}

// type RoomResponse struct {
// 	ID                 string            `json:"id"`
// }
