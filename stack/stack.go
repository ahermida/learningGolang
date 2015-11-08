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
  length int
  lock   sync.Mutex
}

/**
  Stack constructor
*/
func NewStack() *Stack {
  s := &Stack{}
  s.stack = make([]interface{}, 0)
  s.length = 0
  return s
}

/**
  Flesh out the Stack struct
*/

/**  Push Method */
func (s *Stack) Push(item interface{}){
  s.lock.Lock()
  defer s.lock.Unlock()
  s.stack = append(s.stack, item)
  s.length++
}

/**  Pop Method */
func (s *Stack) Pop() (item interface{}){
  s.lock.Lock()
  defer s.lock.Unlock()
  item, s.stack = s.stack[s.length - 1], s.stack[:s.length - 2]
  s.length--
  return
}

/**  Peek Method, get index */
func (s *Stack) Peek(n int) (item interface{}){
  if n < s.length - 1 && n > -1 {
    item = s.stack[s.length - 1 - n]
  } else {
    item = nil
  }
  return
}

/**  Length Method, simply returns */
func (s *Stack) Length() int{
  return s.length
}
