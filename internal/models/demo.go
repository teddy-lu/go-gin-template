package models

type Demo struct {
	BaseModel

	DemoField string `json:"demo_field"`

	CommonTimestampsField
}
