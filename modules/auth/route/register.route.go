package auth_route

import (
	auth_controller "Chat/modules/auth/controller"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Group) {
	e.POST("/register", auth_controller.Register)
}
