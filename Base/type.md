# Type

**類別定義:**  
定義一個客製化的類別名稱
```go
type MyInt1 int

func main(){
    var i MyInt1 = 10
    fmt.Printf("%T\n", i) // MyInt1
}
```

```go
type set map[string]struct{}

func (s set) Add(k string) {
	s[k] = struct{}{}
}

func (s set) Remove(k string) {
	delete(s, k)
}

func (s set) Has(k string) bool {
	_, ok := s[k]
	return ok
}

func main() {
	s := make(set)
	s.Add("10")
	fmt.Println(s.Has("10")) // true
}

```


**類別別名(alias):**  
把原有的類別給上一個名稱，底層還是原有的類別。
```go
type MyInt2 = int

func main(){
    i := 10
    var myInt MyInt2 = i 
    fmt.Printf("%T\n", myInt) // int
}
```


**Example:**
```go
// cannot use i (variable of type int) as type MyInt1 in variable declaration
type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i // error
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
	fmt.Printf("%T, %T\n", i1, i2)
}
````