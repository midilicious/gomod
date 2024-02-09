package main

import (
  "flag"
  "fmt"
  "log"
  "os"
)

var (
	recursive = flag.Bool("recursive",false,"appends changes to all subfolders")
)

func total(rwx string) string {
  acc := 0
  for i := range rwx {
    if string([]rune(rwx)[i]) == "r" { acc += 4 }
    if string([]rune(rwx)[i]) == "w" { acc += 2 }
    if string([]rune(rwx)[i]) == "x" { acc += 1 }
  }
  return fmt.Sprint(acc)
}

func convert(eperm string) {
  fmt.Printf("%s\t", eperm)
  dir := string([]rune(eperm)[0])
  fmt.Printf("%s\t", dir)
  usr := string([]rune(eperm)[1:4])
  fmt.Printf("%s\t", usr)
  group := string([]rune(eperm)[4:7])
  fmt.Printf("%s\t", group)
  everyone := string([]rune(eperm)[7:10])
  fmt.Printf("%s\t", everyone)
  var operm string 
  if dir == "-" {
    operm = fmt.Sprintf("0%s%s%s", total(usr), total(group), total(everyone) ) 
  }
  fmt.Println(operm)
}

func main() {
  flag.Parse()
  // Print flag 
  // fmt.Printf("recursive status:%t", *recursive)
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
