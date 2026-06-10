package storage

import (
	"os"
	"sync"
)

type WAL struct {
	mu   sync.Mutex
	file *os.File
}

func OpenWAL(path string) (*WAL, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &WAL{file: file}, nil
}

func (w *WAL) Write(data []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	_, err := w.file.Write(append(data, '\n'))
	return err
}

func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.file.Close()
}
