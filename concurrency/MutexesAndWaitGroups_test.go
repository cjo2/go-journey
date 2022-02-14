package concurrency

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestConcurrentFileParser_GetFileToProcess(t *testing.T) {

	// No files to process should return false
	c := &ConcurrentFileParser{}
	_, isFile := c.GetFileToProcess()
	assert.False(t, isFile)

	// Get the only file in the slice
	c = &ConcurrentFileParser{filesToBeProcessed: []string{"file1.txt"}}
	filename, isFile := c.GetFileToProcess()
	assert.Equal(t, "file1.txt", filename)
	assert.Empty(t, c.filesToBeProcessed)
	assert.True(t, isFile)

	// Multiple files in the slice, but remove one
	c = &ConcurrentFileParser{filesToBeProcessed: []string{"file1.txt", "file2.txt"}}
	filename, isFile = c.GetFileToProcess()
	assert.Equal(t, "file1.txt", filename)
	assert.Equal(t, 1, len(c.filesToBeProcessed))
	assert.True(t, isFile)
	filename, isFile = c.GetFileToProcess()
	assert.Equal(t, "file2.txt", filename)
	assert.Empty(t, c.filesToBeProcessed)
	assert.True(t, isFile)
}

func TestConcurrentFileParser_Start(t *testing.T) {

	cfp := &ConcurrentFileParser{}

	// create UUIDs (use deterministic for easier reasoning in test)
	random := rand.New(rand.NewSource(1))
	var filenames []string
	fileCount := 100

	// Generate a bunch of filenames
	for i := 0; i < fileCount; i++ {
		id, _ := uuid.NewRandomFromReader(random)
		filenames = append(filenames, id.String()+".txt")
	}

	// Add files to be processed
	cfp.LoadFiles(filenames)
	assert.Equal(t, fileCount, len(cfp.filesToBeProcessed))

	// The following should be blocking
	start := time.Now()
	cfp.Start(fileCount)
	duration := time.Since(start)
	fmt.Println(duration)

	assert.Equal(t, 0, len(cfp.filesToBeProcessed))
	assert.Equal(t, 100, len(cfp.filesCompleted))
}
