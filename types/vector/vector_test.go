package vector

import (
	"testing"
)

func TestPushBack(t *testing.T) {
	v := New()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	if val, err := v.At(0); val != 1 || err != nil {
		t.Errorf("Expected: 1, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(1); val != 2 || err != nil {
		t.Errorf("Expected: 2, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(2); val != 3 || err != nil {
		t.Errorf("Expected: 3, <nil> -> got: %v, %v", val, err)
	}
	// index out of range
	if val, err := v.At(3); val != nil || err == nil {
		t.Errorf("Expected: <nil>, !<nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(-1); val != nil || err == nil {
		t.Errorf("Expected: <nil>, !<nil> -> got: %v, %v", val, err)
	}
	// Front/Back
	if val, err := v.Front(); val != 1 || err != nil {
		t.Errorf("Expected: 1, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.Back(); val != 3 || err != nil {
		t.Errorf("Expected: 3, <nil> -> got: %v, %v", val, err)
	}
	// Len/Clear
	if val := v.Len(); val != 3 {
		t.Errorf("Expected: 3 -> got: %v", val)
	}
	v.Clear()
	if val := v.Len(); val != 0 {
		t.Errorf("Expected: 0 -> got: %v", val)
	}
	// PushFront
	v.PushFront("Three")
	v.PushFront("Two")
	v.PushFront("One")
	if val, err := v.At(0); val != "One" || err != nil {
		t.Errorf("Expected: One, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(1); val != "Two" || err != nil {
		t.Errorf("Expected: Two, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(2); val != "Three" || err != nil {
		t.Errorf("Expected: Three, <nil> -> got: %v, %v", val, err)
	}
	// Insert
	if err := v.Insert(2, "Prefinal"); err != nil {
		t.Errorf("Expected: <nil> -> got: %v", err)
	}
	if err := v.Insert(1, "1.5"); err != nil {
		t.Errorf("Expected: <nil> -> got: %v", err)
	}
	if err := v.Insert(0, "Zero"); err != nil {
		t.Errorf("Expected: <nil> -> got: %v", err)
	}
	// Len
	if val := v.Len(); val != 6 {
		t.Errorf("Expected: 6 -> got: %v", val)
	}
	// Insert index out of range
	if err := v.Insert(-1, "xxx"); err == nil {
		t.Errorf("Expected: !<nil> -> got: %v", err)
	}
	if err := v.Insert(6, "xxx"); err == nil {
		t.Errorf("Expected: !<nil> -> got: %v", err)
	}
	// Len
	if val := v.Len(); val != 6 {
		t.Errorf("Expected: 6 -> got: %v", val)
	}
	if val, err := v.Front(); val != "Zero" || err != nil {
		t.Errorf("Expected: 1, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(2); val != "1.5" || err != nil {
		t.Errorf("Expected: 1.5, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.At(4); val != "Prefinal" || err != nil {
		t.Errorf("Expected: Prefinal, <nil> -> got: %v, %v", val, err)
	}
	// Remove
	if val, err := v.Remove(5); val != "Three" || err != nil {
		t.Errorf("Expected: Three, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.Remove(2); val != "1.5" || err != nil {
		t.Errorf("Expected: 1.5, <nil> -> got: %v, %v", val, err)
	}
	if val, err := v.Remove(0); val != "Zero" || err != nil {
		t.Errorf("Expected: Zero, <nil> -> got: %v, %v", val, err)
	}
	// Len
	if val := v.Len(); val != 3 {
		t.Errorf("Expected: 3 -> got: %v", val)
	}
	// Remove out of range
	if val, err := v.Remove(-1); val != nil || err == nil {
		t.Errorf("Expected: <nil>, !<nil> -> got: %v, %v", val, err)
	}
	if val, err := v.Remove(3); val != nil || err == nil {
		t.Errorf("Expected: <nil>, !<nil> -> got: %v, %v", val, err)
	}
	// IndexOf
	if val := v.IndexOf("Prefinal"); val != 2 {
		t.Errorf("Expected: 2 -> got: %v", val)
	}
	if val := v.IndexOf("One"); val != 0 {
		t.Errorf("Expected: 0 -> got: %v", val)
	}
	if val := v.IndexOf("xxxxxxxxxx"); val != -1 {
		t.Errorf("Expected: -1 -> got: %v", val)
	}
	// x := []int{1, 2, 3}
	// fmt.Println("%v", x[0])
	// fmt.Println("%v", x[2:3])
}
