package schemas

type RefCodeResponse struct {
	Code    string `json:"referral_code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type RefCodeUsedRequest struct {
	ReferralCode string  `json:"referral_code"`
	Id           int64   `json:"id"`
	Domain       string  `json:"domain"`
	Price        float64 `json:"price"`
	Address      string  `json:"address"`
}

type RefCodeRequest struct {
	Domain  string `json:"domain"`
	Address string `json:"address"`
}

type RefCodeCounterRequest struct {
	Counter int64 `json:"counter"`
}

type TrackingResponse struct {
	ReferralCode      string  `json:"referral_code"`
	Id                int64   `json:"id"`
	Count             int64   `json:"count"`
	Level             string  `json:"level"`
	Rate              float64 `json:"rate"`
	WithdrawAvailable float64 `json:"withdraw_available"`
}
