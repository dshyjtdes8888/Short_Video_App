package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ningmeng/data"
	"strconv"
	"time"
)

// tempChat是一个map，用于临时存储用户之间的聊天消息。
var tempChat = map[string][]data.Message{}

// messageIdSequence是一个原子整数，用于生成消息的唯一ID。
var messageIdSequence = int64(1)

// ChatResponse结构体表示获取聊天消息的响应数据。
type ChatResponse struct {
	data.Response
	MessageList []data.Message `json:"message_list"`
}

// MessageAction 函数，保存用户之间的聊天消息。
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	// 检查用户是否存在，如果用户不存在，则返回StatusCode为1，表示用户不存在。
	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	userIdB, _ := strconv.Atoi(toUserId)

	message := data.Message{
		FromUserId: user.Id,
		ToUserId:   userIdB,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := data.Db.Exec("INSERT INTO messages (to_user_id, from_user_id, content, create_time) VALUES (?, ?, ?, STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'))", message.ToUserId, message.FromUserId, message.Content, message.CreateTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 5, StatusMsg: "Failed to create message entry"})
		return
	}

	c.JSON(http.StatusOK, data.Response{StatusCode: 0})
}

// MessageChat 函数用于获取用户之间的聊天消息。
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	preMsgTime := c.Query("pre_msg_time")

	// 检查用户是否存在，如果用户不存在，则返回StatusCode为1，表示用户不存在。
	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	userIdB, _ := strconv.Atoi(toUserId)

	// 转换 preMsgTime 为 int64
	preMsgTimeInt, err := strconv.ParseInt(preMsgTime, 10, 64)
	if err != nil {
		// 处理转换错误
		c.JSON(http.StatusBadRequest, data.Response{StatusCode: -1, StatusMsg: "Invalid preMsgTime"})
		return
	}

	// 查询数据库中用户之间的聊天消息列表，按时间排序
	var messages []data.Message
	if err := data.Db.Where("(from_user_id = ? AND to_user_id = ? AND create_time > ?) OR (from_user_id = ? AND to_user_id = ? AND create_time > ?)", user.Id, userIdB, preMsgTimeInt, userIdB, user.Id, preMsgTimeInt).Order("create_time").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to fetch messages from database"})
		return
	}

	fmt.Println("Fetched messages:", messages)
	c.JSON(http.StatusOK, ChatResponse{Response: data.Response{StatusCode: 0}, MessageList: messages})
}
