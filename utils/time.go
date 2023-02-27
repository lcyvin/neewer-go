package utils

import (
  "time"
)


// yes I know we don't need to check neighbors here but it makes my brain happy
func SeekTimeArrayIdxGT(keys []time.Time, duration time.Duration, high int, low int) int {
  var mid int
  idx :=  -1

  if low == 0 {
    mid = high/2
  }

  if high - low == 1 {
    if time.Since(keys[high]) >= duration {
      return high
    }

    if time.Since(keys[low]) >= duration {
      return low
    }
  }
  
  if time.Since(keys[low]) < duration {
    return -1
  }

  if time.Since(keys[mid]) < duration {
    if time.Since(keys[mid-1]) > duration {
      return mid-1
    }

    idx = SeekTimeArrayIdxGT(keys, duration, mid, low)
  }

  if time.Since(keys[mid]) > duration {
    if time.Since(keys[mid+1]) < duration {
      return mid
    }

    idx = SeekTimeArrayIdxGT(keys, duration, high, mid)
  }

  return idx
}
