package entities

type AttendanceEntity struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	Latitude  float64    `gorm:"column:latitude" json:"latitude"`
	Longitude float64    `gorm:"column:longitude" json:"longitude"`
	Status    string     `gorm:"column:status" json:"status"`
	CreatedAt int64      `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64      `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	User      UserEntity `gorm:"foreignKey:UserID" json:"user"`
	Location  string     `gorm:"-" json:"location"`
}

func (AttendanceEntity) TableName() string {
	return "attendances"
}
