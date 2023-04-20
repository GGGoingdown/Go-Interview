# Map


`map` 是一個reference type. like pointers or slices. (can be nil)  
`map` 的key需要是comparable. 以下的三種資料型態無法當成key
- slice
- maps
- functions

當用以下的方式初始化時，會產生`nil`，所以當要寫入時，則會造成`panic`
```go
func main() {
	var m map[int]string

	if m == nil {
		fmt.Println("m is nil")
	}

	m[0] = "eddie" // panic: assignment to entry in nil map
}
```

讀取map時可以返回兩個值，用來確認key是否存在
```go
func main() {
	m := map[int]string{
		0: "zero",
		1: "one",
	}
	val, ok := m[2]
	fmt.Println(val, ok) // "" false
}
```

**Exploiting zero values:**  
[官網](https://go.dev/blog/maps#TOC_6.)上列出幾個例子告訴我們說如何利用zero values
```go
    type Person struct {
        Name  string
        Likes []string
    }
    var people []*Person

    likes := make(map[string][]*Person) 
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }
```

**Concurrency read and write:**  
`maps` are not safe for concurrent use, 所以當我們需要讀寫時需要加入lock
```go
// Example
var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}

// read
counter.RLock()
val := counter.m["some_key"]
counter.RUnlock()

// write
counter.Lock()
counter.m["some_key"]++
counter.Unlock()
```