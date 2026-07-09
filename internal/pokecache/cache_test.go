package pokecache

import (
	"bytes"
	"testing"
	"time"
)

const (
	key = "https://pokeapi.co/api/v2/location-area"
	val = `{"count" : "", "next" : "", "previous": "", "Results" [{}]"}`
)

func TestCache(t *testing.T) {

	t.Run("creates new value in cache", func(t *testing.T) {

		cache := NewCache(time.Millisecond * 10)

		got, want, ok := AddCache(t, cache)

		CheckBytes(t, got, want, ok)
	})

	t.Run("returns value from cache", func(t *testing.T) {

		cache := NewCache(time.Millisecond * 10)

		_, _, _ = AddCache(t, cache)

		got, ok := cache.Get(key)
		want := []byte(val)

		if ok != true {
			t.Errorf("got %t want true", ok)
		}

		if !bytes.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("deletes cache after time", func(t *testing.T) {

		cache := NewCache(time.Millisecond * 10)

		_, ok := cache.Add(key, []byte(val))

		time.Sleep(time.Millisecond * 15)

		got, ok := cache.Get(key)
		want := []byte(``)

		t.Helper()
		if ok != false {
			t.Errorf("got %t want true", ok)
		}

		if !bytes.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func CheckBytes(t testing.TB, got, want []byte, ok bool) {
	t.Helper()
	if ok != true {
		t.Errorf("got %t want true", ok)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func AddCache(t *testing.T, cache *Cache) ([]byte, []byte, bool) {
	t.Helper()
	got, ok := cache.Add(key, []byte(val))
	want := []byte(val)

	return got, want, ok
}
