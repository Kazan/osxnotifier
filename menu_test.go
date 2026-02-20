package main

import "testing"

func TestMenuCreation(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	menu := createMenu(initializeApp())
	if menu == nil {
		t.Fatal("expected menu to be non-nil")
	}

	if len(menu.Items) != 2 {
		t.Fatalf("expected menu to have exactly 2 items, got %d", len(menu.Items))
	}
}

func TestMenuItems(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	menu := createMenu(initializeApp())
	if menu == nil {
		t.Fatal("expected menu to be non-nil")
	}

	if len(menu.Items) != 2 {
		t.Fatalf("expected menu to have exactly 2 items, got %d", len(menu.Items))
	}

	if menu.Items[0] == nil {
		t.Fatal("expected first menu item to be non-nil")
	}

	if menu.Items[1] == nil {
		t.Fatal("expected second menu item to be non-nil")
	}

	if menu.Items[0].Label != "About" {
		t.Fatalf("expected first menu item label to be About, got %q", menu.Items[0].Label)
	}

	if menu.Items[1].Label != "Exit" {
		t.Fatalf("expected second menu item label to be Exit, got %q", menu.Items[1].Label)
	}

	if menu.Items[0].Action == nil {
		t.Fatal("expected About menu item action to be set")
	}
}
