package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	user := &User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	user.Id = seq
	user.Uuid = uuid.New()
	users[user.Id] = user

	seq += 1

	return c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}
