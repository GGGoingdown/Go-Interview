# Goroutine

### Race condition
```go
var total = 0

func main() {
	count := 10
	wg := &sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				total += 1
			}
		}()
	}

	wg.Wait()
	fmt.Println(total)
}
```


### Goroutine leak
```go
func test() chan int {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	return ch
}

func main() {
	for i := 0; i < 4; i++ {
		test()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

```