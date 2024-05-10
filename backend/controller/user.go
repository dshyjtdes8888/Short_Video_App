package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"ningmeng/data"
)

// UserLoginResponse 结构体表示用户登录的响应数据。
type UserLoginResponse struct {
	data.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserResponse 结构体表示用户信息的响应数据。
type UserResponse struct {
	data.Response
	User data.User `json:"user"`
}

// Register 函数用于处理用户注册的请求。
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user data.User
	// 检查用户是否已存在
	if err := data.Db.Where("name = ?", username).First(&user).Error; err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User already exists"},
		})
	} else {
		newUser := data.User{
			Name:     username,
			Password: password,
		}
		result := data.Db.Create(&newUser)
		if result.Error != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: data.Response{StatusCode: 1, StatusMsg: "User save failed"},
			})
		} else {
			data.Db.Where("name = ?", username).First(&user)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: data.Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    username,
			})
		}
	}
}

// Login 函数用于处理用户登录的请求。
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 查询用户信息
	var user data.User
	if err := data.Db.Where("name = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User doesn't exist or invalid credentials"},
		})
		return
	}

	// 验证密码
	if user.Password != password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "Invalid credentials"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: data.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    user.Name,
	})
}

// UserInfo 函数用于处理获取用户信息的请求。
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//id := c.Query("user_id")

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err == nil {

		c.JSON(http.StatusOK, UserResponse{
			Response: data.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
