package main

import (
  "fmt"
  "os"
  "github.com/venturaville/circonus-golang/circonus"
)

func main() {
  appname := os.Getenv("CIRCONUS_APPNAME")
  if appname == "" { appname = "curl" }
  apitoken := os.Getenv("CIRCONUS_APITOKEN")
  apiserver := os.Getenv("CIRCONUS_APISERVER")
  if appserver == "" { appname = "api.circonus.com" }
  fmt.Println("CIRCONUS_APITOKEN:", apitoken)
  fmt.Println("CIRCONUS_APISERVER:", apiserver)
  fmt.Println("CIRCONUS_APPNAME:", appname)

  c := &circonus.Circonus{appname,apiserver,apitoken}
  c.ListBroker()
  c.ListTag()
}
