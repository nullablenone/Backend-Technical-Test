package company

type Company struct {
	ID   string `gorm:"type:uuid;primary_key"`
	Name string `gorm:"type:text;not null"`
}
