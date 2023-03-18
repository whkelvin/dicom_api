package handler

import (
	"errors"

	fileStorageService "github.com/whkelvin/dicom_api/pkg/features/post_dicom/data/file_storage_service"
	. "github.com/whkelvin/dicom_api/pkg/features/post_dicom/handler/models"
)

func PostDicomHandler(req Request) (Response, error) {

  err := fileStorageService.SaveFile(req.File)
  if(err != nil){
    return Response{}, errors.New("saving file failed.") 
  }


  return Response{
    PlaceHolder: "hello world",
  }, nil
}
