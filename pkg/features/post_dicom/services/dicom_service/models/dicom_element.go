package models

type DicomElement struct {
  Tag string
  TagName string
  VR string
  VRRaw string
  VL uint32
  Value string
}
