package models

type Demo struct {
	BaseModel

	DemoField string `json:"demo_field" gorm:"type:string;size:50"`

	CommonTimestampsField
}
