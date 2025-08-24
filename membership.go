package main

func (s *Store) AddUsersToSegment(segKey string, userIDs []int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	set := s.explicitAdds[segKey]
	if set == nil {
		set = make(map[int64]struct{})
		s.explicitAdds[segKey] = set
	}
	for _, id := range userIDs {
		set[id] = struct{}{}
		if s.explicitRems[segKey] != nil {
			delete(s.explicitRems[segKey], id)
		}
	}
	return nil
}

func (s *Store) RemoveUsersFromSegment(segKey string, userIDs []int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	rem := s.explicitRems[segKey]
	if rem == nil {
		rem = make(map[int64]struct{})
		s.explicitRems[segKey] = rem
	}
	for _, id := range userIDs {
		rem[id] = struct{}{}
		if s.explicitAdds[segKey] != nil {
			delete(s.explicitAdds[segKey], id)
		}
	}
	return nil
}
