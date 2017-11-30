package vector

import (
	"fmt"
)

// TODO: resize(), cap()

// TVector -
type TVector struct {
	data []interface{}
	//cap  int
}

// New -
func New() *TVector {
	return &TVector{}
}

// Len -
func (o *TVector) Len() int {
	return len(o.data)
}

// PushFront -
func (o *TVector) PushFront(item interface{}) {
	o.data = append([]interface{}{item}, o.data...)
}

// PopFront -
func (o *TVector) PopFront() (interface{}, error) {
	if len(o.data) == 0 {
		return nil, fmt.Errorf("TVector.PopFront(): storage is empty")
	}
	ret := o.data[0]
	o.data[0] = nil
	o.data = o.data[1:len(o.data)]
	return ret, nil
}

// PushBack -
func (o *TVector) PushBack(item interface{}) {
	o.data = append(o.data, item)
}

// PopBack -
func (o *TVector) PopBack() (interface{}, error) {
	if len(o.data) == 0 {
		return nil, fmt.Errorf("TVector.PopBack(): storage is empty")
	}
	ret := o.data[len(o.data)-1]
	o.data[len(o.data)-1] = nil
	o.data = o.data[:len(o.data)-1]
	return ret, nil
}

// Insert -
func (o *TVector) Insert(i int, item interface{}) error {
	if i < 0 || i > len(o.data)-1 {
		return fmt.Errorf("TVector.Insert(%v, item): index out of range (len: %v)", i, len(o.data))
	}
	o.data = append(o.data[:i], append([]interface{}{item}, o.data[i:]...)...)
	return nil
}

// Remove -
func (o *TVector) Remove(i int) (interface{}, error) {
	if i < 0 || i > len(o.data)-1 {
		return nil, fmt.Errorf("TVector.Remove(%v): index out of range (len: %v)", i, len(o.data))
	}
	ret := o.data[i]
	o.data[i] = nil
	o.data = append(o.data[:i], o.data[i+1:]...)
	return ret, nil
}

// Clear -
func (o *TVector) Clear() {
	o.data = nil
}

// Front -
func (o *TVector) Front() (interface{}, error) {
	if len(o.data) == 0 {
		return nil, fmt.Errorf("TVector.Front(): storage is empty")
	}
	return o.data[0], nil
}

// Back -
func (o *TVector) Back() (interface{}, error) {
	if len(o.data) == 0 {
		return nil, fmt.Errorf("TVector.Back(): storage is empty")
	}
	return o.data[len(o.data)-1], nil
}

// At -
func (o *TVector) At(i int) (interface{}, error) {
	if i < 0 || i > len(o.data)-1 {
		return nil, fmt.Errorf("TVector.At(%v): index out of range (len: %v)", i, len(o.data))
	}
	return o.data[i], nil
}

// IndexOf -
func (o *TVector) IndexOf(item interface{}) int {
	for i := range o.data {
		if item == o.data[i] {
			return i
		}
	}
	return -1
}

// Data -
func (o *TVector) Data() []interface{} {
	return o.data[:]
}
