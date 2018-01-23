package gcd 

import "math"

func GCDEuclidean(a, b int) int {
  for b != 0 {
    t := b
    b = int(math.Mod(float64(a),float64(b)))
    a = t
  }
  return a
}
