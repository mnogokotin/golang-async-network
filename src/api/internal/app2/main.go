package app2

import (
	"context"
	"fmt"
	"github.com/mnogokotin/golang-async-network/pkg/utils/async"
	"sync"
	"time"
)

func Run() {
	fetcher1 := make(chan interface{})
	fetcher2 := make(chan interface{})

	ctx, cancelCtx := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancelCtx()
	combinedFetcher := fanIn(ctx, fetcher1, fetcher2)

	go async.Receive(combinedFetcher)
	go async.Send(fetcher1, "f1", 500*time.Millisecond)
	go async.Send(fetcher1, "f2", 0)

	time.Sleep(2 * time.Second)
}

func fanIn(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {
	combinedFetcher := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(len(fetchers))

	for _, f := range fetchers {
		f := f
		go func() {
			defer wg.Done()
			for {
				select {
				case output := <-f:
					combinedFetcher <- output
				case <-ctx.Done():
					fmt.Println("fanIn routine done")
					return
				}
			}
		}()
	}

	// Channel cleanup
	go func() {
		wg.Wait()
		close(combinedFetcher)
	}()
	return combinedFetcher
}
