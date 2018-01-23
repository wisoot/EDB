package v1

type AccountDebitingRequest struct {
	Amount uint `json:"amount"`
	Description string `json:"description"`
}
