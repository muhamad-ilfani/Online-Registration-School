package model

type School struct {
	Id       int    `json:"id" form:"id" jsonschema:"required" bson:"id"`
	Name     string `json:"name" form:"name" jsonschema:"required" bson:"name"`
	Capacity int    `json:"capacity" form:"capacity" jsonschema:"required" bson:"capacity"`
	Students []User `json:"students" form:"students" bson:"students"`
}
