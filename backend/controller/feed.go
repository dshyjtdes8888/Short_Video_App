package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ningmeng/data"
	"time"
)

func Feed(c *gin.Context) {
	var videos []*data.Video
	result := data.Db.Preload("Author").Find(&videos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch videos"})
		return
	}

	token := c.Query("token")

	// 更改所有视频的IsFavorite字段为false
	if err := data.Db.Model(&data.Video{}).Where("1 = 1").Update("is_favorite", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to update videos"})
		return
	}

	// 更改所有用户的IsFollow字段为false
	if err := data.Db.Model(&data.User{}).Where("1 = 1").Update("is_follow", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to update users"})
		return
	}

	// 初始化一个切片来存储用户已点赞的视频ID
	var favoriteVideoIDs []int

	//初始化一个切片来存储用户已关注的用户ID
	var relationUserIDs []int

	// 检查用户是否已登录并且具有有效的token参数
	if token != "" {
		var user data.User
		result = data.Db.Where("name = ?", token).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法找到用户"})
			return
		}

		// 查询用户的点赞视频ID
		var favorites []*data.Favorite
		result = data.Db.Where("user_id = ?", user.Id).Find(&favorites)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法查询收藏记录"})
			return
		}

		// 收集用户已点赞的视频ID
		for _, favorite := range favorites {
			favoriteVideoIDs = append(favoriteVideoIDs, favorite.VideoId)
		}

		// 更新用户已点赞的视频的 is_favorite 字段
		if len(favoriteVideoIDs) > 0 {
			if err := data.Db.Model(&data.Video{}).Where("id IN (?)", favoriteVideoIDs).Update("is_favorite", true).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法更新收藏的视频"})
				return
			}
		}

		// 查询用户的关注用户ID
		var relations []*data.Relation
		result = data.Db.Where("user_id = ?", user.Id).Find(&relations)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法查询关注记录"})
			return
		}

		// 收集用户已关注的用户ID
		for _, relation := range relations {
			relationUserIDs = append(relationUserIDs, relation.FollowId)
		}

		// 更新用户已关注用户的 is_follow 字段
		if len(relationUserIDs) > 0 {
			if err := data.Db.Model(&data.User{}).Where("id IN (?)", relationUserIDs).Update("is_follow", true).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "无法更新关注的用户"})
				return
			}
		}
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

	feedResponse := data.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		NextTime:   int(time.Now().Unix()),
		VideoList:  videoList,
	}

	c.JSON(http.StatusOK, feedResponse)
}
