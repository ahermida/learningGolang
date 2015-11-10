/**
  An implementation of the Heap ADT, Use slices. (Mostly Modified Copy Pasta)
*/
package heap

import (
	"sync"
)

/**
  Define Struct of type Heap, use lock to make temporarily immutable on changes.
  Using an interflace slice to represent queue because they're not forced
  to a particular type.
*/
type Heap struct {
	sync.Mutex
	heap []interface{}
	min  bool
}

/**
  Create Heap Function -- One must specify whether max or min heap
*/
func NewHeap(max bool) *Heap {
  if (max) {
  	return &Heap{
  		heap: make([]interface{}, 0),
      min:  false,
  	}
  } else {
    return &Heap{
      heap: make([]interface{}, 0),
      min:  true,
    }
  }
}

/**
  Compare function for Reheap Algorithm
*/
func (h *Heap) Compare(a, b interface{}) bool {
  switch a.(type) {
    case string:
      if h.min {
        return a.(string) <= b.(string)
      } else {
        return b.(string) >= a.(string)
      }
    case int:
      if h.min {
        return a.(int) <= b.(int)
      } else {
        return b.(int) <= a.(int)
      }
    default:
      return false
  }
}

/**
  Getter for items, takes index
*/
func (h *Heap) Get(index int) interface{} {
  if index < len(h.heap) && index > -1 {
  	return h.heap[index]
  } else {
    return nil
  }
}

/**
  Insert item into heap and reheap
*/
func (h *Heap) Insert(item interface{}) {
	h.Lock()
	defer h.Unlock()

	h.heap = append(h.heap, item)
	h.siftUp()

	return
}

/**
  Remove top element on the heap
*/
func (h *Heap) Extract() (item interface{}) {
	h.Lock()
	defer h.Unlock()
	if len(h.heap) == 0 {
		return
	}

	item = h.heap[0]
  last := make([]interface{}, 1)
	last[0] = h.heap[len(h.heap)-1]
	if len(h.heap) == 1 {
		h.heap = nil
		return
	}
	h.heap = append(last, h.heap[1:len(h.heap)-1]...)
	h.siftDown()

	return
}

func (h *Heap) siftUp() {
	for i, parent := len(h.heap)-1, len(h.heap)-1; i > 0; i = parent {
		parent = i >> 1
		if h.Compare(h.heap[i], h.heap[parent]) {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	for i, child := 0, 1; i < len(h.heap) && i<<1+1 < len(h.heap); i = child {
		child = i<<1 + 1

		if child+1 <= len(h.heap)-1 && h.Compare(h.heap[child+1], h.heap[child]) {
			child++
		}

		if h.Compare(h.heap[i], h.heap[child]) {
			break
		}

		h.heap[i], h.heap[child] = h.heap[child], h.heap[i]
	}
}
