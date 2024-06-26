package main

import (
	"log"
	setup "nomikuimatch/setup"
)

func main() {
	_db := setup.DBsetup()
	log.Println(":::Nomikui-DB Setuped:::")
	setup.Echosetup(_db)
	log.Println(":::Nomikui-Echo Setuped:::")
	log.Println(":::Nomikui-Backend Setuped:::")

}
