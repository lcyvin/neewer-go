package neewerlights

import (
	"regexp"
  "fmt"

	"github.com/lcyvin/go-neewer/types/light"
)

type KnownModelStruct struct {
  Models []*light.NeewerLight
}

func (km KnownModelStruct) MatchModel(name string) (light.NeewerLight, error) {
  for _,model := range km.Models {
    if model.NameMatchPattern.MatchString(name) {
      return *model, nil
    }
  }

  return light.NeewerLight{}, fmt.Errorf("no match found")
}

var LightPatterns map[*regexp.Regexp]light.NeewerLight = map[*regexp.Regexp]light.NeewerLight{
  SL90.NameMatchPattern: SL90,
}

var SL90 light.NeewerLight = light.NeewerLight{
  Model: "Neewer SL90 RGB",
  HSISupport: true,
  ANMSupport: true,
  NameMatchPattern: regexp.MustCompile(`NW\-\d+\&\w+$`),
}

var KnownModels KnownModelStruct = KnownModelStruct{

} 
