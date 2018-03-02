package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t1 := NewTransaction("income", "10.0")
	t2 := NewTransaction("income", "-10.0")
	t3 := t1.Add(t2)
	if t3.Amount.StringFixed(2) != "0.00" {
		t.Error(t3.Amount)
	}
}

func TestAdd2(t *testing.T) {
	t1 := NewTransaction("income1", "10.34")
	t2 := NewTransaction("income2", "15.12")
	t3 := t1.Add(t2)
	if t3.Amount.StringFixed(2) != "25.46" {
		t.Error(t3.Amount)
	}
}
