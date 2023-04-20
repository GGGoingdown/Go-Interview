# Embedding

Go語言中沒有繼承的概念，取而代之的是embedding (composition)



**Code:**
```go
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.People.ShowA() // showA showB
	t.ShowA()        // showA showB (same as above)
	t.ShowB()			// teacher showB
}

```