package main
import "fmt"

func SimpleAdding(num int) int {
  sum := 0
  for i:= 1; i <= num; i++ {
    sum = sum + i
  }
  return sum

}

func main() {

  fmt.Println(SimpleAdding(45), "  || The correct answer is 1035.")
  fmt.Println(SimpleAdding(13), "  || The correct answer is 91.")
  fmt.Println(SimpleAdding(2), "  || The correct answer is 3.")
  fmt.Println(SimpleAdding(5), "  || The correct answer is 15.")
  fmt.Println(SimpleAdding(156), "  || The correct answer is 12246.")
  fmt.Println(SimpleAdding(999), "  || The correct answer is 499500.")
  fmt.Println(SimpleAdding(67), "  || The correct answer is 2278.")
  fmt.Println(SimpleAdding(123), "  || The correct answer is 7626.")
  fmt.Println(SimpleAdding(9), "  || The correct answer is 45.")
  fmt.Println(SimpleAdding(10), "  || The correct answer is 55.")

}
