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
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	rps          int
	clientsCount int
)

func main() {
	flag.IntVar(&rps, "rps", 300, "rps")
	flag.IntVar(&clientsCount, "clients", 2, "clients count")

	flag.Parse()

	if err := run(context.Background()); err != nil {
		log.Println(err)

		os.Exit(1)
	}

	os.Exit(0)
}

func run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < clientsCount; i++ {
		g.Go(startClient)
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func startClient() error {
	ch := make(chan struct{}, rps)
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			ch <- struct{}{}
		}
	}()

	for range ch {
		to := rand.Int31()
		go callGet(
			time.Now().Add(time.Duration(to)),
			time.Now().Add(time.Duration(rand.Int31n(to))),
		)

		//go callPost(
		//	time.Now().Add(time.Duration(to)),
		//	int64(rand.Int31n(to)),
		//)
	}

	return nil
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

type getArgs struct {
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

func callGet(dateFrom, dateTo time.Time) error {
	getArgs := getArgs{
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}

	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(&getArgs)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/stat", body)
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
