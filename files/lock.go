package files

import (
	"sync"
	"syscall"

	"github.com/pkg/errors"
)

// FileLock is FileLock object
type FileLock struct {
	l  sync.Mutex
	fd int
}

// NewFileLock is to open file with filelock
func NewFileLock(fileName string) (*FileLock, error) {
	if fileName == "" {
		return nil, errors.New("fileName is empty")
	}

	fd, err := syscall.Open(fileName, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		return nil, err
	}
	return &FileLock{fd: fd}, nil
}

// Lock is to lock file
func (m *FileLock) Lock() error {
	m.l.Lock()
	//排他ロック
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		return err
	}
	return nil
}

// Unlock is to unlock file
func (m *FileLock) Unlock() error {
	//ロック解除
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		return err
	}
	m.l.Unlock()
	return nil
}
