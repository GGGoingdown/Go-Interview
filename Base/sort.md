# Sort

### Custom sorting by functions
需要實作以下三個function
- Less(i, j int)bool
- Len()int
- Swap(i, j int)
```go
import (
	"sort"
	"fmt"
)

type ArrayOfArray [][]int

func (a ArrayOfArray)Len()int{
	return len(a)
}

func (a ArrayOfArray)Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a ArrayOfArray)Less(i, j int)bool{
	return a[i][0] < a[j][0] //compare with 2D-array first value
}


func main(){
	array := [][]int{{1, 5}, {5, 1}, {2, 3}}
	sort.Sort(ArrayOfArray(array))
	fmt.Println(array) // [[1, 5], [2, 3], [5, 1]]
}


```