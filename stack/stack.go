/**
  An implementation of the Stack ADT
*/
package stack

import "sync"

/**
  Define Struct of type Stack, use lock to make temporarily immutable on changes.
  Using an interflace slice to represent stack because they're not forced
  to a particular type.
*/
type Stack struct {
  stack  []interface{}
  length int
  lock   sync.Mutex
}

/**
  Stack constructor
*/
func NewStack() *Stack {
  s := &Stack{}
  s.stack = make([]interface, 0)
  s.length = 0
  return s
}

/**
  Flesh out the Stack struct
*/
func (s *Stack) Push(item interface{}) {
  s.lock.Lock()
  defer s.lock.Unlock()
  copy(s.stack[1:], s.stack)
  s.stack[0] = item;
  s.length++;
}

func (s *Stack) Pop() (item interface{}){

}

func (s *Stack) Peek(){

}

func (s *Stack) Length(){

}
