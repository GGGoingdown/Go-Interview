# Switch

Go語言中的`switch case`為，只要條件滿足則會break，如果要繼續執行則可以使用`fallthrough` keyword。


```go
func main(){
	value := 10
    // without fallthrough
	switch {
		case value <= 10:
			fmt.Println("Value is less than or equal to 10") // only print this line
		case value <= 5:
			fmt.Println("Value is less than or equal to 5")
		default:
			fmt.Println("Value is greater than 10") 
	}

    // with fallthrough
	switch {
		case value <= 10:
			fmt.Println("Value is less than or equal to 10") // print this line
            fallthrough
		case value <= 5:
			fmt.Println("Value is less than or equal to 5") // print this line
		default:
			fmt.Println("Value is greater than 10") 
	}


}
```


### Switch type aeertion
```go
var (
	i interface{}
)

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	default:
		println("type not found")
	}
}

func main() {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))
}

```