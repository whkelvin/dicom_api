package post_dicom

import (
  "github.com/labstack/echo/v4"
  "github.com/whkelvin/dicom_api/pkg/features/post_dicom/controller"
  "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/dicom_service"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/file_storage_service"
)

func Init(e *echo.Echo, route string) {
  fileStorageService := &FileStorageService{}
  dicomService := &DicomService{}
  handler := &handler.PostDicomHandler{FileStorageService: fileStorageService, DicomService: dicomService}

  controller := controller.PostDicomController{
    Echo: e,
    Handler: handler,
    Route: route,
  }

  controller.Init()
}
