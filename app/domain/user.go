package domain

type User struct {
	ID       uint   `json:"id" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;type:varchar(50);unique"`
	Password []byte `json:"-" gorm:"column:password"`
}
