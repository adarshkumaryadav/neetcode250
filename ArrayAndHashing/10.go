https://leetcode.com/problems/design-hashmap/description/

Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
 

Example 1:

Input
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output
[null, null, null, 1, -1, null, 1, null, -1]

-----------------------------------------------

type MyHashMap struct {
	MySlice []int
}

func Constructor() MyHashMap {
	myHash := MyHashMap{
		MySlice: make([]int, 1000001),
	}
	for i := range myHash.MySlice {
		myHash.MySlice[i] = -1
	}
	return myHash
}

func (this *MyHashMap) Put(key int, value int) {
	this.MySlice[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if this.MySlice[key] < 0 {
		return -1
	}
	return this.MySlice[key]
}

func (this *MyHashMap) Remove(key int) {
	this.MySlice[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */