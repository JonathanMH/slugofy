//package slugofy
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// read all the json files from the dependency slugify
func readJsonFiles() map[string]string {
    var pairs map[string]string
    files, err := ioutil.ReadDir("slugify/Resources/rules")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
//        fmt.Println(file.Name())
        fileData, err1 := ioutil.ReadFile("slugify/Resources/rules/" + file.Name())
        if err1 == nil {
//            fmt.Printf("%+v\n", err1)
        } else {
//            fmt.Println("error:", err1)
        }
        //    err := json.Unmarshal(fileData, &pairs)
        err := json.NewDecoder( strings.NewReader( string( fileData ) ) ).Decode( &pairs )
        if err == nil {
//            fmt.Printf( "%+v\n", err )
        } else {
//            fmt.Println( "error:", err )
        }
        //    fmt.Println(file.Name())
    }

    return pairs
}

// first replace spaces with dashes, perhaps extend to all non
// alphanumeric characters
func ReplaceSpaces(input string) string {
  output := strings.Replace(input, " ", "-", -1)
  return output
}

// create a slug of an input string
func Create(input string) string {
    pairs := readJsonFiles()
//    fmt.Printf("pairs: %v\n", pairs)
    tempString := ReplaceSpaces( input )
    var output string
    for _, character := range tempString {
//        fmt.Println( index, " => ", string( character ) )
//        fmt.Println( pairs[ string( character ) ] )
        _, present := pairs[ string( character ) ]
        if present {
            output += pairs[ string( character ) ]
        } else {
            output += string( character )
        }
    }
    return output;
}

func main(){
    foobar := Create( "my test string ö" )
    snafu := Create( "some more öäüß ææåøeé" )
    fmt.Println( foobar );
    fmt.Println( snafu );
}
