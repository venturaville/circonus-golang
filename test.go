package main

import (
  "github.com/venturaville/circonus-golang/circonus"
)

func main() {
  c := circonus.NewCirconus(&circonus.Circonus{})
  c.ListBroker()
//  c.ListTag()
}
