package features
import (
  "github.com/labstack/echo/v4"
  postDicom "github.com/whkelvin/dicom_api/pkg/features/post_dicom"
)

func Init(e *echo.Echo) {
  postDicom.Init(e, "/dicom")
}
