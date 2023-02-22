package main

import (
  "regexp"
  "fmt"
  "github.com/lcyvin/go-neewer/controllers/bluetooth"
  "github.com/lcyvin/go-neewer/vars/lights"
)

func main(){
  matchPatterns := []*regexp.Regexp{lights.SL90.NameMatchPattern}
  knownLights := []string{}

  scanResults, err := bluetooth.Scan(matchPatterns, knownLights, 10, true)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(scanResults[0].LocalName())
}
