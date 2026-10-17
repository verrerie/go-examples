package main

import "fmt"

func main() {
	kl := Constructor(4, []int{7, 7, 7, 7, 8, 3})
	fmt.Println(kl.Add(2))
	fmt.Println(kl.Add(10))
	fmt.Println(kl.Add(9))
	fmt.Println(kl.Add(9))
	kl = Constructor(3, []int{4, 5})
	fmt.Println(kl.Add(6))
	kl = Constructor(1, nil)
	fmt.Println(kl.Add(7))
	kl = Constructor(4, []int{1, 2, 3})
	fmt.Println(kl.Add(4))
	kl = Constructor(4, []int{2, 3, 4})
	fmt.Println(kl.Add(1))
}

// KthLargest requirements
// 1. restate give an array of int and number k, we receive a stream of score (int), we want to return the kth largest score at any time
// 2. initial input is not ordered
// 3. score >= 0 ?, size of stream + input ? invalid input ? max, min ?
type KthLargest struct {
	arr []int
	k   int
}

// 1. naive approach: sort numbers O(L*logL), then insert the new numer O(N), return position N-K+1
// 2. not better: use max-heap, heapify numbers O(L), then add to heap O(logN), pop K times to get the ans, then add them back O(LogN*K)
// 3. better: use min-heap, heapify numbers O(L), pop L-K so only K left - O(L-K * logL), we maintain a heap of K elements
// for every new element, if > peek: pop then add: O(logK), ans = peek
// invariant: keep the min(totalSee, k) largest value in a min-heap
// 4. allowed to mutable the input ?

func Constructor(k int, nums []int) KthLargest {
	// time = O(L)
	heapify(nums)
	// time = max(0, O((L-k)logL)), worst case: O(L*logL) if k is very small
	if len(nums) > k {
		for range len(nums) - k {
			nums = pop(nums)
		}
	}
	return KthLargest{
		arr: nums,
		k:   k,
	}
}

// time = O(L)
func heapify(arr []int) {
	for lastParent := (len(arr) / 2) - 1; lastParent >= 0; lastParent-- {
		index := lastParent
		siftDown(arr, index)
	}
}

// time = O(logL)
func pop(arr []int) []int {
	swap(arr, 0, len(arr)-1)
	arr = arr[:len(arr)-1]
	siftDown(arr, 0)
	return arr
}

// time = O(1)
func peek(arr []int) int {
	return arr[0]
}

// time = O(LogL)
func siftDown(arr []int, start int) {
	index := start
	for index < len(arr) {
		left := index*2 + 1
		right := index*2 + 2
		small := index
		if left < len(arr) && arr[left] < arr[small] {
			small = left
		}
		if right < len(arr) && arr[right] < arr[small] {
			small = right
		}
		if small == index {
			break
		}
		swap(arr, index, small)
		index = small
	}
}

// time = O(logk)
func siftUp(arr []int, start int) {
	index := start
	for index > 0 {
		parent := (index - 1) / 2
		if parent >= 0 && arr[index] < arr[parent] {
			swap(arr, parent, index)
			index = parent
		} else {
			break
		}
	}
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

// Add time = O(logk) once heap is full at size k
func (kl *KthLargest) Add(val int) int {
	if len(kl.arr) < kl.k {
		kl.arr = append(kl.arr, val)
		siftUp(kl.arr, len(kl.arr)-1)
	} else if val > peek(kl.arr) {
		kl.arr[0] = val
		siftDown(kl.arr, 0)
	}
	return kl.arr[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
