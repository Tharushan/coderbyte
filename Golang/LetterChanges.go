package main
import (
  "fmt"
  "unicode"
  "bytes"
  "strings"
)

func LetterChanges(str string) string {
  var buffer bytes.Buffer
  for i := 0; i < len(str); i++ {
    var letter string
    if unicode.IsLetter(rune(str[i])) {
      if (string(str[i])) == "z" {
        letter = "a"
      } else if string(str[i]) == "Z" {
        letter = "A"
      } else {
        letter = string(str[i] + 1)
      }
    } else {
      letter = string(str[i])
    }
    if strings.Contains("aeiou", letter) {
      letter = strings.ToUpper(letter)
    }
    buffer.WriteString(letter)
  }
  return buffer.String()

}

func main() {

  fmt.Println(LetterChanges("abcdz"), "The correct answer is bcdEA")
  fmt.Println(LetterChanges("hello world"), "The correct answer is Ifmmp xpsmE")
  fmt.Println(LetterChanges("sentence"), "The correct answer is tfOUfOdf")
  fmt.Println(LetterChanges("replace!*"), "The correct answer is sfqmbdf!*")
  fmt.Println(LetterChanges("coderbyte"), "The correct answer is dpEfsczUf")
  fmt.Println(LetterChanges("beautiful^"), "The correct answer is cfbvUjgvm^")
  fmt.Println(LetterChanges("oxford"), "The correct answer is pygpsE")
  fmt.Println(LetterChanges("123456789ae"), "The correct answer is 123456789bf")
  fmt.Println(LetterChanges("this long cake@&"), "The correct answer is UIjt mpOh dblf@&")
  fmt.Println(LetterChanges("a b c dee"), "The correct answer is b c d Eff")
  fmt.Println(LetterChanges("a confusing /:sentence:/[ this is not!!!!!!!~"), "The correct answer is b dpOgvtjOh /:tfOUfOdf:/[ UIjt jt OpU!!!!!!!~")

}
