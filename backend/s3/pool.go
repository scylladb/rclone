package s3

import (
	"sync"
	"time"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/lib/pool"
)

type chunkMemPool struct {
	pools map[int]*sync.Pool
	mu    sync.Mutex
}

func (c *chunkMemPool) Get(flushTime time.Duration, bufferSize, poolSize int, useMmap bool) *pool.Pool {
	c.mu.Lock()
	sp, ok := c.pools[bufferSize]
	c.mu.Unlock()

	if ok {
		if v := sp.Get(); v != nil {
			fs.Logf(nil, "Cached pool")
			return v.(*pool.Pool)
		}
	}

	return pool.New(flushTime, bufferSize, poolSize, useMmap)
}

func (c chunkMemPool) Put(bufferSize int, p *pool.Pool) {
	c.mu.Lock()
	sp, ok := c.pools[bufferSize]
	if !ok {
		sp = new(sync.Pool)
		c.pools[bufferSize] = sp
	}
	c.mu.Unlock()

	sp.Put(p)
}

var s3ChunkMemPool = chunkMemPool{
	pools: make(map[int]*sync.Pool, 1),
}
