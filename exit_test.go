package main

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestExitFunction(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	application := initializeApp()
	defer application.Quit()

	originalQuitApp := quitApp
	t.Cleanup(func() {
		quitApp = originalQuitApp
	})

	called := false
	quitApp = func(applicationToQuit fyne.App) {
		if applicationToQuit == nil {
			t.Fatal("expected non-nil app passed to quitApp")
		}
		called = true
	}

	exitApp(application)

	if !called {
		t.Fatal("expected exitApp to call quitApp")
	}
}

func TestExitMenuWiring(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	application := initializeApp()
	defer application.Quit()

	menu := createMenu(application)
	if menu == nil {
		t.Fatal("expected menu to be non-nil")
	}

	if len(menu.Items) < 2 {
		t.Fatalf("expected menu to contain Exit item, got %d items", len(menu.Items))
	}

	exitItem := menu.Items[1]
	if exitItem == nil {
		t.Fatal("expected Exit menu item to be non-nil")
	}

	if exitItem.Label != "Exit" {
		t.Fatalf("expected second menu item label to be Exit, got %q", exitItem.Label)
	}

	if exitItem.Action == nil {
		t.Fatal("expected Exit menu item action to be set")
	}

	originalQuitApp := quitApp
	t.Cleanup(func() {
		quitApp = originalQuitApp
	})

	called := false
	quitApp = func(applicationToQuit fyne.App) {
		if applicationToQuit == nil {
			t.Fatal("expected non-nil app passed from Exit menu action")
		}
		called = true
	}

	exitItem.Action()

	if !called {
		t.Fatal("expected Exit menu action to call exitApp")
	}
}
