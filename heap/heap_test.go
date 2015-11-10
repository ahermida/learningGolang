package heap

import (
	//"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewHeap(false)

	h.Insert("hello")
	h.Insert("yo")
	h.Insert("sup")
	h.Insert("do")
	h.Insert("it")
	h.Insert("just")
	h.Insert("jimmy")
	h.Insert("russel")
	h.Insert("supdude")
	h.Insert("heyhey")

	sorted := []string{}
	for len(h.heap) > 0 {
		sorted = append(sorted, h.Extract().(string))
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

	h.Insert("hello")
	h.Insert("yo")
	h.Insert("sup")
	h.Insert("do")
	h.Insert("it")
	h.Insert("just")
	h.Insert("jimmy")
	h.Insert("russel")
	h.Insert("supdude")
	h.Insert("heyhey")
	sorted := []string{}
	for len(h.heap) > 0 {
		sorted = append(sorted, h.Extract().(string))

	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] < sorted[i+1] {
			t.Log(sorted)
			t.Error()
		}
    t.Log(sorted)
	}


}
