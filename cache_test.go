package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/whynayemnay/pokedex/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestEmptyCache(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	// attempt to get from empty cache
	_, ok := cache.Get("https://test.com")
	if ok {
		t.Errorf("expected to not find a key")
	}
}

func TestMultipleEntries(t *testing.T) {
	const interval = 5 * time.Second
	cache := pokecache.NewCache(interval)

	// Define test cases
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example1.com",
			val: []byte("example1"),
		},
		{
			key: "https://example2.com",
			val: []byte("example2"),
		},
		{
			key: "https://example3.com",
			val: []byte("example3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.val)
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Check %s", c.key), func(t *testing.T) {
			val, ok := cache.Get(c.key)
			if !ok || string(val) != string(c.val) {
				t.Errorf("expected %s key to return '%s', got %v", c.key, string(c.val), val)
			}
		})
	}
}
