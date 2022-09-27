package main

import (
	"bookserver/config"
	"bookserver/message"
	"bookserver/table"
	"bookserver/webserver"
	"context"
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
		message.Printf("Start Server\n")
		if err := s.Srv.Serve((s.L)); err != nil {
			return err
		}
		return nil
	})
	<-ctx.Done()
	if err := s.Srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	message.Println("shutdown")
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
	defer sql.Close()

	// SQLの読み込みテスト
	// data, err := sql.ReadAll(table.Copyfile)
	// if err != nil {

	// }
	// jsondata, _ := json.Marshal(data)
	// fmt.Println(string(jsondata))
	// // SQLの読み込みテスト2
	// v := map[string]interface{}{}
	// v["name"] = "bakatesu"
	// data, err := sql.Search(table.Booknames, v)
	// if err != nil {

	// }
	// jsondata, _ := json.Marshal(data)
	// fmt.Println(string(jsondata))
	// // SQLの読み込みテスト3
	// v := map[string]string{
	// 	"id": "1",
	// }
	// data, err := sql.SearchV2(table.Filelists, v)
	// if err != nil {

	// }
	// t, _ := table.FilelistCovert(data)
	// jsondata, _ := json.Marshal(t)
	// fmt.Println(string(jsondata))

	// //SQLの追加テスト
	// v := map[string]string{
	// 	"id":    "",
	// 	"name":  "test",
	// 	"title": "test",
	// }
	// err1 := sql.Add(table.Booknames, v)
	// if err1 != nil {
	// 	log.Println(err1.Error())
	// }

	ctx := context.Background()
	if err := Run(s, ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	message.Println("program end")

}
