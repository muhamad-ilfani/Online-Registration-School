package model

type User struct {
	Id       int     `json:"id" form:"id" bson:"id"`
	Name     string  `json:"name" form:"name" jsonschema:"required" bson:"name"`
	Email    string  `json:"email" form:"email" jsonschema:"required" bson:"email"`
	Password string  `json:"password" form:"password" jsonschema:"required" bson:"password"`
	Role     string  `json:"role" form:"role" bson:"role"`
	Ip       float32 `json:"ip" form:"ip" jsonschema:"required" bson:"ip"`
	SchoolId int     `json:"schoolid" form:"schoolid" bson:"schoolid"`
}

type UserResponse struct {
	Id    int    `json:"id" form:"id" bson:"id"`
	Name  string `json:"name" form:"name" bson:"name"`
	Email string `json:"email" form:"email" bson:"email"`
	Token string `json:"token" form:"token" bson:"token"`
}
