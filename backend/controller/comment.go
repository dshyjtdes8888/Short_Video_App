package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ningmeng/data"
	"time"
)

// CommentListResponse 结构体表示获取评论列表的响应数据。
type CommentListResponse struct {
	data.Response
	CommentList []CommentListItem `json:"comment_list,omitempty"` // 评论列表，可选字段
}

// CommentActionResponse 结构体表示评论操作的响应数据。
type CommentActionResponse struct {
	data.Response
	Comment data.Comment `json:"comment,omitempty"` // 评论信息，可选字段
}

// CommentAction 评论和删除评论。
func CommentAction(c *gin.Context) {
	token := c.Query("token") //用户名
	actionType := c.Query("action_type")
	content := c.Query("comment_text")
	videoId := c.Query("video_id")

	var video data.Video
	data.Db.Where("id = ?", videoId).First(&video)

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "User doesn't exist"})
	} else {
		if actionType == "1" {
			comment := data.Comment{
				User:       user,
				Video:      video,
				Content:    content,
				CreateDate: time.Now().Format("2006-01-02 15:04"),
			}
			result := data.Db.Create(&comment)
			if result.Error != nil {
				c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "Comment save failed"})
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: data.Response{StatusCode: 1, StatusMsg: "Comment save successful"},
					Comment:  comment,
				})
			}

			// 更新视频的comment_count
			if video.CommentCount >= 0 {
				if err = data.Db.Model(&video).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update video comment count"})
					return
				}
			}

		} else {
			commentId := c.Query("comment_id")
			var comment data.Comment
			if err = data.Db.Where("id = ?", commentId).First(&comment).Error; err != nil {
				c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "Comment doesn't exist"})
			} else {
				err = data.Db.Delete(comment).Error
				if err != nil {
					c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "Comment delete failed"})
				} else {
					c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "Comment delete successfully"})
				}
			}

			// 更新视频的comment_count
			if video.CommentCount >= 0 {
				if err = data.Db.Model(&video).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update video comment count"})
					return
				}
			}

		}
	}
}

// CommentList 返回所有视频的评论列表。
func CommentList(c *gin.Context) {
	videoId := c.Query("video_id")

	var comments []*data.Comment
	result := data.Db.Preload("User").Preload("Video").Where("video_id = ?", videoId).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch comments"})
		return
	}

	var commentList []CommentListItem

	for _, comment := range comments {
		commentListItem := CommentListItem{
			ID:         comment.Id,
			User:       comment.User,
			Video:      comment.Video,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		commentList = append(commentList, commentListItem)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    data.Response{StatusCode: 0},
		CommentList: commentList,
	})
}

type CommentListItem struct {
	ID         int        `json:"id"`
	User       data.User  `json:"user"`
	Video      data.Video `json:"video"`
	Content    string     `json:"content"`
	CreateDate string     `json:"create_date"`
}
