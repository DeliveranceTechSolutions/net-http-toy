package main


import (
	"context"
	"fmt"
	"log"
//	"net/http"
	// "runtime"
	"sync"
	"time"

	_ "net/http/pprof"
)

func count(
	ctx context.Context,  
	wg *sync.WaitGroup, 
	ch chan int, 
	seed int, 
	start time.Time, 
	target int,
) {
	defer wg.Done()
	for i := seed; i <= target; {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		default:
			ch<-i
			i += 2
		}
	}
	endTime := time.Now()
	executionTime := endTime.Sub(start)
	fmt.Println(executionTime)
}

func concurrency() {
	var wg sync.WaitGroup
	timeout := time.Millisecond * 1
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	odd := make(chan int)
	even := make(chan int)
	oddSeed := 1
	evenSeed := 2

	target := 500000

	start := time.Now()
	wg.Add(2)
	go count(
		ctx, 
		&wg, 
		odd, 
		evenSeed,
		start,
		target, 
	)
	go count(
		ctx, 
		&wg, 
		odd, 
		oddSeed,
		start,
		target, 
	)
	go func() {
		wg.Wait()
		close(even)
		close(odd)
	}()
	
	free: for {
		select {
		case o := <-odd:
			fmt.Println(o)
		case e := <-even:
			fmt.Println(e)
		case <-ctx.Done():
			break free
		}
	}
	
	// if channels were used and needed to return in order
	//for o := range odd {
	//	fmt.Println(o)
	//	for e := range even {
	//		fmt.Println(e)
	//		break
	//	}
	//}

	fmt.Println("returning here")

}

type Message struct {
	payload string
}

const (
        SERVER_HOST = "localhost"
        SERVER_PORT = "9988"
        SERVER_TYPE = "tcp"
)

func main() {
//		runtime.GOMAXPROCS(8)
//		concurrency()
	// Server for pprof
	// http.HandleFunc("/concurrency", func(http.ResponseWriter, *http.Request) {	concurrency()	})
	// fmt.Println(http.ListenAndServe("localhost:6060", nil))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := NewSocketConn(SERVER_TYPE, SERVER_HOST, SERVER_PORT)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	defer s.Disconnect(ctx)

	s.Subscribe(ctx)
	for {
		var msg string
		fmt.Scanln(&msg)

		if msg == "exit" {
			break
		}

		err = s.Publish(ctx, msg)
		if err != nil {
			log.Println(fmt.Errorf("Error: Couldn't publish message to subscriber %s", err.Error()))
		}
	}

	return
}
