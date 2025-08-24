package main

import (
	"crypto/rand"
	"encoding/hex"
	"hash/fnv"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	mu           sync.RWMutex
	users        map[int64]*User
	segments     map[string]*Segment
	explicitAdds map[string]map[int64]struct{}
	explicitRems map[string]map[int64]struct{}
	salt         string
}

func NewStore() *Store {
	return &Store{
		users:        make(map[int64]*User),
		segments:     make(map[string]*Segment),
		explicitAdds: make(map[string]map[int64]struct{}),
		explicitRems: make(map[string]map[int64]struct{}),
		salt:         defaultSalt(),
	}
}

func defaultSalt() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// Возвращает true, если пользователь попадает в сегмент с учетом дат и rollout
func (s *Store) UserInSegment(userID int64, seg *Segment, now time.Time) bool {
	if seg.StartDate != nil && now.Before(*seg.StartDate) {
		return false
	}
	if seg.EndDate != nil && now.After(*seg.EndDate) {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if rset, ok := s.explicitRems[seg.Key]; ok {
		if _, removed := rset[userID]; removed {
			return false
		}
	}
	if aset, ok := s.explicitAdds[seg.Key]; ok {
		if _, added := aset[userID]; added {
			return true
		}
	}
	h := fnv.New64a()
	h.Write([]byte(seg.Key))
	h.Write([]byte{':'})
	h.Write([]byte(strconv.FormatInt(userID, 10)))
	h.Write([]byte{':'})
	h.Write([]byte(s.salt))
	bucket := h.Sum64() % 100
	return int(bucket) < seg.Percent
}

func (s *Store) ListSegments() []*Segment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*Segment, 0, len(s.segments))
	for _, seg := range s.segments {
		out = append(out, seg)
	}
	return out
}
