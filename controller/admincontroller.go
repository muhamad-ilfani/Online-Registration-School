package controller

import (
	"context"
	"net/http"
	"project/config"
	"project/model"
	"strconv"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSingleUser(id int) (interface{}, error) {
	var users (*model.User)
	ctx := context.Background()
	if e := config.DB_user.FindOne(ctx, bson.M{"id": id}).Decode(&users); e != nil {
		return nil, e
	}
	return users, nil
}

func UpdateData(users model.User, id int) error {
	ctx := context.Background()
	_, err := config.DB_user.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": &users})
	if err != nil {
		return err
	}
	return nil
}
func UpdateDataSchool(school model.School, id int) error {
	ctx := context.Background()
	_, err := config.DB_school.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": &school})
	if err != nil {
		return err
	}
	return nil
}
func DeleteData(users model.User) error {
	ctx := context.Background()
	_, err := config.DB_user.DeleteOne(ctx, &users)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDataController(c echo.Context) error {
	id := c.QueryParam("id")

	idInt, err := strconv.Atoi(id)
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
	err = UpdateData(*dataUser, idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"data":    datas,
	})
}
func DeleteDataController(c echo.Context) error {
	id := c.QueryParam("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	users, err := GetSingleUser(id_int)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	dataUser, ok := users.(*model.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, "data user broken")
	}
	err = c.Bind(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = DeleteData(*dataUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
func RegisterSchool(c echo.Context) error {
	schools := model.School{}
	c.Bind(&schools)
	ctx := context.Background()
	_, err := config.DB_school.InsertOne(ctx, &schools)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    schools,
	})
}
func GetUserData(c echo.Context) error {
	var users []model.User
	ctx := context.Background()

	cursor_user, err := config.DB_user.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	defer cursor_user.Close(ctx)
	if err = cursor_user.All(ctx, &users); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    users,
	})
}
