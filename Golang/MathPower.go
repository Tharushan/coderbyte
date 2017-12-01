package main

import (
  "fmt"
)

func MathPower(base int, exp int) float64 {
  sum := 1
  expPositive := true
  if exp < 0 {
    expPositive = false
    exp = -1 * exp
  }
  for i:= 0; i < exp; i++ {
    sum = sum * base
  }
  if expPositive == false {
    return float64(float64(1) / float64(sum))
  }
  return float64(sum)
}

func main() {
  fmt.Println(MathPower(10, 0))
  fmt.Println(MathPower(10, 1))
  fmt.Println(MathPower(10, 2))
  fmt.Println(MathPower(2, 1))
  fmt.Println(MathPower(2, 2))
  fmt.Println(MathPower(2, 3))
  fmt.Println(MathPower(2, 4))
  fmt.Println(MathPower(2, 5))
  fmt.Println(MathPower(10, -1))
  fmt.Println(MathPower(2, -2))
  fmt.Println(MathPower(2, -3))
}
