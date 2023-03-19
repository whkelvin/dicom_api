package controller

import (
  "errors"
  "fmt"
  "net/http"
  "path/filepath"
  "regexp"
  "github.com/labstack/echo/v4"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/controller/models"
  "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler"
  handlerModels "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler/models"
  "strconv"
)

/*
HERE IS CHAPGPT's EXPLAINATION OF THAT REGEX:
  ^ - Start of string
  * - Zero or more spaces (optional)
  \( - Opening parenthesis
  ([0-9a-fA-F]{4}) - 4 characters that are either digits (0-9) or letters (a-f or A-F) (captured as submatch 1)
  , * - Comma followed by zero or more spaces
  ([0-9a-fA-F]{4}) - 4 characters that are either digits (0-9) or letters (a-f or A-F) (captured as submatch 2)
  \) - Closing parenthesis
  *$ - Zero or more spaces followed by end of string (optional)
*/
const TAG_REGEX = `^ *\(([0-9a-fA-F]{4}), *([0-9a-fA-F]{4})\) *$`

func parsePostDicomRequest(c echo.Context) (Request, error){
  if !c.QueryParams().Has("tag") {
    return Request{}, errors.New("query param 'tag' is required.")
  }
  tag := c.QueryParams().Get("tag")

  exp := regexp.MustCompile(TAG_REGEX)

  submatches := exp.FindStringSubmatch(tag)

  if len(submatches) != 3 {
    return Request{}, errors.New("query param 'tag' must be in the format of '(ffff, ffff)'")
  }

  file, err := c.FormFile("file")
  if(err != nil || filepath.Ext(file.Filename) != ".dcm") {
    return Request{}, errors.New("form file field 'file' is required.")
  }

  return Request{Tag: tag, File: file}, nil
}

func parseTag(tag string) (uint16, uint16){
  exp := regexp.MustCompile(TAG_REGEX)
  submatches := exp.FindStringSubmatch(tag)

  group64, _:= strconv.ParseUint(submatches[1], 16, 16)
  group16 := uint16(group64)

  element64, _:= strconv.ParseUint(submatches[2], 16, 16)
  element16 := uint16(element64)

  return group16, element16
}

type PostDicomController struct {
  Echo *echo.Echo
  Handler handler.IPostDicomHandler
  Route string
}

func (controller *PostDicomController) Init(){
  controller.Echo.POST(controller.Route, controller.PostDicom)
}

func (controller *PostDicomController) PostDicom(c echo.Context) error {
  req, err := parsePostDicomRequest(c)
  if(err != nil){
    return c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
  }
  
  group, element := parseTag(req.Tag)
  var handlerReq handlerModels.Request = handlerModels.Request{
    TagGroup: group,
    TagElement: element,
    File: req.File,
  }

  handlerRes, err := controller.Handler.PostDicom(handlerReq)
  if(err != nil){
    return c.String(http.StatusInternalServerError, "post dicom file failed.")
  }

  if(handlerRes == nil){
    return c.String(http.StatusNotFound, fmt.Sprintf("Tag %v not found.", req.Tag))
  }

  var res Response = Response{
    Tag: handlerRes.Tag,
    TagName: handlerRes.TagName,
    VR: handlerRes.VR,
    VRRaw: handlerRes.VRRaw,
    VL: handlerRes.VL,
    Value: handlerRes.Value,
  }

  return c.JSON(http.StatusOK, res)
}

