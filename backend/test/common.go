package test

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

// serverAddr表示测试服务器地址。
var serverAddr = "http://175.152.183.55:8080"

// testUserA和testUserB是用于测试的用户账号。
var testUserA = "ningmengTestUserA"
var testUserB = "ningmengTestUserB"

// newExpect函数创建并返回一个httpexpect.Expect实例，用于执行HTTP请求，并进行断言。
func newExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,              // 使用默认的HTTP客户端
		BaseURL:  serverAddr,                      // 设置基本URL为测试服务器地址
		Reporter: httpexpect.NewAssertReporter(t), // 使用默认的断言报告器
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true), // 启用调试打印机，将请求和响应信息输出到测试日志中
		},
	})
}

// getTestUserToken函数用于获取测试用户的认证token和用户ID。
func getTestUserToken(user string, e *httpexpect.Expect) (int, string) {
	// 尝试注册用户
	registerResp := e.POST("/ningmeng/user/register/").
		WithQuery("username", user).WithQuery("password", user).
		WithFormField("username", user).WithFormField("password", user).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := 0
	token := registerResp.Value("token").String().Raw()
	if len(token) == 0 {
		// 如果注册失败，则尝试登录用户
		loginResp := e.POST("/ningmeng/user/login/").
			WithQuery("username", user).WithQuery("password", user).
			WithFormField("username", user).WithFormField("password", user).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		loginToken := loginResp.Value("token").String()
		loginToken.Length().Gt(0)
		token = loginToken.Raw()
		userId = int(loginResp.Value("user_id").Number().Raw())
	} else {
		userId = int(registerResp.Value("user_id").Number().Raw())
	}
	return userId, token
}
