# What's different between `new` and `make` ?

`new`有以下特點:
- 設定zero value (**not initialize**)
- 返回address (pointer)

`make`有以下特點:
- create `slices`, `maps` and `channels` only
- initialized (**not zeroed it**)
- 可以把`make`想像要創建`slices` `map` `channles`的預先準備，(references to data structure等) 


```go
func main() {
	u1 := new(User)
	fmt.Println(u1.Name == "") // true, string is empty (zero value)

	m1 := make(map[int]string)
	s1 := make([]string, 5)    // len=5, cap=5
	c1 := make(chan string, 1) // buffer size 1

}
```