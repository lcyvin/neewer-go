package main

import (
  "regexp"
  "fmt"
  "github.com/lcyvin/go-neewer/controllers/bluetooth"
  "github.com/lcyvin/go-neewer/vars/lights"
)

func main(){
  matchPatterns := []*regexp.Regexp{neewerlights.SL90.NameMatchPattern}
  knownLights := []string{}

  scanResults, err := bluetooth.Scan(matchPatterns, knownLights, 10, true)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(scanResults[0].LocalName())
  fmt.Println(scanResults[0].Address.String())

  fmt.Println("scanning for specific device...")
  specificScan := "E4:AF:00:CB:49:96"
  res, err := bluetooth.ScanForMacAddress(specificScan, 20)
  if err != nil {
    fmt.Println(err)
  }
  
  if res != nil {
    fmt.Println(res.LocalName())
  }
}
