package skiplist

import "testing"

func TestInsert(t *testing.T) {
	sl := NewSkipList[int, string]()
	sl.Insert(10, "test1")
	sl.Insert(20, "test2")

	if val, ok := sl.Search(20); val != "test2" || !ok {
		t.Errorf("wrong value")
	}
}

func TestDelete(t *testing.T) {

	sl := NewSkipList[int, string]()
	sl.Insert(10, "test1")
	sl.Insert(30, "test1")
	sl.Insert(20, "test2")
	sl.Insert(40, "test2")

	sl.Delete(40)
	if _, ok := sl.Search(40); ok {
		t.Error("should not exists.")
	}
}

func TestSearch(t *testing.T) {
	sl := NewSkipList[int, int]()
	for i := 0; i < 1000000; i++ {
		sl.Insert(i, i)
	}

	if v, ok := sl.Search(40); v != 40 || !ok {
		t.Error("wrong value")
	}

	sl.Delete(40)
	if _, ok := sl.Search(40); ok {
		t.Error("should not exists.")
	}

}

func BenchmarkInsert(b *testing.B) {
	sl := NewSkipList[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Insert(i, i)
	}
}

func BenchmarkSearch(b *testing.B) {

	sl := NewSkipList[int, int]()
	for i := 0; i < 1000000; i++ {
		sl.Insert(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Search(i)
	}
}
