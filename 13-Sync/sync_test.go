package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)
	})

	t.Run("counter runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
