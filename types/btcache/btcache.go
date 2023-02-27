package btcache

import (
	"sort"
	"time"

	"github.com/lcyvin/go-neewer/utils"
	"tinygo.org/x/bluetooth"
)

type ScanCache struct {
  Devices map[bluetooth.Addresser]*bluetooth.ScanResult
  DeviceTime map[time.Time]bluetooth.Addresser
  TTL time.Duration
}

func NewScanCache(ttl int) *ScanCache {
  sc := &ScanCache{}

  sc.TTL = time.Duration(ttl) * time.Second
  sc.DeviceTime = make(map[time.Time]bluetooth.Addresser)
  sc.Devices = make(map[bluetooth.Addresser]*bluetooth.ScanResult)

  return sc
}

func (sc *ScanCache) AddDevice(dev *bluetooth.ScanResult) (*ScanCache) {
  sc = sc.PruneCache()

  t := time.Now()

  exists := sc.GetDevice(dev.Address)

  if exists == nil {
    sc.Devices[dev.Address] = dev
  }

  // insert a new cache entry regardless of if it has been scanned again or is new
  sc.DeviceTime[t] = dev.Address

  return sc
}

// if dev doesn't exist, device should be nil and can be tested against in that manner
func (sc *ScanCache) GetDevice(address bluetooth.Addresser) (*bluetooth.ScanResult) {
  sc = sc.PruneCache()

  if dev,ok := sc.Devices[address]; ok != false {
    return dev
  }

  return nil
}

func (sc *ScanCache) PruneCache() *ScanCache {
  if len(sc.DeviceTime) == 0 {
    return sc
  }
  prunedCacheTimers := make(map[time.Time]bluetooth.Addresser)
  prunedCacheDevices := make(map[bluetooth.Addresser]*bluetooth.ScanResult)

  cacheTimeKeys := make([]time.Time, 0, len(sc.DeviceTime))
  for t := range sc.DeviceTime {
    cacheTimeKeys = append(cacheTimeKeys, t)
  }
  sort.Slice(cacheTimeKeys, func(i, j int) bool {
    return cacheTimeKeys[i].Before(cacheTimeKeys[j])
  })

  idx := utils.SeekTimeArrayIdxGT(cacheTimeKeys, sc.TTL, len(cacheTimeKeys)-1, 0)
  if idx != -1 {
    for _,v := range cacheTimeKeys[idx+1:] {
      devAddr := sc.DeviceTime[v]
      dev := sc.Devices[devAddr]

      prunedCacheTimers[v] = devAddr
      prunedCacheDevices[devAddr] = dev
    }

    sc.Devices = prunedCacheDevices
    sc.DeviceTime = prunedCacheTimers
  }

  return sc
}

func (sc *ScanCache) ListDevices() []*bluetooth.ScanResult {
  sc = sc.PruneCache()

  results := make([]*bluetooth.ScanResult, 0)
  for _,v := range sc.Devices {
    results = append(results, v)
  }

  return results
}
