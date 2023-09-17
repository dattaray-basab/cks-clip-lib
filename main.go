package main

import "github.com/dattaray-basab/cks-clip-lib/common"




func main() {
		err := common.PrependToFile("test.txt", "bye")
		if err != nil {
			println(err.Error())
		}
}