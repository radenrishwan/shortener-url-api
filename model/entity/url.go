package entity

type Url struct {
	Id          string `gorm:"primaryKey"`
	Destination string
	Alias       string `gorm:"unique"`
	Clicked     uint64
	CreatedAt   string
	UpdatedAt   string
}
