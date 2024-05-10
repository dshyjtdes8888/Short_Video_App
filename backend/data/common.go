package data

// Response 结构体表示API接口返回的通用响应数据。
type Response struct {
	StatusCode int32  `json:"status_code"`          // 响应状态码
	StatusMsg  string `json:"status_msg,omitempty"` // 响应状态消息，可选字段
}

// Video 结构体表示视频信息。
type Video struct {
	Id            int    `json:"id,omitempty"`
	Author        User   `json:"author" gorm:"foreignKey:AuthorId"` // 此处应根据您的模型定义进行调整
	AuthorId      int    `json:"aid"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count,omitempty"`
	CommentCount  int    `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
}

// Comment 结构体表示评论信息。
type Comment struct {
	Id         int    `json:"id,omitempty"`                  // 评论ID
	User       User   `json:"user" gorm:"foreignKey:UserId"` // 评论用户信息
	UserId     int    `json:"user_id"`
	Video      Video  `json:"video" gorm:"foreignKey:VideoId"`
	VideoId    int    `json:"video_id"`
	Content    string `json:"content,omitempty"`     // 评论内容
	CreateDate string `json:"create_date,omitempty"` // 评论创建日期
}

// Relation 结构体表示关注信息。
type Relation struct {
	Id       int  `json:"id,omitempty"`                  // ID
	User     User `json:"user" gorm:"foreignKey:UserId"` //粉丝
	UserId   int  `json:"user_id"`
	Follow   User `json:"follow" gorm:"foreignKey:FollowId"` // 关注的人
	FollowId int  `json:"follow_id"`
}

type Favorite struct {
	Id      int `json:"id,omitempty"`
	UserId  int `json:"user_id,omitempty"`
	VideoId int `json:"video_id,omitempty"`
}

// User 结构体表示用户信息。
type User struct {
	Id              int    `json:"id,omitempty"`               // 用户ID
	Name            string `json:"name,omitempty"`             // 用户名
	Password        string `json:"password,omitempty"`         //用户密码
	FollowCount     int    `json:"follow_count,omitempty"`     // 关注数
	FollowerCount   int    `json:"follower_count,omitempty"`   // 粉丝数
	IsFollow        bool   `json:"is_follow,omitempty"`        // 是否已关注该用户
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	BackgroundImage string `json:"background_image,omitempty"` //背景图片
	Signature       string `json:"signature,omitempty"`        //签名
	TotalFavorited  string `json:"total_favorited,omitempty"`
	WorkCount       int    `json:"work_count,omitempty"`
	FavoriteCount   int    `json:"favorite_count,omitempty"`
}

// VideoListItem 自定义结构体用于映射视频列表项
type VideoListItem struct {
	ID            int    `json:"id"`
	Author        User   `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// FeedResponse 自定义结构体用于接口响应
type FeedResponse struct {
	StatusCode int             `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	NextTime   int             `json:"next_time"`
	VideoList  []VideoListItem `json:"video_list"`
}

// Message 结构体表示聊天消息。
type Message struct {
	Id         int    `json:"id,omitempty"`           // 消息ID
	ToUserId   int    `json:"to_user_id,omitempty"`   //消息接收者
	FromUserId int    `json:"from_user_id,omitempty"` //消息发送者
	Content    string `json:"content,omitempty"`      // 消息内容
	CreateTime string `json:"create_time,omitempty"`  // 消息创建时间
}

// MessageSendEvent 结构体表示发送聊天消息的事件。
type MessageSendEvent struct {
	UserId     int    `json:"user_id,omitempty"`     // 发送消息的用户ID
	ToUserId   int    `json:"to_user_id,omitempty"`  // 接收消息的用户ID
	MsgContent string `json:"msg_content,omitempty"` // 消息内容
}

// MessagePushEvent 结构体表示推送聊天消息的事件。
type MessagePushEvent struct {
	FromUserId int    `json:"user_id,omitempty"`     // 消息发送者的用户ID
	MsgContent string `json:"msg_content,omitempty"` // 消息内容
}
