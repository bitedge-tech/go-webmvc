package dto

type LoginRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	CaptchaId  string `json:"captcha_id" binding:"required"`
	CaptchaVal string `json:"captcha_val" binding:"required"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type CaptchaResponse struct {
	CaptchaID  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
}

type CaptchaResponseWrapper struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data CaptchaResponse `json:"data"`
}
