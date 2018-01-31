package brut

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math"
	"runtime"
	"sync"
)

const interval = 16

var (
	wg     sync.WaitGroup
	mutex  sync.RWMutex
	done   bool
	result string
)

// Sha256 object
type Sha256 struct {
	hash    []byte
	chars   []byte
	threads int
	length  int
	base    int
}

// NewSha256 constructor
func NewSha256(hash string, chars []byte, length uint8) *Sha256 {
	h, err := hex.DecodeString(hash)
	if err != nil {
		panic(err)
	}
	s := Sha256{
		hash:    h,
		chars:   chars,
		threads: runtime.NumCPU(),
		length:  int(length),
		base:    len(chars),
	}
	return &s
}

// Run process calculations
func (s *Sha256) Run() (string, error) {
	max := int64(math.Pow(float64(s.base), float64(s.length)))
	threads := int64(s.threads)
	delta := max / threads
	done = false
	wg.Add(s.threads)

	for i := int64(0); i < threads; i++ {
		start := i * delta
		stop := start + delta
		if i == threads-1 {
			stop = max
		}
		go s.worker(start, stop, s.chars, s.hash)
	}

	wg.Wait()
	if done {
		return result, nil
	}
	return "", errors.New("Can't find result")
}

func (s *Sha256) worker(start, stop int64, chars, hash []byte) {
	defer wg.Done()
	hasher := sha256.New()
	length := s.length
	base := int64(s.base)
	buf := make([]byte, length, length)
	check := interval

	for i := start; i < stop; i++ {
		check--
		if check == 0 {
			mutex.RLock()
			exit := done
			mutex.RUnlock()
			if exit {
				return
			}
			check = interval
		}

		num := i
		for j := 0; j < length; j++ {
			buf[j] = chars[num%base]
			num /= base
		}

		hasher.Reset()
		hasher.Write(buf)
		if bytes.Equal(hash, hasher.Sum(nil)) {
			mutex.Lock()
			result = string(buf)
			done = true
			mutex.Unlock()
			return
		}
	}
}
