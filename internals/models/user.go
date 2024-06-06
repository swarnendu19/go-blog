package models

type User struct {
	ID       uint64    `gorm:"primary_key:auto_increment" json;"id"`
	Name     string    `gorm:"type:varchar(255)" json;"id"`
	Email    string    `gorm:"uniqueIndex; type: varchar(255)" json: "name"`
	Password string    `grom:"not null" json:"-"`
	Comments []Comment `json:"-"`
}
