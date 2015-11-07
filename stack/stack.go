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

/**  Push Method, add to front*/
func (s *Stack) Push(item interface{}){
  s.lock.Lock()
  defer s.lock.Unlock()
  copy(s.stack[1:], s.stack)
  s.stack[0] = item
  s.length++
}

/**  Pop Method, remove front -- implicit return */
func (s *Stack) Pop() (item interface{}){
  s.lock.Lock()
  defer s.lock.Unlock()
  item, s.stack = s.stack[0], s.stack[1:]
  s.length--
}

/**  Peek Method, get index -- implicit return */
func (s *Stack) Peek(n int) (item interface{}){
  if n < s.length && n > -1 {
    item = s.stack[n]
  } else {
    item = nil
  }
}

/**  Length Method, simply returns */
func (s *Stack) Length() int{
  return s.length
}
