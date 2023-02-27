package bluetooth

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"tinygo.org/x/bluetooth"
  "github.com/lcyvin/go-neewer/types/btcache"
)

var adapter = bluetooth.DefaultAdapter
var sc = btcache.NewScanCache(45)

func Scan(matchPatterns []*regexp.Regexp, knownDevices []string, timeout int, exitAfterMatch bool) ([]*bluetooth.ScanResult, error) {
  devs := make([]*bluetooth.ScanResult, 0)

  err := adapter.Enable()
  if err != nil {
    log.Fatalf("could not enable bluetooth adapter: %v", err)
  }

  timer := time.NewTicker(time.Duration(timeout) * time.Second)
  results := make(chan bluetooth.ScanResult)

  go ScanForLight(adapter, results)

  for {
    select {
    case <- timer.C:
      if len(devs) != 0 {
        return devs, nil
      }

      return devs, fmt.Errorf("Scanner timed out after %d seconds", timeout)

    case dev := <- results:
      sc.AddDevice(&dev)
      for _,matchStr := range matchPatterns {
        if matchStr.MatchString(dev.LocalName()) {
        devs = append(devs, &dev)
        if exitAfterMatch {
          return devs, nil
          }
        }
      }
    }
  }
}

func ScanForLight(adapter *bluetooth.Adapter, resChan chan bluetooth.ScanResult) {
  // start our scan
  adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
    resChan <- device
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

func ScanForMacAddress(address string, timeout int) (*bluetooth.ScanResult, error) {
  // check if the device is in the cache
  mac, err := bluetooth.ParseMAC(address)
  if err != nil {
    log.Fatalf("Could not parse mac address %s: %v", address, err)
  }

  dev := sc.GetDevice(bluetooth.Address{
    MACAddress: bluetooth.MACAddress{MAC: mac},  
  })

  if dev != nil {
    return dev, nil
  }

  timeoutDuration := time.Duration(timeout) * time.Second
  ticker := time.NewTicker(timeoutDuration)

  result := make(chan bluetooth.ScanResult)

  go ScanForLight(adapter, result)

  for {
    select {
      case <- ticker.C:
        err := adapter.StopScan()
        if err != nil {
          log.Fatal(err)
        }
        return &bluetooth.ScanResult{}, fmt.Errorf("Scan timed out after %d seconds", timeout)
      case dev,ok := <- result:
        if !ok {
          log.Fatalf("got not-okay response")
        }
        fmt.Println(dev.Address.String())
        if dev.Address.String() == address {
          err := adapter.StopScan()
          if err != nil {
            log.Fatal(err)
          }
          return &dev, nil    
        }
    }
  }
}
