package main

import (
  "fmt"
  "log"
  "os"
)

func total(rwx string) {
  var acc int
  acc = 0
  if string(rwx[0]) == "r" { acc =+ 4 }
  if string(rwx[1]) == "w" { acc =+ 2 }
  if string(rwx[2]) == "x" { acc =+ 1 }
  oct := string(acc)
  fmt.Printf(oct)
//  return oct
}

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
  var operm [4]string 
  if dir == "-" { operm[0] = "0" }
//  operm[1] = total(user)
  fmt.Printf("%s", operm[:])
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
