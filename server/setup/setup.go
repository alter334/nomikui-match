package setup

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBsetup() *sqlx.DB{
	
	log.Println(time.Now(),":::Start DBsetup")

	NS := os.Getenv("NS")
	var user, pass, addr, dbname string

	if NS == "false" {
		user = os.Getenv("NS_MARIADB_USER")
		pass = os.Getenv("NS_MARIADB_PASSWORD")
		addr = os.Getenv("NS_MARIADB_HOSTNAME") + ":" + os.Getenv("NS_MARIADB_PORT")
		dbname = os.Getenv("NS_MARIADB_DATABASE")
	} else {
		user = os.Getenv("LOCAL_MYSQL_USER")
		pass = os.Getenv("LOCAL_MYSQL_PASSWORD")
		addr = os.Getenv("LOCAL_MYSQL_HOSTNAME") + ":" + os.Getenv("LOCAL_MYSQL_PORT")
		dbname = os.Getenv("LOCAL_MYSQL_DATABASE")
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("ERROR_DBsetup:",err)
	}

	conf := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err!=nil{
		log.Fatal("ERROR_DBsetup:",err)
	}
	log.Println(time.Now(),":::Complete DBsetup")

	return _db

}
