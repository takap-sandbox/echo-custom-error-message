package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError);
	if !ok {
		c.Echo().DefaultHTTPErrorHandler(err, c)
		return
	}
	switch m := he.Message.(type) {
	case []string:
		err = c.JSON(he.Code, map[string]interface{}{"message": m})
	default:
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusUnauthorized, []string{"エラーなのだ", "エラーなのだ2"})
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
