package lru

import (
	"testing"
)

func TestNew(t *testing.T) {
	cache := New[string, int](5)

	if cache.cap != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.cap)
	}

	if cache.items == nil {
		t.Error("Expected items map to be initialized")
	}

	if cache.ll == nil {
		t.Error("Expected linked list to be initialized")
	}

	if len(cache.items) != 0 {
		t.Errorf("Expected empty cache, got %d items", len(cache.items))
	}
}

func TestAddAndGet(t *testing.T) {
	cache := New[string, int](3)

	// Test adding and getting a single item
	cache.Add("key1", 100)

	value, ok := cache.Get("key1")
	if !ok {
		t.Error("Expected to find key1")
	}
	if value != 100 {
		t.Errorf("Expected value 100, got %d", value)
	}

	// Test getting non-existent key
	value, ok = cache.Get("nonexistent")
	if ok {
		t.Error("Expected not to find nonexistent key")
	}
	if value != 0 { // zero value for int
		t.Errorf("Expected zero value, got %d", value)
	}
}

func TestUpdateExistingKey(t *testing.T) {
	cache := New[string, int](3)

	cache.Add("key1", 100)
	cache.Add("key1", 200) // Update existing key

	value, ok := cache.Get("key1")
	if !ok {
		t.Error("Expected to find key1")
	}
	if value != 200 {
		t.Errorf("Expected updated value 200, got %d", value)
	}

	// Should still have only 1 item
	if len(cache.items) != 1 {
		t.Errorf("Expected 1 item after update, got %d", len(cache.items))
	}
}

func TestCapacityLimit(t *testing.T) {
	cache := New[string, int](2)

	cache.Add("key1", 1)
	cache.Add("key2", 2)

	// Cache should be full
	if len(cache.items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(cache.items))
	}

	// Add third item - should evict oldest
	cache.Add("key3", 3)

	// Should still have 2 items
	if len(cache.items) != 2 {
		t.Errorf("Expected 2 items after eviction, got %d", len(cache.items))
	}

	// key1 should be evicted (least recently used)
	_, ok := cache.Get("key1")
	if ok {
		t.Error("Expected key1 to be evicted")
	}

	// key2 and key3 should exist
	if _, ok = cache.Get("key2"); !ok {
		t.Error("Expected key2 to exist")
	}
	if _, ok = cache.Get("key3"); !ok {
		t.Error("Expected key3 to exist")
	}
}

func TestLRUOrdering(t *testing.T) {
	cache := New[string, int](3)

	cache.Add("key1", 1)
	cache.Add("key2", 2)
	cache.Add("key3", 3)

	// Access key1 to make it most recently used
	cache.Get("key1")

	// Add key4 - should evict key2 (least recently used)
	cache.Add("key4", 4)

	// key2 should be evicted
	if _, ok := cache.Get("key2"); ok {
		t.Error("Expected key2 to be evicted")
	}

	// key1, key3, key4 should exist
	if _, ok := cache.Get("key1"); !ok {
		t.Error("Expected key1 to exist")
	}
	if _, ok := cache.Get("key3"); !ok {
		t.Error("Expected key3 to exist")
	}
	if _, ok := cache.Get("key4"); !ok {
		t.Error("Expected key4 to exist")
	}
}

func TestGetUpdatesOrder(t *testing.T) {
	cache := New[string, int](3)

	cache.Add("key1", 1)
	cache.Add("key2", 2)
	cache.Add("key3", 3)

	// Access key1 to move it to front
	cache.Get("key1")

	// Access key2 to move it to front
	cache.Get("key2")

	// Add new item - should evict key3 (now least recently used)
	cache.Add("key4", 4)

	// key3 should be evicted
	if _, ok := cache.Get("key3"); ok {
		t.Error("Expected key3 to be evicted")
	}

	// Others should exist
	if _, ok := cache.Get("key1"); !ok {
		t.Error("Expected key1 to exist")
	}
	if _, ok := cache.Get("key2"); !ok {
		t.Error("Expected key2 to exist")
	}
	if _, ok := cache.Get("key4"); !ok {
		t.Error("Expected key4 to exist")
	}
}

func TestDifferentTypes(t *testing.T) {
	// Test with string keys and string values
	stringCache := New[string, string](2)
	stringCache.Add("hello", "world")

	value, ok := stringCache.Get("hello")
	if !ok || value != "world" {
		t.Error("String cache failed")
	}

	// Test with int keys and string values
	intCache := New[int, string](2)
	intCache.Add(42, "answer")

	value2, ok := intCache.Get(42)
	if !ok || value2 != "answer" {
		t.Error("Int key cache failed")
	}
}

func TestSingleCapacity(t *testing.T) {
	cache := New[string, int](1)

	cache.Add("key1", 1)
	if _, ok := cache.Get("key1"); !ok {
		t.Error("Expected key1 to exist")
	}

	// Add second item - should evict first
	cache.Add("key2", 2)

	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be evicted")
	}
	if _, ok := cache.Get("key2"); !ok {
		t.Error("Expected key2 to exist")
	}

	if len(cache.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(cache.items))
	}
}

func TestZeroValues(t *testing.T) {
	cache := New[string, int](2)

	// Add zero value
	cache.Add("zero", 0)

	value, ok := cache.Get("zero")
	if !ok {
		t.Error("Expected to find zero key")
	}
	if value != 0 {
		t.Errorf("Expected zero value, got %d", value)
	}
}

// Benchmark tests
func BenchmarkAdd(b *testing.B) {
	cache := New[int, int](1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Add(i%1000, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cache := New[int, int](1000)

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Add(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkMixed(b *testing.B) {
	cache := New[int, int](1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			cache.Add(i%1000, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}
