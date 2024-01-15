package main

import (
  "fmt"
  "log"
  "os"
)

func convert(eperm string) {
  fmt.Printf("%s\t", eperm)
  dir := string([]rune(eperm)[0])
  fmt.Printf("%s\t", dir)
  user := string([]rune(eperm)[1:4])
  fmt.Printf("%s\t", user)
  group := string([]rune(eperm)[4:7])
  fmt.Printf("%s\t", group)
  everyone := string([]rune(eperm)[7:10])
  fmt.Printf("%s\t", everyone)
//  fmt.Printf(eperms)
//  return operm
}

func main() {
  // stats
  fmt.Printf("Enter the file: ")
  var target string
  // next two lines for streamlining testing
  fmt.Printf("\n")
  target = "example.file"
  stats, err := os.Lstat(target)  
  if err != nil {
    log.Fatal(err)
  }
  fileMode := stats.Mode()
  unixPerms := fileMode & os.ModePerm
  eperm := fmt.Sprintf("%v", unixPerms)
  convert(eperm)
//  fmt.Printf("%s\t%s\t%s\n", string([]rune(seperm)[1:4]), eperm, "asdf"[1])
}
