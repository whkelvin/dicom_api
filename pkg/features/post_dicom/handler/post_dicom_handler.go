package handler

import (
  "errors"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/file_storage_service"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler/models"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/dicom_service"
)

type IPostDicomHandler interface {
  PostDicom(req Request) (*Response, error)
}

type PostDicomHandler struct {
  FileStorageService IFileStorageService
  DicomService IDicomService
}

func (handler *PostDicomHandler) PostDicom(req Request) (*Response, error) {

  path, err := handler.FileStorageService.SaveFile(req.File)
  if(err != nil){
    return nil, errors.New("saving file failed.")
  }

  result := handler.DicomService.GetTagElement(path, req.TagGroup, req.TagElement)
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

  err = handler.DicomService.SaveAsPngs(path)
  if(err != nil){
    return nil, errors.New("saving pngs failed.")
  }

  return res, nil
}
