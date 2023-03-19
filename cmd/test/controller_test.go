package main

import (
  "testing"
  "github.com/whkelvin/dicom_api/pkg/features/post_dicom/controller"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler/models"
  "github.com/labstack/echo/v4"
  "net/http"
  "net/http/httptest"
)

type PostDicomHandlerMock struct{}
func (mock *PostDicomHandlerMock) PostDicom(req Request) (*Response, error) {
  return &Response{}, nil
}

func TestPostDicomShouldReturn400WhenInputIsInvalid(t *testing.T){
  e := echo.New()

  req := httptest.NewRequest(http.MethodPost, "/", nil)

  rec := httptest.NewRecorder()
  ctx := e.NewContext(req, rec)

  handler := PostDicomHandlerMock{}
  controller := controller.PostDicomController{
    Echo: nil,
    Handler: &handler,
    Route: "",
  }

  controller.PostDicom(ctx)
  
  if rec.Code != 400 {
    t.Errorf("Result was incorrect, expected: %d, got: %d", 400, rec.Code)
  }
}

