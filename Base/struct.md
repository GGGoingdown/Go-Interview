# Struct


**Caution:**  
`struct`可以比較，但需要注意以下事項:
1. 只能比較是否相等，不能比大小。 在compile time會出現error (invalid operation: sn1 >= sn2 (operator >= not defined on struct))
2. 同類型的`struct`才可以比較。 不同類型的`struct`比較時會在compile time 出現error (invalid operation: sn1 == sn2 (mismatched types student and struct{age int; name string; gender string}))
3. 如果`struct`中含有以下資料類型，則**無法**比較。 如有的話則會在compile time出現error (invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared))
    - `map`
    - `slice`
    - `func`


**Examples:**
```go
// invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
func main(){
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 { // error
		fmt.Println("sm1 == sm2")
	}
}
```


```go
// 就算sn1跟sn2的`struct` properties are all the same.
// they are not equally.
// sn1 == sn3

type student struct {
	age  int
	name string
}

func main() {
	sn1 := student{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}
	sn3 := student{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	if sn1 == sn3 {
		fmt.Println("sn1 == sn3")
	}
}
```