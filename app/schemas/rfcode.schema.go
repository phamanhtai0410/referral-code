package schemas

type RefCodeResponse struct {
	Id      int64  `json:"id"`
	Code    string `json:"referral_code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type RefCodeUsedRequest struct {
	ReferralCode string  `json:"referral_code" validate:"required"`
	Domain       string  `json:"domain" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Address      string  `json:"address" validate:"required"`
}

type RefCodeRequest struct {
	Domain  string `json:"domain"`
	Address string `json:"address"`
}

type RefCodeCounterRequest struct {
	Counter int64 `json:"counter"`
}

type TrackingResponse struct {
	ReferralCode string  `json:"referral_code"`
	Count        int64   `json:"count"`
	Level        string  `json:"level"`
	Rate         float64 `json:"rate"`
	TotalEarn    float64 `json:"total_earn"`
}
