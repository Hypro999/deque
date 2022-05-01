package deque

import (
	"testing"
)

type basicType interface {
	bool |
		string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

func assertEqual[T basicType](a T, b T, t *testing.T) {
	if a != b {
		t.Errorf("want: %v, got: %v", a, b)
	}
}

func TestRemoveEmpty(t *testing.T) {
	var d Deque[int]

	_, err := d.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst did not produce an error with empty deque")
	}

	_, err = d.RemoveLast()
	if err == nil {
		t.Error("RemoveLast did not produce an error with empty deque")
	}

	assertEqual(d.size, 0, t)
}

func TestAddFirstGet(t *testing.T) {
	var d Deque[int]

	d.AddFirst(1)

	v, err := d.GetFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	v, err = d.GetLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	d.AddFirst(2)

	v, err = d.GetFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 2, t)

	v, err = d.GetLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 2, t)
}
func TestAddLastGet(t *testing.T) {
	var d Deque[int]

	d.AddLast(1)

	v, err := d.GetFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	v, err = d.GetLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	d.AddLast(2)

	v, err = d.GetFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 2, t)

	v, err = d.GetLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 2, t)
}

func TestAddFirstRemoveFirst(t *testing.T) {
	var d Deque[int]

	d.AddFirst(1)
	d.AddFirst(2)

	v, err := d.RemoveFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 1, t)

	v, err = d.RemoveFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 0, t)

	_, err = d.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst did not produce an error with empty deque")
	}
}
func TestAddLastRemoveLast(t *testing.T) {
	var d Deque[int]

	d.AddLast(1)
	d.AddLast(2)

	v, err := d.RemoveLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 1, t)

	v, err = d.RemoveLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 0, t)

	_, err = d.RemoveLast()
	if err == nil {
		t.Error("RemoveLast did not produce an error with empty deque")
	}
}

func TestAddFirstRemoveLast(t *testing.T) {
	var d Deque[int]

	d.AddFirst(1)
	d.AddFirst(2)

	v, err := d.RemoveLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	v, err = d.RemoveLast()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 0, t)

	_, err = d.RemoveLast()
	if err == nil {
		t.Error("RemoveLast did not produce an error with empty deque")
	}
}

func TestAddLastRemoveFirst(t *testing.T) {
	var d Deque[int]

	d.AddLast(1)
	d.AddLast(2)

	v, err := d.RemoveFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 1, t)
	assertEqual(d.size, 1, t)

	v, err = d.RemoveFirst()
	if err != nil {
		t.Error("expected err == nil")
	}
	assertEqual(v, 2, t)
	assertEqual(d.size, 0, t)

	_, err = d.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst did not produce an error with empty deque")
	}
}
