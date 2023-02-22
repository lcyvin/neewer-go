package types

import (
  "regexp"
)

type LightStatus int

const (
  ON LightStatus = iota
  STBY
)

type LightMode int

const (
  HSI LightMode = iota
  CCT
  ANIM
)

type Device interface {
  Connect() error
  Write() error
  Match(string) (error, bool)
}

type NeewerLight struct {
  Model string
  MacAddress string
  HSISupport bool
  ANMSupport bool
  Status LightStatus
  Mode LightMode
  NameMatchPattern *regexp.Regexp
}

func Connect() error {
  return nil
}
