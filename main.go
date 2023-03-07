package main

import (
	"flag"
	config "go-gorm-hierarchy/configs"
	"go-gorm-hierarchy/db"
	"go-gorm-hierarchy/services/chart"
)

func main() {
	configs := flag.String("config", "configs", "set configs path, default as: 'configs'")
	flag.Parse()

	err := config.InitConfig(*configs)
	if err != nil {
		panic(err)
	}

	if config.CF.Database.Enable {
		err = db.New(config.CF.Database)
		if err != nil {
			panic(err)
		}
	}

	srv := chart.NewService()
	_, err = srv.GetAll()
	if err != nil {
		panic(err)
	}
}
