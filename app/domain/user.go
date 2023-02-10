package domain

type User struct {
	ID       uint   `gorm:"column:id;"`
	Name     string `gorm:"column:name;"`
	Email    string `gorm:"column:email;type:varchar(50);unique"`
	Password []byte `gorm:"column:password"`
}
