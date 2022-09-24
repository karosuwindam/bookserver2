package main

import (
	"bookserver/config"
	"bookserver/table"
	"bookserver/webserver"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func Run(s *webserver.Server, ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("Start Server")
		if err := s.Srv.Serve((s.L)); err != nil {
			return err
		}
		return nil
	})
	<-ctx.Done()
	if err := s.Srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	log.Println("shutdown")
	return eg.Wait()
}

func Setup() (*table.Config, *webserver.Server, error) {
	//set up config
	cfgdata, err := config.Envread()
	if err != nil {
		return nil, nil, err
	}
	//set up DB
	sql, err := table.Setup(cfgdata)
	if err != nil {
		return nil, nil, err
	}

	//Set up webserver
	cfg, err := webserver.NewSetup(cfgdata)
	if err != nil {
		return nil, nil, err
	}
	h, err := setupbaseRoute()
	if err != nil {
		return nil, nil, err
	}
	err = h.setupdatabase(sql)
	if err != nil {
		return nil, nil, err
	}
	h.setupRoute(cfg)
	s, err := cfg.NewServer()
	return sql, s, err

}

func main() {
	sql, s, err := Setup()
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}

	// SQLの読み込みテスト
	// data, err := sql.ReadAll(table.Copyfile)
	// if err != nil {

	// }
	// jsondata, _ := json.Marshal(data)
	// fmt.Println(string(jsondata))
	// SQLの読み込みテスト2
	id := 1
	data, err := sql.Read(table.Copyfile, id)
	if err != nil {

	}
	jsondata, _ := json.Marshal(data)
	fmt.Println(string(jsondata))

	defer sql.Close()
	ctx := context.Background()
	if err := Run(s, ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("program end")

}
