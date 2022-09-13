package table

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
