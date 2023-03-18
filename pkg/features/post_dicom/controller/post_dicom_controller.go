package controller

import (
  "errors"
  //"fmt"
  "net/http"
  "path/filepath"

  "github.com/labstack/echo/v4"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/controller/models"
)

func parsePostDicomRequest(c echo.Context) (Request, error){
  if !c.QueryParams().Has("tag") {
    return Request{}, errors.New("query param 'tag' is required.")
  }
  tag := c.QueryParams().Get("tag")

  file, err := c.FormFile("file")
  if(err != nil || filepath.Ext(file.Filename) != ".dcm") {
    return Request{}, errors.New("form file field 'file' is required.")
  }

  return Request{Tag: tag, File: file}, nil
}

func PostDicom(c echo.Context) error {
  //req, err := parsePostDicomRequest(c)
  //if(err != nil){
  //  return c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
  //}

  return c.String(http.StatusOK, "hello world")
}


