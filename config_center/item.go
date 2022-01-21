package config_center

import "sync"

type Entry struct {
	Key   string
	Value string
}

type Item struct {
	Action    int64
	Namespace string
	Key       string
	value     string
	Entries   []*Entry
	IsDefault bool
	sync.RWMutex
}

func (v *Item) SetValue(val string) {
	v.Lock()
	defer v.Unlock()
	v.value = val
}

func (v *Item) GetValue() string {
	v.RLock()
	defer v.RUnlock()
	return v.value
}
