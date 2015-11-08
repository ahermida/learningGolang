/**
  Test of the Stack ADT
*/
package stack

import (
         "fmt"
         "testing"
)

func TestNew(t *testing.T) {
  stack := NewStack()
  stack.Length()
  stack.Push("hello")
  stack.Push(1)
  fmt.Printf("%d\n", len(stack.stack))
  if stack.stack[1] != 1 {
    t.Error()
  }
  stack.Push("How are you?\n")
  stack.Push("I'm well, you?")
  if stack.stack[0] != "hello" || stack.stack[3] != "I'm well, you?" {
     fmt.Printf("stack.stack")
     t.Error()
  }
  fmt.Printf("%d\n", len(stack.stack))
  stack.Push("Yoo")
  fmt.Printf("%d\n", len(stack.stack))
  popped := stack.Pop()
  fmt.Printf("%d\n", len(stack.stack))
  if popped != "Yoo" {
    t.Error()
  }
  if stack.Length() != 4 {
    fmt.Printf("%d", stack.Length())
    t.Error()
  }

  peeked := stack.Peek(3)
  if peeked != "hello" {
    //t.Error()
  }
}
