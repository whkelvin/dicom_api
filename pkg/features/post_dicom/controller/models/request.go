package models
import (
  "mime/multipart"
)

type Request struct {
  Tag string
  File *multipart.FileHeader
}
