package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()
	config, err := NewConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := dbConn(*config)
	if err != nil {
		log.Println(err)
		return
	}

	app := NewApplication(db, config)

	app.run()
}

func dbConn(cfg config) (*sql.DB, error) {
	fmt.Println(cfg.Server.Db.Dsn)
	db, err := sql.Open("postgres", cfg.Server.Db.Dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.Server.Db.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Server.Db.MaxIdleConns)
	duration, err := time.ParseDuration(cfg.Server.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
