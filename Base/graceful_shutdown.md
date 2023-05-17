# Graceful shutdown

以下展示兩種方法來做graceful shutdown
- after all job finished
- immediately

此兩種方法最大的差異性在於 `worker` 方法裡，當有接收到`ctrl + v`訊號時，`context` 會收到error訊號。

### Shutdown after all job finished
```go
package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Consumer struct
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func processJob(workerID, jobID int) {
	sleep := time.Duration(getRandomNum())
	log.Printf("worker %d received job %d need to process %d s\n", workerID, jobID, sleep)
	time.Sleep(sleep * time.Second)
	log.Printf("worker %d finished job %d\n", workerID, jobID)
}

func (c *Consumer) Enqueue(jobID int) {
	log.Printf("send job %d", jobID)
	c.inputChan <- jobID
}

func (c *Consumer) StartConsumer(con context.Context) {
	for {
		select {
		case job := <-c.inputChan:
			c.jobsChan <- job
		case <-con.Done():
			close(c.jobsChan)
			return
		}
	}
}

func (c *Consumer) Worker(workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("start worker %d", workerID)

	for job := range c.jobsChan {
		processJob(workerID, job)
	}

	log.Printf("stop worker %d", workerID)
}

func ContextCancelWithFunc(c context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(c)
	go func() {
		osC := make(chan os.Signal, 1)
		signal.Notify(osC, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(osC)

		select {
		case <-ctx.Done():
		case <-osC:
			log.Println("cancel from ctrl+c event")
			cancel()
			f()
		}
	}()
	return ctx
}

func main() {

	wg := sync.WaitGroup{}
	workerPool := 2
	wg.Add(workerPool)
	done := make(chan struct{})
	consumer := Consumer{
		inputChan: make(chan int, 5),
		jobsChan:  make(chan int, workerPool),
	}

	ctx := ContextCancelWithFunc(context.Background(), func() {
		wg.Wait()
		close(done)
	})

	go consumer.StartConsumer(ctx)

	for i := 1; i <= workerPool; i++ {
		go consumer.Worker(i, &wg)
	}

	for i := 0; i < 5; i++ {
		consumer.Enqueue(i)
	}
	<-done
}

```


### Shutdown immediately


```go
package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Consumer struct
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func processJob(workerID, jobID int) {
	sleep := time.Duration(getRandomNum())
	log.Printf("worker %d received job %d need to process %d s\n", workerID, jobID, sleep)
	time.Sleep(sleep * time.Second)
	log.Printf("worker %d finished job %d\n", workerID, jobID)
}

func (c *Consumer) Enqueue(jobID int) {
	log.Printf("send job %d", jobID)
	c.inputChan <- jobID
}

func (c *Consumer) StartConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.inputChan:
			c.jobsChan <- job
			if ctx.Err() != nil {
				close(c.jobsChan)
				return
			}
		case <-ctx.Done():
			close(c.jobsChan)
			return
		}
	}
}

func (c *Consumer) Worker(ctx context.Context, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("start worker %d", workerID)

	for {
		select {
		case job := <-c.jobsChan:
			if ctx.Err() != nil {
				log.Printf("get next job %d by worker %d", job, workerID)
				log.Printf("stop worker %d", workerID)
				return
			}
			processJob(workerID, job)
		case <-ctx.Done():
			log.Printf("stop worker %d", workerID)
			return
		}

	}
}

func ContextCancelWithFunc(c context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(c)
	go func() {
		osC := make(chan os.Signal, 1)
		signal.Notify(osC, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(osC)

		select {
		case <-ctx.Done():
		case <-osC:
			log.Println("cancel from ctrl+c event")
			cancel()
			f()
		}
	}()
	return ctx
}

func main() {

	wg := sync.WaitGroup{}
	workerPool := 2
	wg.Add(workerPool)
	done := make(chan struct{})
	consumer := Consumer{
		inputChan: make(chan int, 5),
		jobsChan:  make(chan int, workerPool),
	}

	ctx := ContextCancelWithFunc(context.Background(), func() {
		wg.Wait()
		close(done)
	})

	go consumer.StartConsumer(ctx)

	for i := 1; i <= workerPool; i++ {
		go consumer.Worker(ctx, i, &wg)
	}

	for i := 0; i < 5; i++ {
		consumer.Enqueue(i)
	}
	<-done
}

```