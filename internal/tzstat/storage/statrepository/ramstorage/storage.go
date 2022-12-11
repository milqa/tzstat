package ramstorage

import (
	"sort"
	"sync"
)

type storage struct {
	mu   sync.RWMutex
	data []event
}

type event struct {
	datetime int64
	value    int64
}

func (s *storage) insert(datetime, value int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := sort.Search(len(s.data), func(i int) bool { return s.data[i].datetime >= datetime })

	s.data = append(s.data, event{})
	copy(s.data[idx+1:], s.data[idx:])

	s.data[idx] = event{
		datetime: datetime,
		value:    value,
	}
}

func (s *storage) getEventsWithDatetime(datetimeFrom, datetimeTo int64) []event {
	s.mu.RLock()
	defer s.mu.RUnlock()

	firstIdx := sort.Search(len(s.data), func(i int) bool { return s.data[i].datetime >= datetimeFrom })
	lastIdx := sort.Search(len(s.data), func(i int) bool { return s.data[i].datetime > datetimeTo })

	result := make([]event, lastIdx-firstIdx)

	copy(result, s.data[firstIdx:lastIdx])

	return result
}
