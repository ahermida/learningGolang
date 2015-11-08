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
  stack.Push(" sir, \n")
  stack.Push("How are you?\n")
  stack.Push("I'm well, you?")
  if stack.stack[0] != "hello" || stack.stack[3] != "I'm well, you?" {
     fmt.Printf("stack.stack")
     t.Error()
  }
  stack.Push("Yoo")
  popped := stack.Pop()
  if popped != "Yoo" {
    t.Error()
  }
  if stack.Length() != 4 {
    t.Error()
  }

  peeked := stack.Peek(3)
  if peeked != "hello" {
    //t.Error()
  }
}
