package main
import "fmt"

func FirstFactorial(num int) int {
  sum := 1
  for ; num > 0; num-- {
    sum = sum * num
  }
  return sum
}

func main() {
  fmt.Println(FirstFactorial(3), " || The correct answer is 6.")
  fmt.Println(FirstFactorial(4), " || The correct answer is 24.")
  fmt.Println(FirstFactorial(5), " || The correct answer is 120.")
  fmt.Println(FirstFactorial(6), " || The correct answer is 720.")
  fmt.Println(FirstFactorial(7), " || The correct answer is 5040.")
  fmt.Println(FirstFactorial(8), " || The correct answer is 40320.")
  fmt.Println(FirstFactorial(9), " || The correct answer is 362880.")
  fmt.Println(FirstFactorial(10), " || The correct answer is 3628800.")
}
