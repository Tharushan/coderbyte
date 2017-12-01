package main
import (
  "bytes"
  "fmt"
)

func FirstReverse(str string) string {
  var buffer bytes.Buffer
  for i := len(str) - 1; 0 <= i; i-- {
    buffer.WriteString(string(str[i]))
  }
  return buffer.String()

}

func main() {

  fmt.Println(FirstReverse("Coderbyte"), " || The correct answer is etybredoC")
  fmt.Println(FirstReverse("I Love Coding"), " || The correct answer is gnidoC evoL I")
  fmt.Println(FirstReverse("h333llLo"), " || The correct answer is oLll333h")
  fmt.Println(FirstReverse("Yo0"), " || The correct answer is 0oY")
  fmt.Println(FirstReverse("thisiscool"), " || The correct answer is loocsisiht")
  fmt.Println(FirstReverse("commacomma!"), " || The correct answer is !ammocammoc")
  fmt.Println(FirstReverse("123456789"), " || The correct answer is 987654321")
  fmt.Println(FirstReverse("lettersz!23z"), " || The correct answer is z32!zsrettel")
  fmt.Println(FirstReverse("aq"), " || The correct answer is qa")

}
