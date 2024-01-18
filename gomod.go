package main

import (
  "fmt"
  "log"
  "os"
)

func total(rwx string) {
  var acc, i int
  var oct string
  acc = 0
  i = 1
  oct = "0"
  for i <= 3 {
    if string(rwx[i]) == "r" { acc =+ 4 }
    if string(rwx[i]) == "w" { acc =+ 2 }
    if string(rwx[i]) == "x" { acc =+ 1 }
    i += 1
    oct += string(acc)

    // TODO I don't know, maybe build a list for oct then write it into a string?
  }
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
