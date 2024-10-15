package user

import (
	"database/sql"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(e *echo.Group, us UserService) {
	handler := &UserHandler{
		userService: us,
	}

	users := e.Group("/users")

	users.GET("/:id", handler.getUserByID)
}

func (h *UserHandler) getUserByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	username, err := h.userService.GetUserByID(ctx, userId)
	if err != nil && err == sql.ErrNoRows {
		return c.JSON(404, map[string]interface{}{"message": "user not found"})
	}

	return c.JSON(200, map[string]interface{}{"username": username})
}
