package cuckoofilter_test

import (
	"bufio"
	"hash/fnv"
	"os"
	"testing"

	"github.com/stanxii/cuckoofilter"
)

func TestMultipleInsertions(t *testing.T) {
	cf := cuckoofilter.New()

	fd, err := os.Open("dict/words")
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var words [][]byte
	var wordCount uint
	for scanner.Scan() {
		word := []byte(scanner.Text())

		if !cf.Lookup(word) && cf.Insert(word) {
			wordCount++
		}
		words = append(words, word)
	}

	size := cf.Count()
	if size != wordCount {
		t.Errorf("Expected word count = %d, not %d", wordCount, size)
	}

	for _, word := range words {
		cf.Delete(word)
	}

	size = cf.Count()
	if size != 0 {
		t.Errorf("Expected word count = 0, not %d", size)
	}
}

func TestBasicInsertion(t *testing.T) {
	cf := cuckoofilter.New()
	if !cf.Insert([]byte("buongiorno")) {
		t.Errorf("Wasn't able to insert very first word, 'buongiorno'")
	}

	size := cf.Count()
	if size != 1 {
		t.Errorf("Expected size after insertion to be 1, not %d", size)
	}

	if !cf.Lookup([]byte("buongiorno")) {
		t.Errorf("Expected to find 'buongiorno' in filter set membership query")
	}

	if !cf.Delete([]byte("buongiorno")) {
		t.Errorf("Expected to be able to delete 'buongiorno' in filter")
	}

	if cf.Lookup([]byte("buongiorno")) {
		t.Errorf("Did not expect to find 'buongiorno' in filter after deletion")
	}

	size = cf.Count()
	if size != 0 {
		t.Errorf("Expected size after deletion to be 0, not %d", size)
	}
}

func TestInitialization(t *testing.T) {
	cf := cuckoofilter.New()
	size := cf.Count()
	if size != 0 {
		t.Errorf("Expected initial size to be 0, not %d", size)
	}
}

func TestConfigurationOptions(t *testing.T) {
	cf := cuckoofilter.New(
		cuckoofilter.Size(1<<18),
		cuckoofilter.BucketSize(4),
		cuckoofilter.FingerprintSize(2),
		cuckoofilter.MaximumKicks(500),
		cuckoofilter.HashFn(fnv.New64()),
	)
	size := cf.Count()
	if size != 0 {
		t.Errorf("Expected size to be 10, not %d", size)
	}
}

func BenchmarkInsertionAndDeletion(b *testing.B) {
	cf := cuckoofilter.New()
	for n := 0; n < b.N; n++ {
		cf.Insert([]byte("buongiorno"))
		cf.Delete([]byte("buongiorno"))
	}
}
