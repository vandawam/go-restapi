package models

type Product struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}
