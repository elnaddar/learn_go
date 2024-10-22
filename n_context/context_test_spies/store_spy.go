package context_test_spies

import (
	"context"
	"log"
	"testing"
	"time"
)

func SpyStore(response string, t *testing.T) *spyStore {
	return &spyStore{response, t}
}

type spyStore struct {
	response string
	t        *testing.T
}

// func (s *spyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *spyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

func (s *spyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
