# HashMap


```go
const TableSize = 26

type Bucket struct {
	Key   string
	Value string
	Next  *Bucket
}

type LinkedList struct {
	Head *Bucket
}

type HashTable struct {
	Table [TableSize]*LinkedList
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (h *HashTable) hash(s string) int {
	// Create a variable to store the total of all the characters
	total := 0
	// Convert the string into a slice of bytes
	byteSlice := []byte(s)
	// Loop through the slice and add up the totals
	for _, v := range byteSlice {
		total += int(v)
	}
	// Use modulus to get the remainder of the sum divided by the table size
	return total % TableSize
}

// Insert takes a key and value as arguments and inserts it into the hash
func (h *HashTable) Insert(key string, value string) {
	// Get the index of the key, using the hash function
	index := h.hash(key)

	// If the index is empty, create a new linked list with a head node
	if h.Table[index] == nil {
		h.Table[index] = &LinkedList{
			Head: &Bucket{
				Key:   key,
				Value: value,
			},
		}
		return
	}

	// Otherwise, traverse the linked list and append the new node
	curBucket := h.Table[index].Head
	for curBucket.Next != nil {
		curBucket = curBucket.Next
	}
	curBucket.Next = &Bucket{
		Key:   key,
		Value: value,
	}

}

// Get returns the value associated with the key in the hash table.
// It returns an empty string and false if the key is not found.
func (h *HashTable) Get(key string) (string, bool) {
	// Get the index for the key
	index := h.hash(key)

	// If the index is not present in the hash table, return false
	if h.Table[index] == nil {
		return "", false
	}

	// Find the bucket for the key
	curBucket := h.Table[index].Head
	for curBucket != nil {
		// If the key is found, return the value
		if curBucket.Key == key {
			return curBucket.Value, true
		}
		curBucket = curBucket.Next
	}

	// If the key is not found, return false
	return "", false
}


// Delete removes a key from the hash table
// Returns true if the key was found and removed, false otherwise
func (h *HashTable) Delete(key string) bool {

	// Get the index where the key should be stored
	index := h.hash(key)

	// If the bucket is empty, return false
	if h.Table[index] == nil {
		return false
	}

	// Set the current bucket to the head of the chain
	curBucket := h.Table[index].Head

	// If the key is the first element in the bucket, remove it and return true
	if curBucket.Key == key {
		h.Table[index].Head = curBucket.Next
		return true
	}

	// Iterate through the bucket until we find the key or reach the end
	for curBucket.Next != nil {
		if curBucket.Next.Key == key {
			curBucket.Next = curBucket.Next.Next
			return true
		}
		curBucket = curBucket.Next
	}

	// We've reached the end of the bucket and haven't found the key
	return false
}
```