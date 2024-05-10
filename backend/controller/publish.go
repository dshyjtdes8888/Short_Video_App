package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ningmeng/data"
	"os/exec"
	"path/filepath"
	"time"
)

// VideoListResponse 结构体表示获取视频列表的响应数据。
type VideoListResponse struct {
	data.Response
	VideoList []*data.Video `json:"video_list"`
}

// Publish 函数用于发布视频，将上传的文件保存到public目录。
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		fmt.Println("token is null")
	}

	// 检查用户是否存在，如果用户不存在，则返回StatusCode为1，表示用户不存在。
	var user data.User
	result := data.Db.Where("Name = ?", token).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 获取上传的文件
	data2, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, data.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 获取上传文件的基本文件名
	filename := filepath.Base(data2.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	// 拼接最终保存文件的路径
	saveFile := filepath.Join("./public/", finalName)

	playurl := "http://192.168.47.9:8080/static/" + finalName
	title := c.PostForm("title")

	//发布的用户作品数+1
	data.Db.Model(&data.User{Id: user.Id}).Update("work_count", gorm.Expr("work_count + ?", 1))

	// 保存上传的文件到public目录中
	if err := c.SaveUploadedFile(data2, saveFile); err != nil {
		c.JSON(http.StatusOK, data.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 提取封面图片
	coverImageFile := filepath.Join("./public/", "cover_"+finalName+".jpg")
	err = extractCoverImage(saveFile, coverImageFile)
	if err != nil {
		c.JSON(http.StatusOK, data.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	coverurl := "http://192.168.47.9:8080/static/" + "cover_" + finalName + ".jpg"

	v := &data.Video{PlayUrl: playurl, CoverUrl: coverurl, AuthorId: user.Id, Title: title}
	res := data.Db.Create(v) //传递结构体，链式调用，返回对象
	fmt.Println(res.Error)   //获取error，用链式调用返回的对象获取
	fmt.Println(v.Id)        //返回插入数据的主键，自增

	c.JSON(http.StatusOK, data.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// 提取视频封面图像
func extractCoverImage(videoPath, outputPath string) error {
	// 使用 Ffmpeg 提取封面图像
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:01", "-vframes", "1", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to extract cover image: %v", err)
	}
	return nil
}

// PublishList 函数返回所有用户的发布视频列表
func PublishList(c *gin.Context) {
	userid := c.Query("user_id")
	var videos []*data.Video
	result := data.Db.Where("author_id = ?", userid).Find(&videos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}

	var videoList []data.VideoListItem // 自定义的结构体，用于映射视频列表项

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
		StatusMsg:  "Success",
		VideoList:  videoList,              // DemoVideos是演示视频列表，实际列表内容可根据需求更改
		NextTime:   int(time.Now().Unix()), // 设置NextTime为当前时间的Unix时间戳，用于下次请求时分页加载数据
	})
}
