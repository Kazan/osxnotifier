package main

import (
	"bytes"
	"strings"
	"testing"

	"fyne.io/fyne/v2"
)

type mockDesktopApp struct {
	fyne.App
	setSystemTrayIconCalled bool
	setSystemTrayIconArg    fyne.Resource
}

func (m *mockDesktopApp) SetSystemTrayMenu(*fyne.Menu) {}

func (m *mockDesktopApp) SetSystemTrayIcon(icon fyne.Resource) {
	m.setSystemTrayIconCalled = true
	m.setSystemTrayIconArg = icon
}

func (m *mockDesktopApp) SetSystemTrayWindow(fyne.Window) {}

func TestIconResource(t *testing.T) {
	t.Helper()

	resource := getIconResource()
	if resource == nil {
		t.Fatal("expected icon resource to be non-nil")
	}

	content := resource.Content()
	if len(content) == 0 {
		t.Fatal("expected icon resource content to be non-empty")
	}

	trimmedContent := bytes.TrimSpace(content)
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	isPNG := len(trimmedContent) >= len(pngSignature) && bytes.Equal(trimmedContent[:len(pngSignature)], pngSignature)
	isSVG := bytes.HasPrefix(trimmedContent, []byte("<svg")) || bytes.Contains(trimmedContent, []byte("<svg"))
	if !isPNG && !isSVG {
		t.Fatal("expected icon resource content to be a valid PNG or SVG")
	}
}

func TestSystemTraySetup(t *testing.T) {
	t.Helper()
	application := &mockDesktopApp{}

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("system tray setup panicked: %v", recovered)
		}
	}()

	setupSystemTray(application)

	if !application.setSystemTrayIconCalled {
		t.Fatal("expected SetSystemTrayIcon to be called")
	}

	if application.setSystemTrayIconArg == nil {
		t.Fatal("expected SetSystemTrayIcon to receive a non-nil icon")
	}
}

func TestIconIsTemplateStyled(t *testing.T) {
	t.Helper()

	content := strings.ToLower(string(getIconResource().Content()))
	if strings.Contains(content, `fill="black"`) {
		t.Fatal("expected icon SVG to be template-styled without hard-coded black fill")
	}
}
