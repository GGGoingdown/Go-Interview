# Context Manager

### Context with cancel

```go
func doSomeWork(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d received close signal\n", num)
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doSomeWork(ctx, 1)
	go doSomeWork(ctx, 2)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

```


### Context with deadline
```go
func doSomeWork(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d received close signal\n", num)
			return
		default:
		}
	}
}

func main() {
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	go doSomeWork(ctx, 1)
	go doSomeWork(ctx, 2)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

```