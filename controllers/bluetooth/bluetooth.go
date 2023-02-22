package bluetooth

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func Scan(matchPatterns []*regexp.Regexp, knownDevices []string, timeout int, exitAfterMatch bool) ([]*bluetooth.ScanResult, error) {
  devs := make([]*bluetooth.ScanResult, 0)

  err := adapter.Enable()
  if err != nil {
    log.Fatalf("could not enable bluetooth adapter: %v", err)
  }

  fmt.Println("scanning...")
  timer := time.NewTicker(time.Duration(timeout) * time.Second)
  results := make(chan *bluetooth.ScanResult)

  go ScanForLight(adapter, results)

  for {
    select {
    case tick := <- timer.C:
      if len(devs) != 0 {
        return devs, nil
      }

      return devs, fmt.Errorf("Scanner timed out after %d seconds", tick.Second())

    case dev := <- results:
      for _,matchStr := range matchPatterns {
        if matchStr.MatchString(dev.LocalName()) {
        devs = append(devs, dev)
        if exitAfterMatch {
          return devs, nil
          }
        }
      }
    }
  }
}

func ScanForLight(adapter *bluetooth.Adapter, resChan chan *bluetooth.ScanResult) {
  // start our scan
  adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
    resChan <- &device
  })
}

func isKnownDevice(newAddress string, knownAddresses []string) bool {
  for _,v := range knownAddresses {
    if v == newAddress {
      return true
    }
  }

  return false
}
