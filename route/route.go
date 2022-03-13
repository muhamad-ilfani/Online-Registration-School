package route

import (
	"project/constants"
	"project/controller"
	"project/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.RegisterAccount)

	eUser := e.Group("/user")
	eUser.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUser.Use(echo.WrapMiddleware(middleware.RegisterPermission))
	eUser.GET("/register", controller.GetSchoolData)
	eUser.POST("/register", controller.RegisterUser)

	eAdmin := e.Group("/admin")
	eAdmin.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eAdmin.Use(echo.WrapMiddleware(middleware.AdminOnly))
	eAdmin.GET("/get", controller.GetUserData)
	eAdmin.POST("/register", controller.RegisterSchool)
	eAdmin.PUT("/update", controller.UpdateDataController)
	eAdmin.PUT("/delete", controller.DeleteDataController)

	return e
}
