package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func clearData() {
	fmt.Println("clearData")
	db, err := gorm.Open(sqlite.Open("/etc/x-ui/x-ui.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("UPDATE inbounds set up = 0,down=0;")
}

func main() {
	c := cron.New()
	c.AddFunc("@monthly", clearData)
	c.Start()
	select {}
}
