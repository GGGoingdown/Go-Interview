# Type assertions


**Definition:**
只有`interface type`可以做type assertion，如果不支援此型別，則會造成`panic`

**Example:**
```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

### Type assertion for map 
```go
import (
	"fmt"
	"strconv"
)

func main() {
	var test = map[string]interface{}{
		"test01": []interface{}{"a", "b"},
		"test02": []int{1, 2},
	}

	oids1 := make([]string, len(test["test01"].([]interface{})))
	for i, v := range test["test01"].([]interface{}) {
		oids1[i] = v.(string)
	}
	fmt.Println(oids1)

	oids2 := make([]string, len(test["test02"].([]int)))
	for i, v := range test["test02"].([]int) {
		oids2[i] = strconv.Itoa(v)
	}
	fmt.Println(oids2)

	oids3 := make([]int, len(test["test02"].([]int)))
	for i, v := range test["test02"].([]int) {
		oids3[i] = v
	}
	fmt.Println(oids3)
}
```