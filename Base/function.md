# Function

**Detail:**  
go 支援函數返回多個值，返回的值也可以賦予名稱，如果其中一個返回的值有賦予名稱，其他的也**必須**一併提供名稱，要不然會在compile time 出現error (error: mixed named and unnamed parameters)

如果返回的值有賦予名稱，在創建的時候會是`zero-value`

```go
func fun1(x, y int) (sum int, err error) {
	fmt.Println(sum, err) // 0 nil
	return x + y, nil
}

```