package file_storage_service

import (
  "errors"
  "io"
  "mime/multipart"
  "os"
  "github.com/google/uuid"
  "fmt"
  "path/filepath"
)

type IFileStorageService interface {
  SaveFile(file *multipart.FileHeader) (string, error)
}

type FileStorageService struct {}

func (service *FileStorageService) SaveFile(file *multipart.FileHeader) (string, error) {
  src, err := file.Open()
  if(err!= nil){
    return "", errors.New("saving file failed.")
  }
  defer src.Close()
  uuid := uuid.New()
  
  dstPath := fmt.Sprintf("/home/whkelvin/Projects/golang/dicom_api/assets/uploaded/dcm/%v.dcm", uuid)
  dst, err := os.Create(dstPath)
  if err != nil {
    return "", errors.New("saving file failed.")
  }
  defer dst.Close()

  if _, err = io.Copy(dst, src); err != nil {
    return "", errors.New("saving file failed.")
  }

  absPath, err := filepath.Abs(dstPath)
  if err != nil {
    return "", errors.New("saving file failed.")
  }

  return absPath, nil
}
