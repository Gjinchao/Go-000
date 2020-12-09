package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "ping")
}

func main() {
	sigC := make(chan os.Signal, 1)
	defer close(sigC)
	signal.Notify(sigC)
	//facade.ServeHttp(":8888", sigC)
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	s := &http.Server{
		Addr:    ":8080",
		Handler: myHandler{},
	}
	g.Go(func() error {
		defer cancel()
		return s.ListenAndServe()
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case sig := <-sigC:
			switch sig {
			case os.Interrupt:
				return s.Shutdown(ctx)
			case syscall.SIGKILL:
				os.Exit(-1)
				return nil
			default:
				return nil
			}
		}
	})
	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
