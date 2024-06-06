package cache

import (
	"sync"
	"time"
)

type image struct {
	Bytes          []byte
	LastAccessTime time.Time
}

type cacheImages struct {
	heap  []string
	cache map[string]*image
	mutex sync.RWMutex
	size  int
}

var Store cacheImages

type Option struct {
	Size  int
	Delay time.Duration
}

func New(opt Option) {

	Store = cacheImages{
		heap:  []string{},
		cache: make(map[string]*image, 1000),
		size:  opt.Size,
	}

	go func() {
		for {
			Store.handlerCleanup()
			time.Sleep(opt.Delay)
		}
	}()

}

func (ic *cacheImages) handlerCleanup() {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

	currentTime := time.Now()
	for key, img := range ic.cache {
		if currentTime.Sub(img.LastAccessTime) > 10*time.Minute {
			delete(ic.cache, key)
		}
	}
}

func (ic *cacheImages) GetImage(key string) (*[]byte, bool) {
	ic.mutex.RLock()
	defer ic.mutex.RUnlock()

	if img, ok := ic.cache[key]; ok {
		img.LastAccessTime = time.Now()
		return &img.Bytes, true
	}

	return nil, false
}

func (ic *cacheImages) SetImage(key string, data []byte) {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

	if len(ic.heap) >= ic.size {
		key := ic.heap[0]
		delete(ic.cache, key)
		ic.heap = append(ic.heap[:1], ic.heap[2:]...)
	}

	img := image{Bytes: data, LastAccessTime: time.Now()}
	ic.cache[key] = &img
	ic.heap = append(ic.heap, key)

}
