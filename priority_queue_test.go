package gopq

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	pq := New()
	elements := []int{5, 3, 7, 8, 6, 2, 9}
	for _, e := range elements {
		pq.Insert(e, e)
	}

	sort.Ints(elements)
	for _, e := range elements {
		item := pq.Pop()

		i := item.(int)
		if e != i {
			t.Fatalf("expected %v, got %v", e, i)
		}
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	pq := New()
	pq.Insert("foo", 3)
	pq.Insert("bar", 4)
	pq.UpdatePriority("bar", 2)

	item := pq.Pop()

	if item.(string) != "bar" {
		t.Fatal("priority update failed")
	}
}

func TestPriorityQueueLen(t *testing.T) {
	pq := New()
	if pq.Len() != 0 {
		t.Fatal("empty queue should have length of 0")
	}

	pq.Insert("foo", 1)
	pq.Insert("bar", 1)
	if pq.Len() != 2 {
		t.Fatal("queue should have lenght of 2 after 2 inserts")
	}
}

func TestDoubleAddition(t *testing.T) {
	pq := New()
	pq.Insert("foo", 2)
	pq.Insert("bar", 3)
	pq.Insert("bar", 1)

	if pq.Len() != 2 {
		t.Fatal("queue should ignore inserting the same element twice")
	}

	item := pq.Pop()
	if item.(string) != "foo" {
		t.Fatal("queue should ignore duplicate insert, not update existing item")
	}
}

func TestPopEmptyQueue(t *testing.T) {
	pq := New()

	require.Panics(t, func() { pq.Pop() }, "should panic when performing pop on empty queue")
}

func TestUpdateNonExistingItem(t *testing.T) {
	pq := New()

	pq.Insert("foo", 4)
	pq.UpdatePriority("bar", 5)

	if pq.Len() != 1 {
		t.Fatal("update should not add items")
	}

	item := pq.Pop()
	if item.(string) != "foo" {
		t.Fatalf("update should not overwrite item, expected \"foo\", got \"%v\"", item.(string))
	}
}
