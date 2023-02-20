package cache

import (
	"log"
	"os"
)

type FlushOperation string

const (
	PUT FlushOperation = "PUT"
	DEL                = "DEL"
)

type UnflushedObject struct {
	verb FlushOperation
	key  string
}

type MemoryFirstCache struct {
	flush_size int
	directory  string
	store      map[string]string
	flushQueue chan UnflushedObject
}

func New(flush_size int, directory string) MemoryFirstCache {
	m := make(map[string]string)
	c := make(chan UnflushedObject)
	mfc := MemoryFirstCache{
		flush_size, directory, m, c,
	}
    err := os.Mkdir(directory, 0777)
    if err != nil {
        log.Fatal(err)}
	return mfc
}
