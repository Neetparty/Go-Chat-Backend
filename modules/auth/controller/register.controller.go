package auth_controller

import (
	"Chat/dto"
	auth_service "Chat/modules/auth/service"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) (err error) {

	user := dto.TRegister{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Image:    c.FormValue("image"),
	}

	e := c.Bind(&user)

	if e != nil {
		return e
	}

	newUser := auth_service.Register_Service(user)

	return c.JSON(200, newUser)

}
