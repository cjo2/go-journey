package concurrency

import (
	"sync"
	"time"
)

type ConcurrentFileParser struct {
	filesToBeProcessed []string
	filesInProgress    map[string]bool
	filesCompleted     []string
	queueLock          sync.Mutex
	process            sync.WaitGroup
}

// GetFileToProcess is a FIFO implementation
func (c *ConcurrentFileParser) GetFileToProcess() (string, bool) {
	c.queueLock.Lock()
	defer c.queueLock.Unlock()

	var toReturn string

	if len(c.filesToBeProcessed) == 0 {
		return "", false
	}

	toReturn = c.filesToBeProcessed[0]
	c.filesToBeProcessed = c.filesToBeProcessed[1:]
	return toReturn, true
}

func (c *ConcurrentFileParser) AcknowledgeFileProcessed(filename string) {
	c.queueLock.Lock()
	defer c.queueLock.Unlock()

	c.filesCompleted = append(c.filesCompleted, filename)
}

// ProcessFile simulates a "blocking" function like file I/O
// All we're doing is making this function sleep, then return a boolean
func (c *ConcurrentFileParser) ProcessFile() {
	filename, isFile := c.GetFileToProcess()
	if isFile {
		c.process.Add(1)
		defer c.process.Done()
		time.Sleep(5 * time.Second)
		c.AcknowledgeFileProcessed(filename)
	}
}

func (c *ConcurrentFileParser) LoadFiles(f []string) {
	c.filesToBeProcessed = append(c.filesToBeProcessed, f...)
}

func (c *ConcurrentFileParser) Start(fileCount int) {
	for i := 0; i < fileCount; i++ {
		go c.ProcessFile()
	}
	c.process.Wait()
}
