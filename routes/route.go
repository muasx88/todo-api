package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	c "github.com/muasx/todo_api/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo API")
	})

	e.GET("/todos", c.GetAll)
	e.POST("/todos", c.PostData)
	e.PUT("/todos/:id", c.UpdateData)
	e.DELETE("/todos/:id", c.DeleteData)

	return e
}
