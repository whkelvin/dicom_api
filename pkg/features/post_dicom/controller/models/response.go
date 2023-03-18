package models

type Response struct {
  Tag string `json:"tag"`
  TagName string `json:"tagName"`
  VR string `json:"vr"`
  VRRaw string `json:"vrRaw"`
  VL uint32 `json:"vl"`
  Value string `json:"value"`
}
