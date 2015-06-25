package main

import (
  "fmt"
  "github.com/venturaville/circonus-golang/circonus"
)

func main() {
  c := circonus.NewCirconus(&circonus.Circonus{})
  data := c.ListBroker()
  fmt.Println("Brokers:")
  for _,element := range data {
    fmt.Println(element.(map[string]interface{})["_name"])
  }
//  fmt.Println("%v", data)
//  entry := data[0]
//  fmt.Println("%v", entry)
//  fmt.Println("%v", entry.(map[string]interface{})["_name"])

//  c.ListTag()
}
