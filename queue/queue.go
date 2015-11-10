/**
  An implementation of the Queue ADT
*/
package queue

import "sync"

/**
  Define Struct of type Queue, use lock to make temporarily immutable on changes.
  Using an interflace slice to represent queue because they're not forced
  to a particular type.
*/
type Queue struct {
  queue  []interface{}
  sync.Mutex
}

/**
  Queue Constructor
*/
func NewQueue() *Queue {
  q := &Queue{}
  q.queue = make([]interface{}, 0)
  return q
}

/**
  Flesh out the Queue struct
*/
/**  Enqueue Method, add to back */
func (q *Queue) Enqueue(item interface{}){
  q.Lock()
  defer q.Unlock()
  q.queue = append(q.queue, item)
}

/**  Dequeue Method remove from front */
func (q *Queue) Dequeue() (item interface{}) {
  q.Lock()
  defer q.Unlock()
  item, q.queue = q.queue[0], q.queue[1:]
  return
}

/**  Peek Method, get index */
func (q *Queue) Peek(n int) (item interface{}) {
  item = q.queue[n]
  return
}

/**  Length Method, simply returns */
func (q *Queue) Length() int{
  return len(q.queue)
}
