/**
  Test of the queue ADT
*/
package queue

import (
         "fmt"
         "testing"
)

func TestNew(t *testing.T) {
  queue := NewQueue()
  queue.Length()
  queue.Enqueue("hello")
  queue.Enqueue(1)
  if queue.queue[1] != 1 {
    t.Error()
  }
  queue.Enqueue("How are you?\n")
  queue.Enqueue("I'm well, you?")
  if queue.queue[0] != "hello" || queue.queue[3] != "I'm well, you?" {
     fmt.Printf("queue.queue")
     t.Error()
  }
  queue.Enqueue("Yoo")
  popped := queue.Dequeue()
  if popped != "hello" {
    t.Error()
  }
  if queue.Length() != 4 {
    fmt.Printf("%d", len(queue.queue))
    t.Error()
  }

  peeked := queue.Peek(3)
  if peeked != "Yoo" {
    fmt.Printf(queue.queue[3].(string))
    fmt.Printf("%d", queue.Length())
    t.Error()
  }
}
