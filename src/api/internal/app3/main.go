package app3

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Run() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	subscription := NewSubscription(ctx, NewFetcher("example.com"), 1)

	time.AfterFunc(3*time.Second, func() {
		cancelCtx()
		fmt.Println("context canceled")
	})

	for item := range subscription.Updates() {
		fmt.Println(item)
	}
}

func (s *sub) serve(ctx context.Context, checkFrequency int) {
	clock := time.NewTicker(time.Duration(checkFrequency) * time.Second)
	type fetchResult struct {
		fetched item
		err     error
	}
	fetchDone := make(chan fetchResult, 1)

	for {
		select {
		case <-clock.C:
			go func() {
				fetched, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, err}
			}()
		case result := <-fetchDone:
			fetched := result.fetched
			if result.err != nil {
				log.Printf("Fetch error: %v \n Waiting the next iteration\n", result.err.Error())
				break
			}
			s.updates <- fetched
		case <-ctx.Done():
			return
		}
	}
}

type Subscription interface {
	Updates() <-chan item
}

type Fetcher interface {
	Fetch() (item, error)
}

type item struct {
	id int
}

type fetcher struct {
	uri string
}

type sub struct {
	fetcher Fetcher
	updates chan item
}

func NewSubscription(ctx context.Context, fetcher Fetcher, freq int) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan item),
	}
	go s.serve(ctx, freq)
	return s
}

func (s *sub) Updates() <-chan item {
	return s.updates
}

func NewFetcher(uri string) Fetcher {
	f := &fetcher{
		uri: uri,
	}
	return f
}

func (f *fetcher) Fetch() (item, error) {
	i := item{id: rand.Intn(100)}
	return i, nil
}
