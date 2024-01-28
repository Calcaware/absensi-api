package entities

import "time"

type UserEntity struct {
	ID          int                `gorm:"column:id;primaryKey" json:"id"`
	Fullname    string             `gorm:"column:fullname" json:"fullname"`
	NIK         string             `gorm:"column:nik" json:"nik"`
	Email       string             `gorm:"column:email" json:"email"`
	Password    string             `gorm:"column:password" json:"password"`
	Phone       string             `gorm:"column:phone" json:"phone"`
	Address     string             `gorm:"column:address" json:"address"`
	GenderID    int                `gorm:"column:gender_id" json:"gender_id"`
	Gender      GenderEntity       `gorm:"foreignKey:GenderID" json:"gender"`
	Birthdate   time.Time          `gorm:"column:birthdate" json:"birthdate"`
	CreatedAt   time.Time          `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time         `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Attendances []AttendanceEntity `gorm:"foreignKey:UserID" json:"attendances"`
}

type GenderEntity struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (UserEntity) TableName() string {
	return "users"
}
func (GenderEntity) TableName() string {
	return "genders"
}
