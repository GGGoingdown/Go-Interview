# Iota

**Definition:**
`iota` 初始化為0，_(底線)表示不賦值
`iota`的之前如果有值的話，則會接續之後做初始化。

**Example:**
```go
const (
	x = iota    // 0
	_           // 1
	y           // 2
	z = "zz"    // "zz"
	k           // "zz"
	p = iota    // 5 因為上面有一個iota了(初始化過)
)

```


**Example:**
```go
const (
	x = "e"    // 0         
	y = iota   // 1
    z          // 2
)

```