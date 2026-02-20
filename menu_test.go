package main

import "testing"

func TestMenuCreation(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	menu := createMenu(nil)
	if menu == nil {
		t.Fatal("expected menu to be non-nil")
	}

	if len(menu.Items) != 1 {
		t.Fatalf("expected menu to have exactly 1 item, got %d", len(menu.Items))
	}
}

func TestMenuItems(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	menu := createMenu(nil)
	if menu == nil {
		t.Fatal("expected menu to be non-nil")
	}

	if len(menu.Items) != 1 {
		t.Fatalf("expected menu to have exactly 1 item, got %d", len(menu.Items))
	}

	if menu.Items[0] == nil {
		t.Fatal("expected menu item to be non-nil")
	}

	if menu.Items[0].Label != "About" {
		t.Fatalf("expected menu item label to be About, got %q", menu.Items[0].Label)
	}

	if menu.Items[0].Action == nil {
		t.Fatal("expected About menu item action to be set")
	}
}
