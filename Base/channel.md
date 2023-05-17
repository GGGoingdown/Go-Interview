# Channel

### Ticker
```go
func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for {
			if v, ok := <-ch; ok {
				log.Printf("val:%d\n", v)
			}
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 30; i++ {
		select {
		// how to make sure all number send to ch channel?
		case ch <- i:
		case <-tick.C:
			log.Printf("%d: case <-tick.C\n", i)
		}

		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
```


### Timeout issue
```go
func job1(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	done <- true
}

func job2(done chan bool) {
	time.Sleep(600 * time.Millisecond)
	done <- true
}

func doJob(wg *sync.WaitGroup, f func(chan bool)) error {
	done := make(chan bool)
	defer wg.Done()
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(500 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		if err := doJob(wg, job1); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job2); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job1); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job2); err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

```


### Worker pool
```go
const concurrency = 3

func main() {
	// put tasks on channel
	tasks := make(chan int, 100)
	go func() {
		for j := 1; j <= 9; j++ {
			tasks <- j
		}
		close(tasks)
	}()

	// waitgroup, and close results channel when work done
	results := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(concurrency)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= concurrency; i++ {
		go func(id int) {
			defer wg.Done()

			for t := range tasks {
				fmt.Println("worker", id, "processing job", t)
				results <- t * 2
				time.Sleep(time.Second)
			}
		}(i)
	}

	// loop over results until closed (see above)
	for r := range results {
		fmt.Println("result", r)
	}
}
```

# Buffer and Unbuffer channel
```go
func worker(jobChan <-chan int) {
	for job := range jobChan {
		fmt.Println("current job:", job)
		time.Sleep(3 * time.Second)
		fmt.Println("finished job:", job)
	}
}

func enqueue(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}

func main() {
	// make a channel with a capacity of 1.
	jobChan := make(chan int, 1)

	// start the worker
	go worker(jobChan)

	fmt.Println(enqueue(1, jobChan)) // true
	fmt.Println(enqueue(2, jobChan)) // true
	fmt.Println(enqueue(3, jobChan)) // false

	fmt.Println("waiting the jobs")
	time.Sleep(10 * time.Second)
}

```

# Share by memory or Sh