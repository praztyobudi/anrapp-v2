package dto

type LoginResponse struct {
	Message      string                 `json:"message"`
	Token        string                 `json:"token"`
	RefreshToken string                 `json:"refreshToken"`
	Data         map[string]interface{} `json:"data"`
}
