package user

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))
	if users[id] == nil {
		return fmt.Errorf("user of id %v was not found", id)
	}

	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[id].Age = u.Age
	users[id].Email = u.Email
	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}
