package main

import (
	"bookserver/webserver"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
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

func Setup() (*webserver.Server, error) {
	cfg, err := webserver.NewSetup()
	if err != nil {
		return nil, err
	}
	h, err := setupbaseRoute()
	if err != nil {
		return nil, err
	}
	h.setupRoute(cfg)
	return cfg.NewServer()

}

func main() {
	s, err := Setup()
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	ctx := context.Background()
	if err := Run(s, ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("end")

}
