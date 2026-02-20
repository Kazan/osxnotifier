package main

import (
	"strings"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
)

type fakeInformationDialog struct {
	onClosed func()
}

func (d *fakeInformationDialog) Show() {}

func (d *fakeInformationDialog) Hide() {}

func (d *fakeInformationDialog) SetDismissText(label string) {}

func (d *fakeInformationDialog) SetOnClosed(closed func()) {
	d.onClosed = closed
}

func (d *fakeInformationDialog) Refresh() {}

func (d *fakeInformationDialog) Resize(size fyne.Size) {}

func (d *fakeInformationDialog) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (d *fakeInformationDialog) Dismiss() {
	if d.onClosed != nil {
		d.onClosed()
	}
}

func TestAboutDialogContent(t *testing.T) {
	t.Helper()

	if appName != "OSX Notifier" {
		t.Fatalf("expected appName to be %q, got %q", "OSX Notifier", appName)
	}

	if appVersion != "1.0.0" {
		t.Fatalf("expected appVersion to be %q, got %q", "1.0.0", appVersion)
	}

	if appDescription != "A macOS menu bar application" {
		t.Fatalf("expected appDescription to be %q, got %q", "A macOS menu bar application", appDescription)
	}

	content := aboutDialogContent()
	if !strings.Contains(content, "OSX Notifier") {
		t.Fatalf("expected dialog content to contain %q", "OSX Notifier")
	}

	if !strings.Contains(content, appVersion) {
		t.Fatalf("expected dialog content to contain app version %q", appVersion)
	}

	if !strings.Contains(content, appDescription) {
		t.Fatalf("expected dialog content to contain app description %q", appDescription)
	}
}

func TestAboutDialogCallable(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("showAboutDialog panicked: %v", recovered)
		}
	}()

	done := make(chan struct{})
	go func() {
		showAboutDialog(nil)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("showAboutDialog did not return in time")
	}
}

func TestAboutDialogCallableWithTestApp(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	application := test.NewApp()
	defer application.Quit()

	originalRunOnUIThread := runOnUIThread
	originalNewInformationDialog := newInformationDialog
	t.Cleanup(func() {
		runOnUIThread = originalRunOnUIThread
		newInformationDialog = originalNewInformationDialog
	})

	uiThreadExecuted := make(chan struct{}, 1)
	dialogCreated := make(chan struct{}, 1)
	runOnUIThread = func(fn func()) {
		uiThreadExecuted <- struct{}{}
		fn()
	}
	newInformationDialog = func(title, message string, parent fyne.Window) dialog.Dialog {
		dialogCreated <- struct{}{}
		return &fakeInformationDialog{}
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("showAboutDialog panicked with test app: %v", recovered)
		}
	}()

	showAboutDialog(application)

	select {
	case <-uiThreadExecuted:
	case <-time.After(2 * time.Second):
		t.Fatal("expected showAboutDialog to execute code in fyne.Do callback")
	}

	select {
	case <-dialogCreated:
	case <-time.After(2 * time.Second):
		t.Fatal("expected showAboutDialog to create an information dialog")
	}
}

func TestAboutDialogClosesParentWindowWhenDialogCloses(t *testing.T) {
	t.Helper()
	t.Setenv("FYNE_DRIVER", "test")

	application := test.NewApp()
	defer application.Quit()

	originalRunOnUIThread := runOnUIThread
	originalNewInformationDialog := newInformationDialog
	t.Cleanup(func() {
		runOnUIThread = originalRunOnUIThread
		newInformationDialog = originalNewInformationDialog
	})

	runOnUIThread = func(fn func()) {
		fn()
	}

	fakeDialog := &fakeInformationDialog{}
	newInformationDialog = func(title, message string, parent fyne.Window) dialog.Dialog {
		return fakeDialog
	}

	initialWindows := len(application.Driver().AllWindows())

	showAboutDialog(application)

	if got := len(application.Driver().AllWindows()); got != initialWindows+1 {
		t.Fatalf("expected window count to increase by 1 after showing About dialog, got %d (initial %d)", got, initialWindows)
	}

	fakeDialog.Dismiss()

	deadline := time.Now().Add(2 * time.Second)
	for {
		if len(application.Driver().AllWindows()) == initialWindows {
			break
		}

		if time.Now().After(deadline) {
			t.Fatal("expected About parent window to close when dialog closes")
		}

		time.Sleep(10 * time.Millisecond)
	}
}
