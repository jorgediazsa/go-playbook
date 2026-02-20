package goroutines

import (
	"errors"
)

// Context: Bounded Concurrent Executor
// You are building an image uploader. If a user uploads 1,000 images, you want to
// upload them to S3 concurrently, but you do NOT want to spawn 1,000 goroutines at once
// because it would overwhelm your network card or S3 rate limits.
//
// Why this matters: Instead of launching an unbounded number of goroutines,
// idiomatic systems use "WaitGroups" to ensure all workers finish before returning,
// and limit concurrency manually (using a semaphore channel, though here we'll
// just spawn N static workers and feed them work).
//
// Requirements:
// 1. Refactor `UploadAll` to spawn exactly `maxConcurrent` worker goroutines.
// 2. These workers should listen to a channel of filenames, upload them, and exit
//    when the channel is closed.
// 3. `UploadAll` must use a `sync.WaitGroup` to block until ALL workers have finished.
// 4. If ANY upload errors, the very first error encountered must be returned by `UploadAll`
//    at the end. The other workers can finish their current files, but we should return
//    the error once everything has exited.
//
// Note: We use a channel for the work queue because it is the idiomatic way
// to distribute work across a pool of goroutines cleanly.

func MockUploadS3(filename string) error {
	if filename == "corrupt.jpg" {
		return errors.New("failed to upload corrupt image")
	}
	return nil
}

func UploadAll(filenames []string, maxConcurrent int) error {
	// BUG: This current code is sequential. It is slow and uses 1 routine.
	// TODO: Spawn `maxConcurrent` workers. Use an unbuffered channel to feed them
	// `filenames`. Capture the first error across any worker, and block using
	// WaitGroup until all workers are done.

	for _, f := range filenames {
		if err := MockUploadS3(f); err != nil {
			return err
		}
	}

	return nil
}
