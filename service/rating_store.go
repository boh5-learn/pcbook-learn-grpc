package service

import "sync"

// RatingStore is an interface to store laptop ratings
type RatingStore interface {
	Add(laptopId string, score float64) (*Rating, error)
}

// Rating contains the rating information of a laptop
type Rating struct {
	Count uint32
	Sum   float64
}

// InMemoryRatingStore stores laptop ratings in memory
type InMemoryRatingStore struct {
	mutex  sync.RWMutex
	rating map[string]*Rating
}

// Add adds a new laptop score to the store and returns its rating
func (s *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rating := s.rating[laptopID]
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	s.rating[laptopID] = rating
	return rating, nil
}

// NewInMemoryRatingStore returns a new InMemoryRatingStore
func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: make(map[string]*Rating),
	}
}
