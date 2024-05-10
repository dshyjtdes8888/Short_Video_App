package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ningmeng/data"
	"strconv"
	"time"
)

// FavoriteAction 函数响应点赞，1为点赞，2为取消点赞。
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actiontype, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	// 检查用户是否存在，如果用户不存在，则返回StatusCode为1，表示用户不存在。
	var user data.User
	result := data.Db.Where("Name = ?", token).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if actiontype == 1 {
		// 更新用户的favorite_count和视频的喜欢状态和计数
		if err := data.Db.Model(&data.User{Id: user.Id}).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 2, StatusMsg: "Failed to update user favorite count"})
			return
		}
		/*		if err := tx.Model(&Video{Id: vid}).Update("is_favorite", true).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, Response{StatusCode: 3, StatusMsg: "Failed to update video favorite status"})
				return
			}*/
		if err := data.Db.Model(&data.Video{Id: int(vid)}).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 4, StatusMsg: "Failed to update video favorite count"})
			return
		}

		// 添加到favorite表
		favorite := data.Favorite{
			UserId:  user.Id,
			VideoId: int(vid),
		}
		if err := data.Db.Create(&favorite).Error; err != nil {
			c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 5, StatusMsg: "Failed to create favorite entry"})
			return
		}

		c.JSON(http.StatusOK, data.Response{StatusCode: 0})

	} else {
		// 处理取消点赞操作
		if user.FavoriteCount > 0 {
			if err := data.Db.Model(&data.User{Id: user.Id}).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 6, StatusMsg: "Failed to update user favorite count"})
				return
			}
			/*			if err := tx.Model(&Video{Id: vid}).Update("is_favorite", false).Error; err != nil {
						tx.Rollback()
						c.JSON(http.StatusInternalServerError, Response{StatusCode: 7, StatusMsg: "Failed to update video favorite status"})
						return
					}*/

			// 从favorite表中删除对应的条目
			if err := data.Db.Where("user_id = ? AND video_id = ?", user.Id, vid).Delete(data.Favorite{}).Error; err != nil {
				c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 8, StatusMsg: "Failed to delete favorite entry"})
				return
			}

			// 更新视频的favorite_count
			var video data.Video
			if err := data.Db.Where("id = ?", vid).Find(&video).Error; err != nil {
				c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 9, StatusMsg: "Failed to find video"})
				return
			}
			if video.FavoriteCount > 0 {
				if err := data.Db.Model(&data.Video{Id: int(vid)}).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update video favorite count"})
					return
				}
			}
			c.JSON(http.StatusOK, data.Response{StatusCode: 0})

		}
	}
}

// FavoriteList 函数返回用户的点赞视频列表。
func FavoriteList(c *gin.Context) {
	// 更改所有视频的IsFavorite字段为false
	if err := data.Db.Model(&data.Video{}).Where("1 = 1").Update("is_favorite", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to update videos"})
		return
	}

	// 初始化一个切片来存储用户已收藏的视频ID
	var favoriteVideoIDs []int

	// 获取用户的身份信息，例如令牌
	token := c.Query("token")
	userid := c.Query("user_id")

	if token != "" {
		// 根据令牌查找用户
		var user data.User
		result := data.Db.Where("id = ?", userid).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法找到用户"})
			return
		}

		// 查询用户的点赞视频ID
		var favorites []*data.Favorite
		result = data.Db.Where("user_id = ?", user.Id).Find(&favorites)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法查询收藏记录"})
			return
		}

		// 收集用户已收藏的视频ID
		for _, favorite := range favorites {
			favoriteVideoIDs = append(favoriteVideoIDs, favorite.VideoId)
		}

		// 更新用户已收藏的视频的 is_favorite 字段
		if len(favoriteVideoIDs) > 0 {
			if err := data.Db.Model(&data.Video{}).Where("id IN (?)", favoriteVideoIDs).Update("is_favorite", true).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法更新收藏的视频"})
				return
			}
		}

		var videos []*data.Video
		result = data.Db.Where("is_favorite = ?", true).Find(&videos)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch videos"})
			return
		}

		var videoList []data.VideoListItem

		for _, video := range videos {
			videoListItem := data.VideoListItem{
				ID:            video.Id,
				Author:        video.Author,
				PlayURL:       video.PlayUrl,
				CoverURL:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    video.IsFavorite,
				Title:         video.Title,
			}
			videoList = append(videoList, videoListItem)
		}

		c.JSON(http.StatusOK, data.FeedResponse{
			StatusCode: 0,
			StatusMsg:  "成功",
			VideoList:  videoList,
			NextTime:   int(time.Now().Unix()),
		})
	} else {
		c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "User doesn't login"})
	}
}
