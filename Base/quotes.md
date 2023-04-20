# Single quotes V.S Double quotes V.S Back quotes

**Double quotes:**
It is used to define a string.

**Back quotes:**
A string encoded in back quotes is a raw literal string and **doesn’t honor any kind of escaping**.

**Single quotes:**
To declare either a byte or a rune we use single quotes. While declaring byte we have to specify the type
**A single quote will allow only one character**

```go
func main() {

    //String in double quotes
    x := "tit\nfor\ttat"
    fmt.Println("Priting String in Double Quotes:")
    fmt.Printf("x is: %s\n", x) 
	// x is: tit
	// for   tat
    
   //String in back quotes
    y := `tit\nfor\ttat`
    fmt.Println("\nPriting String in Back Quotes:")
    fmt.Printf("y is: %s\n", y) 
	// tit\nfor\ttat


	var b byte = 'a'
	r := '你' // rune type
	fmt.Println(b, reflect.TypeOf(b)) // 97 uint8
	fmt.Println(r, reflect.TypeOf(r)) // 20320 int32
}

```