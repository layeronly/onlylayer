package memory

import (
	"testing"

	"github.com/onlylayer/onlylayer/blockchain/storage"
)

func TestStorage(t *testing.T) {
	t.Helper()

	f := func(t *testing.T) (storage.Storage, func()) {
		t.Helper()

		s, _ := NewMemoryStorage(nil)

		return s, func() {}
	}
	storage.TestStorage(t, f)
}