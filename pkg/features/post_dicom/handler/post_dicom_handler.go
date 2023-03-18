package handler

import (
  "errors"
  fileStorageService "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/file_storage_service"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler/models"
  dicomService "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/dicom_service"
  "github.com/labstack/gommon/log"
)

func PostDicomHandler(req Request) (*Response, error) {

  path, err := fileStorageService.SaveFile(req.File)
  if(err != nil){
    log.Error("saving file failed.")
    return nil, errors.New("saving file failed.")
  }

  result := dicomService.GetTagElement(path, req.TagGroup, req.TagElement)
  if(result == nil){
    return nil, nil
  }

  var res *Response = &Response{
    Tag: result.Tag,
    TagName: result.TagName,
    VR: result.VR,
    VRRaw: result.VRRaw,
    VL: result.VL,
    Value:result.Value,
  }

  err = dicomService.SaveAsPngs(path)
  if(err != nil){
    log.Error("saving pngs failed.")
    return nil, errors.New("saving pngs failed.")
  }

  return res, nil
}
