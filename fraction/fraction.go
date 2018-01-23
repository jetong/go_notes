package fraction	// used in interface1.go

import (
  "fmt"
  "math"
  "github.com/jetong/go_pract/gcd"
)

type Fraction struct {
  Num int
  Den int
}

// implement String() method of Stringer interface
func (f Fraction) String() string {
  f.Reduce()
  // have the negative sign print in front of Num rather than Den
  if f.Den < 0 {
    f.Num = -f.Num
    f.Den = -f.Den
  }
  // case: whole number
  if math.Mod(float64(f.Num), float64(f.Den)) == 0 {
    return fmt.Sprintf("%v", f.Num/f.Den)
  // case: normal, proper fraction
  }else if math.Abs(float64(f.Num)) < math.Abs(float64(f.Den)) {
    return fmt.Sprintf("%v/%v", f.Num, f.Den)
  }else{
  // case: print fraction with whole part
    whole := f.Num/f.Den
    f.Num = int(math.Mod(float64(f.Num), float64(f.Den)))
    if whole < 0 {
      return fmt.Sprintf("%v & %v/%v", whole, math.Abs(float64(f.Num)), math.Abs(float64(f.Den)))
    }else{
      return fmt.Sprintf("%v & %v/%v", whole, f.Num, f.Den)
    }
  }
}

func (f1 Fraction) Add(f2 Fraction) Fraction {
  if f1.Den == f2.Den {
    return Fraction{f1.Num + f2.Num, f1.Den}
  }else{
    return Fraction{f1.Num*f2.Den + f2.Num*f1.Den, f1.Den*f2.Den}
  }
}

func (f *Fraction) Reduce() {
  gcd := gcd.GCDEuclidean(f.Num, f.Den)
  f.Num /= gcd
  f.Den /= gcd
}
