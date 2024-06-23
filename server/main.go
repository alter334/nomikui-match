package main

import (
	"log"
	setup "nomikuimatch/pkg/_setup"
)

func main() {
	_ = setup.DBsetup()
	log.Println(":::Nomikui-DB Setuped:::")
	setup.Echosetup()
	log.Println(":::Nomikui-Echo Setuped:::")
	log.Println(":::Nomikui-Backend Setuped:::")

}
