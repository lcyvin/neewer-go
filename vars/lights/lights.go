package lights

import (
	"regexp"
  "fmt"

	"github.com/lcyvin/go-neewer/types"
)

type KnownModelStruct struct {
  Models []*types.NeewerLight
}

func (km KnownModelStruct) MatchModel(name string) (types.NeewerLight, error) {
  for _,model := range km.Models {
    if model.NameMatchPattern.MatchString(name) {
      return *model, nil
    }
  }

  return types.NeewerLight{}, fmt.Errorf("no match found")
}

var LightPatterns map[*regexp.Regexp]types.NeewerLight = map[*regexp.Regexp]types.NeewerLight{
  SL90.NameMatchPattern: SL90,
}

var SL90 types.NeewerLight = types.NeewerLight{
  Model: "Neewer SL90 RGB",
  HSISupport: true,
  ANMSupport: true,
  NameMatchPattern: regexp.MustCompile(`NW\-\d+\&\w+$`),
}

var KnownModels KnownModelStruct = KnownModelStruct{

} 
