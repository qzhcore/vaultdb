package storage

import "sync"

type Engine struct {
	mu    sync.RWMutex
	store map[string]string
	wal   *WAL
}

func NewEngine(walPath string) (*Engine, error) {
	wal, err := OpenWAL(walPath)
	if err != nil {
		return nil, err
	}
	return &Engine{
		store: make(map[string]string),
		wal:   wal,
	}, nil
}

func (e *Engine) Set(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	err := e.wal.Write([]byte(key + ":" + value))
	if err != nil {
		return err
	}

	e.store[key] = value
	return nil
}

func (e *Engine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, exists := e.store[key]
	return val, exists
}
