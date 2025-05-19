https://leetcode.com/problems/design-hashset/description
Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

type MyHashSet struct {
    list []bool
}


func Constructor() MyHashSet {
    return MyHashSet{list : make([]bool, 1000001)}
}


func (this *MyHashSet) Add(key int)  {
    this.list[key]=true
}


func (this *MyHashSet) Remove(key int)  {
    this.list[key]=false
}


func (this *MyHashSet) Contains(key int) bool {
    return this.list[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */