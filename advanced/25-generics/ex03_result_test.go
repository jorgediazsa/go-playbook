package generics

import (
	"testing"
)

func TestScrapeWorkerResult(t *testing.T) {
	/* TODO: Uncomment this test after refactoring ex03_result.go!

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	urls := []string{"http://good.com", "http://bad.com"}

	// Create the generic channel
	resultsCh := make(chan Result[string], 2)

	go ScrapeWorker(ctx, urls, resultsCh)

	// Read Result 1 (Good)
	res1 := <-resultsCh
	if res1.Err != nil || res1.Value != "<html>Body of http://good.com</html>" {
		t.Fatalf("Failed to scrape good URL, got: %v", res1)
	}

	// Read Result 2 (Bad)
	res2 := <-resultsCh
	if res2.Err == nil {
		t.Fatalf("FAILED: Expected to receive an error for bad.com over the channel, got nil error and value: %q", res2.Value)
	}
	*/
}
