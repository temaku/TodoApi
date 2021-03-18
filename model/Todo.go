package model

type  Todo struct{
	Id          int  `json:"id"`
	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	Description   string `json:"description" gorm:"type:varchar(255);not null"`

}
type Error struct {
	Code int
	Message string
}
