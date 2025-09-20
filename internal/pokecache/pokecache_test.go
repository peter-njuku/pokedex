package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Errorf("cache if nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()
	cache.Add("Key1", []byte("val1"))
	actual, ok := cache.Get("Key1")
	if !ok {
		t.Errorf("Key 1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("Value does not match")
	}
}

func TestReapLoopCache(t *testing.T) {
	cache := NewCache()
	cache.Add("Key1", []byte("val1"))
	actual, ok := cache.Get("Key1")
	if !ok {
		t.Errorf("Key 1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("Value does not match")
	}
}
