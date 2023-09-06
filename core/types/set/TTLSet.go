package set

import (
	mlog "github.com/MarconiFoundation/log"
	"time"
)

/*
	TTLSet is similar to Set except
	entries only 'exist' for a set amount of time (time.Duration)
	This follows an access model (expired entries are removed at access time and not returned)
	This is more efficient than running a background timer/goroutine to remove expired entries for the
	time being
*/

type TTLSet struct {
	m   map[string]time.Time
	ttl time.Duration
}

/*
	Creates a new instance of TTLSet
	This is like a regular set except each entry
	has an 'updated time'
	If the value is requested and a duration of ttl
	has passed then the value will be removed from the set
*/
func NewTTLSet(ttl time.Duration) *TTLSet {
	s := &TTLSet{}
	s.m = make(map[string]time.Time)
	s.ttl = ttl
	return s
}

/*
	Adds a new value to the set
	If the value is already present then updatedTime gets
	updated
*/
func (s *TTLSet) Add(value string) {
	mlog.GetLogger().Info("Adding ", value, " ", time.Now())
	s.m[value] = time.Now()
}

/*
	Removes a value from the set
*/
func (s *TTLSet) Remove(value string) {
	delete(s.m, value)
}

/*
	Gets a value from the set (if exists or not expired)
	Otherwise returns an error
*/
func (s *TTLSet) Contains(value string) bool {
	if _, present := s.m[value]; present {
		if !s.isExpired(value) {
			return true
		}
		s.Remove(value)
		return false
	}
	return false
}

/*
	Checks if a value has expired
	Assumes the value is present in the set
*/
func (s *TTLSet) isExpired(value string) bool {

	entry := s.m[value]
	timePassed := time.Since(entry)
	mlog.GetLogger().Info(value, "Time passed", timePassed, "TTL", s.ttl)
	if timePassed > s.ttl {
		return true
	}
	return false
}

/*
	Removes all expired entries in the set
*/
func (s *TTLSet) RemoveExpired() {
	// Goes through the set and removes all expired entries
	for key := range s.m {
		if s.isExpired(key) {
			s.Remove(key)
		}
	}
}

/*
	Gets all values from the set (minus expired ones)
*/
func (s *TTLSet) GetValues() []string {
	// Removes all expired entries and returns the remaining
	mlog.GetLogger().Info("Called")
	values := make([]string, 0, len(s.m))

	for key := range s.m {
		mlog.GetLogger().Info(key, "time updated", s.m[key])
		if s.isExpired(key) {
			mlog.GetLogger().Info(key, "Has expired")
			s.Remove(key)
		} else {
			// If it is not expired, add it
			mlog.GetLogger().Info(key, "Still available")
			values = append(values, key)
		}
	}
	return values
}

func (s *TTLSet) Size() int {
	s.RemoveExpired()
	return len(s.m)
}
