package file_storage_service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(file *multipart.FileHeader) error {

  src, err := file.Open()
  if(err!= nil){
    return errors.New("file opened failed")
  }
  defer src.Close()

  dst, err := os.Create("shouldbeguid.dcm")
  if err != nil {
    return errors.New("file creation failed.")
  }
  defer dst.Close()

  if _, err = io.Copy(dst, src); err != nil {
    return err
  }

  return nil
}
