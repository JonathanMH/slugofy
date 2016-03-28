package slugofy

//import "encoding/json"
import "strings"

func Hello() string {
  a := "hello"
  return a
}

func ReplaceSpaces(input string) string {
  output := strings.Replace(input, " ", "-", -1)
  return output
}

func Create(input string) string {
  output := ReplaceSpaces( input )
  return output;
}
