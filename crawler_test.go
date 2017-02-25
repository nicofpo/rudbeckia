package rudbeckia

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	crawler := NewCrawler(FEED_URL)

	videos, err := crawler.Fetch()
	if err != nil {
		t.Fatal(err)
	}

	for i, video := range videos {
		t.Logf("%d: %#v", i, video)
	}
}
