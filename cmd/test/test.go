package main

import (
  "fmt"
  "regexp"
)

func main(){
  tag := "(1234, 2345)"
  exp := regexp.MustCompile(`^ *\(([0-9a-fA-F]+), *([0-9a-fA-F]+)\) *$`)
  submatches := exp.FindStringSubmatch(tag)

  fmt.Println(submatches[1])
  fmt.Println(submatches[2])
}

