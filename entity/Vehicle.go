package entity

type Vehicle struct {
	ID            int    `gorm:"type:int;not null;primary_key"`
	Name          string `gorm:"type:varchar(100);not null"`
	Brand         string `gorm:"type:varchar(100);not null"`
	TypeOfVehicle string `gorm:"type:varchar(10);not null"`
	CanDrive      bool   `gorm:"type:boolean"`
}
