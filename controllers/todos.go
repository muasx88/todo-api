package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muasx/todo_api/models"
)

func GetAll(c echo.Context) error {
	result, err := models.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func PostData(c echo.Context) error {
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.Store(todo.Name, todo.Complete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.Update(id, todo.Name, todo.Complete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := models.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
