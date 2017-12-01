package main
import (
  "fmt"
  "regexp"
  "log"
  "strings"
)

func LongestWord(sen string) string {
  var longest string
  reg, err := regexp.Compile("[^a-zA-Z0-9]+")
  if err != nil {
    log.Fatal(err)
  }
  words := strings.Split(reg.ReplaceAllString(sen, " "), " ")
  for _, word := range words {
    if len(longest) < len(word) {
      longest = word
    }
  }

  return longest
}

func main() {
  fmt.Println(LongestWord("hello world"), " || The correct answer is hello")
  fmt.Println(LongestWord("this is some sort of sentence"), " || The correct answer is sentence")
  fmt.Println(LongestWord("longest word!!"), " || The correct answer is longest")
  fmt.Println(LongestWord("a beautiful sentence^&!"), " || The correct answer is beautiful")
  fmt.Println(LongestWord("oxford press"), " || The correct answer is oxford")
  fmt.Println(LongestWord("123456789 98765432"), " || The correct answer is 123456789")
  fmt.Println(LongestWord("letter after letter!!"), " || The correct answer is letter")
  fmt.Println(LongestWord("a b c dee"), " || The correct answer is dee")
  fmt.Println(LongestWord("a confusing /:sentence:/[ this is not!!!!!!!~"), " || The correct answer is confusing")

}
