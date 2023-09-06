package set

import (
	"testing"
	"time"
)

const DEFAULT_TIME_DURATION_MS = 1

func TestTTLSet_Add(t *testing.T) {

	// Creating a ttl of 1 MS
	ttl := time.Duration(DEFAULT_TIME_DURATION_MS) * time.Millisecond
	set := NewTTLSet(ttl)
	set.Add("test")
	presentBefore := set.Contains("test")
	// Sleep at least 2 seconds
	time.Sleep(ttl)
	presentAfter := set.Contains("test")

	// Checking that the value was added to the set
	if presentBefore != true {
		t.Error("Error: expected true but got false")
	}

	// Checking that the value was 'removed' after expiring
	if presentAfter != false {
		t.Error("Error: expected false but got true")
	}

	t.Log("Success")
}

func TestTTLSet_isExpired(t *testing.T) {

	// Creating a ttl of 1 MS
	ttl := time.Duration(DEFAULT_TIME_DURATION_MS) * time.Millisecond
	set := NewTTLSet(ttl)

	set.Add("test")
	time.Sleep(ttl)

	// Testing to see if the value is expired
	if isExpired := set.isExpired("test"); isExpired {
		t.Log("Success")
	} else {
		t.Error("Error: expected true but got ", isExpired)
	}
}

func TestTTLSet_RemoveExpired(t *testing.T) {

	// Add a bunch of items to the set and check if RemoveExpired removes all expired entries
	// Creating a ttl of 1 MS
	ttl := time.Duration(DEFAULT_TIME_DURATION_MS) * time.Millisecond
	set := NewTTLSet(ttl)

	values := [...]string{"foo", "bar", "baz"}

	for _, val := range values {
		set.Add(val)
	}

	setValues := set.GetValues()

	if len(setValues) != len(values) {
		t.Error("Error: expected ", len(setValues), " but got ", len(values))
	}

	// Sleep for a bit to allow expiration times to be reached
	time.Sleep(ttl)

	set.RemoveExpired()

	// Expect set to be empty at this point
	if len(set.m) != 0 {
		t.Error("Error: expected 0 but got ", len(set.m))
	}

	t.Log("Success")
}

func TestTTLSet_GetValues(t *testing.T) {

	// Creating a ttl of 1 MS
	ttl := time.Duration(DEFAULT_TIME_DURATION_MS) * time.Millisecond
	set := NewTTLSet(ttl)

	values := [...]string{"foo", "bar", "baz"}

	for _, val := range values {
		set.Add(val)
	}

	setValues := set.GetValues()

	if len(setValues) != len(values) {
		t.Error("Error: expected ", len(setValues), " but got ", len(values))
	}

	// Sleep for a bit to allow expiration times to be reached
	time.Sleep(ttl)

	updatedValues := set.GetValues()

	// Expect set to be empty at this point
	if len(updatedValues) != 0 {
		t.Error("Error: expected 0 but got ", len(updatedValues))
	}

	t.Log("Success")
}
