package controller

import (
	"context"
	"net/http"
	"project/config"
	"project/middleware"
	"project/model"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSchoolData(c echo.Context) error {
	var schools []model.School
	ctx := context.Background()

	cursor_school, err := config.DB_school.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	defer cursor_school.Close(ctx)
	if err = cursor_school.All(ctx, &schools); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    schools,
	})
}
func RegisterAccount(c echo.Context) error {
	users := model.User{}
	c.Bind(&users)
	ctx := context.Background()
	users.Role = "user"
	_, err := config.DB_user.InsertOne(ctx, &users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    users,
	})
}
func InsertUser(user model.User, id int) error {
	ctx := context.Background()
	_, err := config.DB_school.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$push": bson.M{"students": user}})
	if err != nil {
		return err
	}
	return nil
}
func RegisterUser(c echo.Context) error {
	Id := c.QueryParam("id")
	idInt, err := strconv.Atoi(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	datas, err := GetSingleUser(idInt)
	if err != nil || datas == nil {
		return c.JSON(http.StatusNotFound, err)
	}
	dataUser, ok := datas.(*model.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, "user data broken")
	}
	err = c.Bind(&datas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	IdSchool := dataUser.SchoolId
	err = InsertUser(*dataUser, IdSchool)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"data":    datas,
	})
}

func LoginUserController(c echo.Context) error {
	users := model.User{}
	ctx := context.Background()
	c.Bind(&users)
	err := config.DB_user.FindOne(ctx, bson.M{"email": users.Email, "password": users.Password}).Decode(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err,
		})
	}
	token, err := middleware.CreateToken(users.Id, users.Email, users.Name, users.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err,
		})
	}
	usersResponse := model.UserResponse{users.Id, users.Name, users.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"data":    usersResponse,
	})
}
