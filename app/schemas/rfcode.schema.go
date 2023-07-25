package schemas

type RefCodeResponse struct {
	Code    int64 `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type RefCodeUsedRequest struct {
	RefCode int64  `json:"refcode"`
	Domain  string  `json:"domain"`
	Price   float64 `json:"price"`
	Address string  `json:"address"`
}

type RefCodeRequest struct {
	RefCode int64 `json:"ref_code"`
	Address string `json:"address"`
}

type RefCodeCounterRequest struct {
	Counter int64 `json:"counter"`
}
