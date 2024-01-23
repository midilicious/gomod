package main

import (
  "fmt"
  "log"
  "os"
)

func total(rwx string) {
  acc := 0
  for i := range rwx {
    if string([]rune(rwx)[i]) == "r" { acc += 4 }
    if string([]rune(rwx)[i]) == "w" { acc += 2 }
    if string([]rune(rwx)[i]) == "x" { acc += 1 }
  }
  // return acc int
}

func convert(eperm string) {
  fmt.Printf("%s\t", eperm)
  dir := string([]rune(eperm)[0])
  fmt.Printf("%s\t", dir)
  usr := string([]rune(eperm)[1:4])
  fmt.Printf("%s\t", usr)
  // ousr := total(usr)
  total(usr)
  // fmt.Println(ousr)
  // TODO figure out how to return values from the total function without 
  // getting a "too many return values" error
  group := string([]rune(eperm)[4:7])
  fmt.Printf("%s\t", group)
  everyone := string([]rune(eperm)[7:10])
  fmt.Printf("%s\t", everyone)
  var operm [4]string 
  if dir == "-" { operm[0] = "0" }
  fmt.Printf("%s", operm[:])
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
