package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	appName        = "OSX Notifier"
	appVersion     = "1.0.0"
	appDescription = "A macOS menu bar application"
)

var runOsaScript = func(script string) error {
	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

func escapeAppleScriptString(value string) string {
	escaped := strings.ReplaceAll(value, "\\", "\\\\")
	return strings.ReplaceAll(escaped, "\"", "\\\"")
}

func showAboutDialog() {
	message := fmt.Sprintf("%s\\n\\nVersion: %s\\n\\n%s", appName, appVersion, appDescription)

	script := fmt.Sprintf(
		`display dialog "%s" with title "About %s" buttons {"OK"} default button "OK"`,
		escapeAppleScriptString(message),
		escapeAppleScriptString(appName),
	)

	_ = runOsaScript(script)
}
