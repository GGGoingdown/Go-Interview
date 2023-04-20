# Defer

**Explain:**
下面的程式碼會分別印出:
deferTest3
deferTest2
deferTest1
panic: panicTest

因為`defer `不受`panic`的限制，所以依然會執行。
`defer`執行的順序為FILO(`Stack`)
當程式遇到`defer`時會先計算當下`defer function`所讀取的值

```go
func main() {
	deferTest()
}

func deferTest() {
	defer fmt.Println("deferTest1")
	defer fmt.Println("deferTest2")
	defer fmt.Println("deferTest3")

	panic("panicTest")
}
```



```go

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	// 會先計算calc("10", a, b)
	// 返回的結果給 calc("1", a, b) <- 此時尚未執行，加入defer stack
	a = 0
	//先前的defer function不影響
	defer calc("2", a, calc("20", a, b))
	// 會先計算calc("20", a, b)
	// 返回的結果給 calc("2", a, b) <- 此時尚未執行，加入defer stack
	b = 1

	// output:
	// 10 1 2 3
	// 20 0 2 2
	//  2 0 2 2
	//  1 1 3 4
}


```