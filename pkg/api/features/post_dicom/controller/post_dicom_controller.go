package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func PostDicom(c echo.Context) error {
  return c.String(http.StatusOK, "hello world")
}
