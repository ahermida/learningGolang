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
func (s *Stack) Push(item []interface{}) {
  

}

func (s *Stack) Pop(){

}

func (s *Stack) Length(){

}
