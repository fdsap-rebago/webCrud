package model

type (
	Users struct {
		Id       uint   `json:"id" gorm:"primaryKey"`
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
