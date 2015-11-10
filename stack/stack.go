/**
  An implementation of the Stack ADT, Use slices (top is the last element in slice)
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
  sync.Mutex
}

/**
  Stack constructor
*/
func NewStack() *Stack {
  s := &Stack{}
  s.stack = make([]interface{}, 0)
  return s
}

/**
  Flesh out the Stack struct
*/

/**  Push Method, add to front*/
func (s *Stack) Push(item interface{}){
  s.Lock()
  defer s.Unlock()
  s.stack = append(s.stack, item)
}

/**  Pop Method */
func (s *Stack) Pop() (item interface{}){
  s.Lock()
  defer s.Unlock()
  //goes up to the last index, exclusive
  item, s.stack = s.stack[len(s.stack) - 1], s.stack[:len(s.stack) - 1]
  return
}

/**  Peek Method, get index -- implicit return */
func (s *Stack) Peek(n int) (item interface{}){
  if n < len(s.stack) - 1 && n > -1 {
    item = s.stack[len(s.stack) - 1 - n]
  } else {
    item = nil
  }
  return
}

/**  Length Method, simply returns */
func (s *Stack) Length() int{
  return len(s.stack)
}
