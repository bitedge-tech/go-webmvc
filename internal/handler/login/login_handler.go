package login

import (
	"fmt"
	"go-webmvc/internal/dto"
	"go-webmvc/internal/handler"
	"go-webmvc/internal/repository/model"
	"go-webmvc/internal/repository/query"
	"go-webmvc/internal/util"
	"image/color"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
)

// 验证验证码
// captchaId: 验证码ID
// captchaVal: 用户输入的验证码值
// 返回值: 验证结果，true表示验证通过，false表示验证失败
func verifyCaptcha(captchaId, captchaVal string) bool {
	return base64Captcha.DefaultMemStore.Verify(captchaId, captchaVal, true)
}

// Login 用户登录接口
// @Summary 用户登录
// @Description 用户通过用户名和密码登录，成功后返回JWT Token
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "登录请求参数"
// @Success 200 {object} dto.BaseResponse{data=dto.LoginResponse{}} "登录成功"
// @Failure 500 {object} dto.BaseResponse "登录失败"
// @Router /login/index [post]
func Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Failed(c, "参数绑定失败")
		return
	}
	if !verifyCaptcha(req.CaptchaId, req.CaptchaVal) {
		handler.Failed(c, "验证码错误")
		return
	}
	var user model.User
	// 用 query.SysUser 查询
	u, err := query.User.Where(query.User.Username.Eq(req.Username)).First()
	if err != nil || u == nil {
		handler.Failed(c, "用户名或密码错误")

		return
	}
	user = *u
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password+user.Salt)) != nil {
		handler.Failed(c, "用户名或密码错误")
		return
	}
	token, err := util.GenerateTokenWithUser(user.ID, user.Username)
	if err != nil {
		handler.Failed(c, "Token生成失败")
		return
	}
	resp := dto.LoginResponse{
		Token:    token,
		UserId:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}
	handler.Success(c, &resp)
}

// GetCaptcha 获取图形验证码接口
// @Summary 获取图形验证码
// @Description 获取图形验证码，返回验证码ID和Base64编码的图片
// @Tags 用户认证
// @Produce json
// @Success 200 {object} dto.BaseResponse{Data=map[string]string} "验证码生成成功"
// @Failure 200 {object} dto.BaseResponse "验证码生成失败"
// @Router /login/captcha [get]
func GetCaptcha(c *gin.Context) {

	id, b64, err := getCaptcha()
	if err != nil {
		handler.Failed(c, "验证码生成失败")
		return
	}
	resp := map[string]string{
		"captcha_id":  id,
		"captcha_img": b64,
	}
	handler.Success(c, &resp)
}

// GetCaptchaImg 获取图形验证码图片接口（直接返回图片HTML）
// @Summary 获取图形验证码图片（HTML）
// @Description 获取图形验证码图片，直接返回HTML img标签
// @Tags 用户认证
// @Produce html
// @Success 200 {string} string "验证码图片HTML"
// @Failure 200 {string} string "验证码生成失败"
// @Router /login/captcha_img [get]
func GetCaptchaImg(c *gin.Context) {
	_, b64, err := getCaptcha()
	if err != nil {
		c.String(http.StatusInternalServerError, "验证码生成失败")
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, fmt.Sprintf("<img src='%s' />", b64))
}

/* ************************ 以下是私有方法 **************************************/

func getCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverString(
		40,                               // 高度
		120,                              // 宽度
		4,                                // 长度
		5,                                // 噪点数量
		base64Captcha.OptionShowSineLine, // 干扰线类型  |base64Captcha.OptionShowSlimeLine|base64Captcha.OptionShowSineLine
		"abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789", // 字符集
		&color.RGBA{R: 236, G: 240, B: 241, A: 255},              // 背景色
		nil,                          // 字体存储
		[]string{"wqy-microhei.ttc"}, // 字体
	)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64, _, err := captcha.Generate()

	return id, b64, err
}
