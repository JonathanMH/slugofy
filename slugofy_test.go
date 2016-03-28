package slugofy_test

import "fmt"
import "testing"
import "."

func TestHello(t *testing.T){
  expected := "hello"
  actual := slugofy.Hello()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestReplaceSpaces(t *testing.T){
  expected := "Hello-World!"
  actual := slugofy.ReplaceSpaces("Hello World!")
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got: '%s'\n", expected, actual)
  } else {
    fmt.Printf("Test passed, expexted: '%s', got: '%s'\n", expected, actual)
  }
}

func TestCreate(t *testing.T){
  expected := "Jonathan-M-Hethey"
  actual := slugofy.Create("Jonathan M Hethey")
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got: '%s'\n", expected, actual)
  } else {
    fmt.Printf("Test passed, expexted: '%s', got: '%s'\n", expected, actual)
  }
}
