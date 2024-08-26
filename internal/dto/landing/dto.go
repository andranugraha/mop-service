package landing

type LandingRequest struct {
	Code string `json:"code" validate:"required"`
}

type LandingResponseData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
