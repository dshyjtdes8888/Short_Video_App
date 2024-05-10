package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ningmeng/data"
)

// UserListResponse 结构体表示获取用户列表的响应数据。
type UserListResponse struct {
	data.Response
	UserList []data.User `json:"user_list"`
}

// RelationAction 用于关注和取消关注的响应
func RelationAction(c *gin.Context) {
	token := c.Query("token") //用户名
	actionType := c.Query("action_type")
	ToUserId := c.Query("to_user_id")

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "User doesn't exist"})
	} else {
		var follow data.User
		if err := data.Db.Where("id = ?", ToUserId).First(&follow).Error; err != nil {
			c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "FollowUser doesn't exist"})
		}
		if actionType == "1" {
			relation := data.Relation{
				User:     user,
				Follow:   follow,
				UserId:   user.Id,
				FollowId: follow.Id,
			}
			result := data.Db.Create(&relation)
			if result.Error != nil {
				c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "relation failed"})
			} else {
				c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "relation save successfully"})
			}

			// 更新自己的follow_count
			if user.FollowCount >= 0 {
				if err = data.Db.Model(&user).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update user follow count"})
					return
				}
			}

			// 更新关注者的follower_count
			if follow.FollowerCount >= 0 {
				if err = data.Db.Model(&follow).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update user follower count"})
					return
				}
			}

		} else {
			var relation data.Relation
			if err := data.Db.Where("user_id=?", user.Id).Where("follow_id = ?", ToUserId).First(&relation).Error; err != nil {
				c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "FollowUser doesn't exist"})
			} else {
				err = data.Db.Delete(relation).Error
				if err != nil {
					c.JSON(http.StatusOK, data.Response{StatusCode: 1, StatusMsg: "relation delete failed"})
				} else {
					c.JSON(http.StatusOK, data.Response{StatusCode: 0, StatusMsg: "relation delete successfully"})
				}
			}

			// 更新用户的follow_count
			if user.FollowCount >= 0 {
				if err = data.Db.Model(&user).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update user follow count"})
					return
				}
			}

			// 更新用户的follower_count
			if follow.FollowerCount >= 0 {
				if err = data.Db.Model(&follow).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
					c.JSON(http.StatusInternalServerError, data.Response{StatusCode: 10, StatusMsg: "Failed to update user follower count"})
					return
				}
			}

		}
	}
}

// FollowList 函数返回用户的关注列表
func FollowList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err == nil {
		var follows []*data.Relation
		result := data.Db.Preload("User").Where("user_id = ?", userId).Find(&follows)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch relations"})
			return
		}

		// 更改所有用户的IsFollow字段为false
		if err := data.Db.Model(&data.User{}).Where("1 = 1").Update("is_follow", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to update users"})
			return
		}

		//初始化一个切片来存储用户已关注的用户ID
		var relationUserIDs []int

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

		var userList []data.User
		for _, follow := range follows {
			var followuser data.User
			result = data.Db.Where("id = ?", follow.FollowId).First(&followuser)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch relations"})
				return
			}
			userList = append(userList, followuser)
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: data.Response{
				StatusCode: 0,
			},
			UserList: userList, // DemoUser是演示用户，实际列表内容可根据需求更改
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}

// FollowerList 函数返回用户的粉丝列表
func FollowerList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err == nil {
		var followers []*data.Relation
		result := data.Db.Preload("User").Where("follow_id = ?", userId).Find(&followers)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch relations"})
			return
		}

		// 更改所有用户的IsFollow字段为false
		if err := data.Db.Model(&data.User{}).Where("1 = 1").Update("is_follow", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to update users"})
			return
		}

		//初始化一个切片来存储用户已关注的用户ID
		var relationUserIDs []int

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

		var userList []data.User
		for _, follower := range followers {
			var followeruser data.User
			result = data.Db.Where("id = ?", follower.UserId).First(&followeruser)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch relations"})
				return
			}
			userList = append(userList, followeruser)
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: data.Response{
				StatusCode: 0,
			},
			UserList: userList, // DemoUser是演示用户，实际列表内容可根据需求更改
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// FriendList 函数返回所有的好友列表
func FriendList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")

	var user data.User
	if err := data.Db.Where("name = ?", token).First(&user).Error; err == nil {
		// 找到互相关注的用户
		var userList []data.User
		result := data.Db.Raw("SELECT u.* FROM users u "+
			"JOIN relations r1 ON u.id = r1.user_id "+
			"JOIN relations r2 ON u.id = r2.follow_id "+
			"WHERE r1.follow_id = ? AND r2.user_id = ?", userId, userId).Scan(&userList)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "status_msg": "Failed to fetch friends"})
			return
		}

		c.JSON(http.StatusOK, UserListResponse{
			Response: data.Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: data.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
