package v1

type AccountCreditingRequest struct {
	Amount uint `json:"amount"`
	Description string `json:"description"`
}
