# Slice and Array

`slice` 跟 `array`最大的差異為:
- `slice` 是沒有固定長度的`array`
- `array` 是初始化的時候就給定長度

`slice`有length(`len`) and capacity(`cap`)，`cap`是實際儲存資料的地方(底層就是`array`)，length則可以把它想像成一個window在`cap`上移動，所以`len`的長度勢必小於等於`cap`。  
`slice`最重要的一點是: **當`cap`達到上限時又新增資料時，底層的`array`會做置換的動作(copy original data to new array)**

**Important:**  
如果要對`slice` 或是 `array` 做切片處理，須注意以下事項  
假設定義一個`slice -> []int{1, 2, 3}`，長度為3，容量為3
對這個slice做切片`[i:j:k]`時結果如下，len=j-i，cap=k-i，如果j and k沒有設定的話則預設為原先size (設定時也不能超過原來的大小), i則為0
```go
	s := []int{1, 2, 3}
	s1 := s[:1] // [1], len=1-0=1, cap=3-0=3
	s2 := s[1:] // [2, 3], len=3-1=2, cap=3-1=2
```


**Example with linked:**
```go
func main() {
    // s 與 s1 底層的array為相同的array 
    // 所以當s1更改值時 s會一起更改
	s := []int{2, 3, 5}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=3 cap=3 [2 3 5]
	s1 := s
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1) // len=3 cap=3 [2 3 5]
	s1[0] = 9
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1) // len=3 cap=3 [9 3 5]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)    // len=3 cap=3 [9 3 5]
}
```

**Example with nolinked:**  
```go
func main() {
    // 當我們對s新增值時，由於s的cap以達到上限，所以會置換底層的array
    // 所以當s1更改值時，因為底層的array不同。 所以不會影響到s
	s := []int{2, 3, 5}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=3 cap=3 [2 3 5]
	s1 := s
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1) // len=3 cap=3 [2 3 5]
	s = append(s, 10)                                      // 因為 s 的容量不夠，所以會重新配置一個新的 slice
	s1[0] = 9                                              // s not effect
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1) // len=3 cap=3 [9 3 5]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)    // len=3 cap=3 [9 3 5]
}
```

**Copy:**  
複製slice時須注意 destination slice的長度`len`(not capacity)，因為這會決定複製的數量
**Example with copy:**
```go
	s := []int{2, 3, 5}
	s1 := make([]int, len(s))                              // len=3 cap=3
	copyied1 := copy(s1, s)                                // length of destination slice must greater than source
	fmt.Println(copyied1)                                  // 3
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1) // len=3 cap=3 [2 3 5]

	s2 := make([]int, 0, len(s)) // len=0 cap=3
	copyied2 := copy(s2, s)
	fmt.Println(copyied2)                                  // 0, because length of s2 is 0
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=0 cap=3 []
```

**Extend and re-slice:**
```go
func main() {
	s := []int{2, 3, 5}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=3 cap=3 [2 3 5]

	s = s[:1]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=1 cap=3 [2]

	s = s[:3]                                           // extend its length.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=3 cap=3 [2, 3, 5]

	s = s[1:]                                           // drop its first value, capacity changed

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=2 cap=2 [3, 5]

}

```

當要把slice or array傳入function時須注意以下事情:
- 如果function內要更改array值，需傳入pointer
- 如果function內要更改slice值，直接傳入即可，因為傳入function內的是`slice struct`而非array

**Example:**  
```go
func changeArray(a [3]int) {
	a[0] = 100
}
func changeArrayWithPointer(a *[3]int) {
	a[0] = 69
}

func changeSlice(a []int) {
	a[0] = 100
}

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	changeArray(a)
	fmt.Println(a) // [1 2 3]
	changeSlice(b)
	fmt.Println(b) // [100 2 3]
	changeArrayWithPointer(&a)
	fmt.Println(a) // [1 2 3]
}

```

### Slice and Array 初始化的比較 
```go
func main() {
	var a [3]int // init with zero-value
	fmt.Println(a) // [0, 0, 0]
	var b []int
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b) // len=0 cap=0 []
	if b == nil{
		fmt.Println("b is nil before append value") // this will execute
	}
	b = append(b, 1)
	if b == nil{
		fmt.Println("b is nil after append value")  // this not execute
	}
	c := make([]int, 3) // init with zero-value
	fmt.Printf("len=%d cap=%d %v\n", len(c), cap(c), c) // len=3 cap=3 [0 0 0]
}
```