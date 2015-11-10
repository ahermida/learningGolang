package heap

import (
	//"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewHeap(false)

	h.Insert(8)
	h.Insert(7)
	h.Insert(6)
	h.Insert(3)
	h.Insert(1)
	h.Insert(0)
	h.Insert(2)
	h.Insert(5)
	h.Insert(9)
	h.Insert(20)

	sorted := []int{}
	for len(h.heap) > 0 {
		sorted = append(sorted, h.Extract().(int))
	}
	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error()
		}
    t.Log(sorted)
	}

}

func TestMaxHeap(t *testing.T) {
	h := NewHeap(true)

  	h.Insert(8)
  	h.Insert(7)
  	h.Insert(6)
  	h.Insert(3)
  	h.Insert(1)
  	h.Insert(0)
  	h.Insert(2)
  	h.Insert(5)
  	h.Insert(9)
  	h.Insert(20)

	sorted := []int{}
	for len(h.heap) > 0 {
		sorted = append(sorted, h.Extract().(int))
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] < sorted[i+1] {
			t.Log(sorted)
			t.Error()
		}
    t.Log(sorted)
	}
}
