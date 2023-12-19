package w

import "sync"

var (
	Lock *lock
)

func init() {
	Lock = &lock{}
	Lock.lock_map = sync.Map{}
}

// Lock.Lock("a").Lock()
// Lock.Lock("a").Unlock()
func (k *lock) Lock(key string) *sync.Mutex {
	v, _ := k.lock_map.LoadOrStore(key, &sync.Mutex{})
	return v.(*sync.Mutex)
}
