package routes

import (
  "github.com/labstack/echo/v4"
  postDicomRoute "github.com/whkelvin/dicom_api/pkg/api/features/post_dicom/route"
)

func Init(e *echo.Echo) {
  postDicomRoute.Init(e)
}

