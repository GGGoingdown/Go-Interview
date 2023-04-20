# Appepnd

以下程式說明`append`function 內部執行的機制。 
```go
func append(slice []Type, elems ...Type) []Type {
    // Get the capacity of the slice
    length := len(slice)
    capacity := cap(slice)
    
    // Check if the slice has enough capacity to hold the new elements
    if length + len(elems) > capacity {
        // If not, create a new underlying array with double the capacity
        newCapacity := capacity * 2
        if newCapacity < length + len(elems) {
            newCapacity = length + len(elems)
        }
        newSlice := make([]Type, length, newCapacity)
        copy(newSlice, slice)
        slice = newSlice
    }
    
    // Append the new elements to the slice
    slice = slice[:length + len(elems)]
    for i, elem := range elems {
        slice[length+i] = elem
    }
    return slice
}

```