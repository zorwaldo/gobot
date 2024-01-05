package model

type Chat struct {
	Id   int    `gorm:"column:chatid"`
	Name string `gorm:"column:name"`
}
