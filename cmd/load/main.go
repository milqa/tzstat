package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/milQA/tzstat/pkg/grsdown"
)

var (
	job          string
	rps          int
	clientsCount int
)

func main() {
	flag.StringVar(&job, "job", "get", "'get' or 'post' job")
	flag.IntVar(&rps, "rps", 50, "rps for client")
	flag.IntVar(&clientsCount, "clients", 10, "clients count")

	flag.Parse()

	grsdown.Run(context.Background(), func(ctx context.Context) error {
		switch job {
		case "get", "post":
		default:
			log.Println("invalid job: use 'get' or 'post'")

			return fmt.Errorf("invalid job: use 'get' or 'post'")
		}

		log.Println("start loader")

		if err := run(ctx); err != nil {
			log.Println(err)

			return err
		}

		return nil
	})
}

func run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < clientsCount; i++ {
		g.Go(startClient(ctx))
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func startClient(ctx context.Context) func() error {
	return func() error {
		ch := make(chan struct{}, rps)
		ticker := time.NewTicker(time.Second / time.Duration(rps))
		defer ticker.Stop()

		var f func()
		switch job {
		case "get":
			f = func() {
				to := rand.Int31()
				callGet(
					time.Now().Add(time.Duration(to)),
					time.Now().Add(time.Duration(rand.Int31n(to))),
				)
			}
		case "post":
			f = func() {
				to := rand.Int31()
				callPost(
					time.Now().Add(time.Duration(to)),
					int64(rand.Int31n(1000)),
				)
			}
		}

		go func() {
			for range ticker.C {
				ch <- struct{}{}
			}
		}()

		for range ch {
			select {
			case <-ch:
				go f()
			case <-ctx.Done():
				return nil
			}
		}

		return nil
	}
}

type postArgs struct {
	Date  time.Time `json:"datetime"`
	Value int64     `json:"value"`
}

func callPost(date time.Time, value int64) error {
	postArgs := postArgs{
		Date:  date,
		Value: value,
	}

	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(&postArgs)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/stat", body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("call post status code = %d", resp.StatusCode)
	}

	return nil
}

func callGet(dateFrom, dateTo time.Time) error {
	resp, err := http.Get("http://localhost:8080/api/stat")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("call post status code = %d", resp.StatusCode)
	}

	return nil
}
