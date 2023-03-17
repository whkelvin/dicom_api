package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/gommon/log"
  "github.com/whkelvin/dicom_api/pkg/api/routes"
)

func main() {
  var e *echo.Echo = echo.New()

  log.SetLevel(log.DEBUG)
  log.SetHeader("${time_rfc3339} ${level}")

  routes.Init(e)

  e.Logger.Fatal(e.Start(":1323"))
}
