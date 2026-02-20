package main

import "testing"

func TestCreateMenuDoesNotPanic(t *testing.T) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("createMenu panicked: %v", recovered)
		}
	}()

	createMenu()
}

func TestCreateMenuCanBeCalledRepeatedly(t *testing.T) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("createMenu panicked on repeated calls: %v", recovered)
		}
	}()

	createMenu()
	createMenu()
}
