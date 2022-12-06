//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream) chan *Tweet {
	queue := make(chan *Tweet, 5)
	go func() {
		defer close(queue)

		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				break
			}
			queue <- tweet

		}
	}()
	return queue

}

func consumer(tweets chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	time.Sleep(time.Nanosecond)
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	wg := new(sync.WaitGroup)

	// Producer
	queue := producer(stream)

	// Consumer
	// add 3 consumer

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go consumer(queue, wg)
	}

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
