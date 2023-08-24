package main

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
)

func main() {
  err := common.Refactor("oldString", "newString", "*.txt", "*.json")
  if err != nil {
    println(err) 
  }
}