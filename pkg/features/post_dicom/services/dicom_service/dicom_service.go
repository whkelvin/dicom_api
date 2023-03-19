package dicomservice

import (
  "errors"
  "fmt"
  "image/png"
  "os"
  "github.com/suyashkumar/dicom"
  "github.com/suyashkumar/dicom/pkg/tag"
  . "github.com/whkelvin/dicom_api/pkg/features/post_dicom/services/dicom_service/models"
  "github.com/google/uuid"
)

type IDicomService interface {
  GetTagElement(dicomFilePath string, group uint16, element uint16) (*DicomElement)
  SaveAsPngs(dicomFilePath string) error
}

type DicomService struct{}

func (service *DicomService) GetTagElement(dicomFilePath string, group uint16, element uint16) (*DicomElement){
  content, _ := dicom.ParseFile(dicomFilePath, nil)

  var t tag.Tag = tag.Tag{Group: group, Element: element}
  result, err := content.FindElementByTag(t)

  if(err != nil){
    return nil
  }

  var tagName string
  if tagInfo, err := tag.Find(result.Tag); err == nil {
    tagName = tagInfo.Name
  }

  var dicomElement *DicomElement = &DicomElement{
    Tag: result.Tag.String(),
    TagName: tagName,
    VR: result.ValueRepresentation.String(),
    VRRaw: result.RawValueRepresentation,
    VL: result.ValueLength,
    Value: result.Value.String(),
  }

  return dicomElement
}

func (service *DicomService) SaveAsPngs(dicomFilePath string) error {
  dataset, err := dicom.ParseFile(dicomFilePath, nil)
  if(err != nil){
    return errors.New("cannot parse dicom file.")
  }

  pixelDataElement, err := dataset.FindElementByTag(tag.PixelData)
  if(err != nil) {
    return err
  }

  pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)
  uuid := uuid.New()

  for _, fr := range pixelDataInfo.Frames {
    img, err := fr.GetImage()
    if err == nil {
      f, err := os.Create(fmt.Sprintf("/home/whkelvin/Projects/golang/dicom_api/assets/uploaded/png/%v.png", uuid))
      if err == nil {
        png.Encode(f, img)
      }
      f.Close()
    }
  }
  return nil
}
