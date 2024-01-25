package model

type User struct {
	ID         int64  `gorm:"column:id;type:bigint(20) auto_increment;not null;primaryKey" json:"id"`
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	InsertTime int64  `gorm:"column:insert_time;type:bigint(20);not null" json:"insert_time"`
	UpdateTime int64  `gorm:"column:update_time;type:bigint(20);not null" json:"update_time"`
}

func (User) TableName() string {
	return "user"
}
