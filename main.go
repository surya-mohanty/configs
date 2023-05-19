package main

import (
	"log"

	"github.com/surya-mohanty/configs/dbconfigs/mongoconfig"
)

func main() {

	mongoClient, err := mongoconfig.GetConnection()
	if err != nil {
		log.Panic(err)
	} else {
		log.Panic(mongoClient)
	}
}
