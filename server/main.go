package main

import (
	"log"
	"nomikuimatch/setup"
)

func main() {
	_ = setup.DBsetup()
	log.Println(":::Nomikui-backend Setuped:::")
	

}
