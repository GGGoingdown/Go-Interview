# For

**Caution:**  
`for loop`裡面的變數 (此範例的`i`)記憶體位置都一樣。
```go
for i := 0; i <= 2; i++ {
    fmt.Printf("%p\n", &i)
    // 0xc00001e0a8
}
```

如以下的範例:
```go
package main

import "fmt"

func main() {
    slice := []int{0, 1, 2, 3}
    m := make(map[int]*int)

    for key, val := range slice {
        m[key] = &val 
    }

    for k, v := range m {
        fmt.Println(k, "->", *v) // all the v = 3
    }
}
```