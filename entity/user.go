package entity

type User struct{
	Id            uint64 `gorm:"primayKey" json:"id"`
	Name          string `gorm:"name" json:"name"`
	Password      string `gorm:"password"`
	Phone         string `gorm:"phone"`
	FollowCount   int64  `gorm:"follow_count" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"follower_count" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"is_follow" json:"is_follow,omitempty"`
	CreateAt      string `gorm:"create_at"`
	UpdateAt      string `gorm:"update_at"`
}

// 后台管理用户登录值对象ValueObject，用来传递参数
type AdminWebUserLoginVO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 后台管理前端，用户信息传输模型
type AdminWebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}