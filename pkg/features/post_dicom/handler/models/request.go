package models
import (
  "mime/multipart"
)

type Request struct {
  TagGroup uint16
  TagElement uint16
  File *multipart.FileHeader
}
