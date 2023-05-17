# Error handler

### Error type
```go
import (
	"errors"
	"fmt"
)

func foo1() error {
	return errors.New("foo1 error")
}

func foo2() error {
	return fmt.Errorf("foo2 error")
}

func main() {
	var err error
	err = foo1()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = foo2()
	if err != nil {
		fmt.Println(err.Error())
	}

}
```

### Custom error
```go
type CustomError struct {
	Detail string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("custom error %s", c.Detail)
}

func foo1() error {
	return CustomError{"ops! something went wrong"}
}

func main() {
	if err := foo1(); err != nil {
		fmt.Println(err.Error())
	}

}
```