# Select

Go的select可以讀取channel的資料，如有多個case的時候是隨機讀取的。
另外如果select讀取的channel is blocking，select本身也會blocked，為了不要block住，我們可以使用`default`
來讓select執行預設的程式。

**Example without default case:**
```go
func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("Start send data to channel after 1 second")
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	select {
	case <-ch1: // block here until data is received from channel
		fmt.Println("Received data from channel")
	}

}
```
**Example with default case:**
```go
func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("Start send data to channel after 1 second")
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	select {
	case <-ch1:
		fmt.Println("Received data from channel")
	default: // always execute default case in this code
		fmt.Println("default case")
	}

}
```