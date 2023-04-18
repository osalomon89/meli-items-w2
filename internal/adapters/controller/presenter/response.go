package presenter

type ItemResponse struct {
	Error bool     `json:"error"`
	Data  ItemJson `json:"data"`
}

type ItemsResponse struct {
	Error bool       `json:"error"`
	Data  []ItemJson `json:"data"`
}

type ItemDeletedResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
