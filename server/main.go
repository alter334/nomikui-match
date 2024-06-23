package main

import (
	"log"
	"nomikuimatch/pkg/setup"
)

func main() {
	_ = setup.DBsetup()
	log.Println(":::Nomikui-backend Setuped:::")
	setup.Echosetup()

}
