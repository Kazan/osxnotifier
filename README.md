# OSX Notifier

A lightweight macOS menu bar application written in Go that displays a bell icon in the menu bar with "About" and "Quit" options. Built with the systray library for true menu bar-only behavior.

## Features

✅ **True Menu Bar-Only App** - No Dock icon, runs silently in the menu bar
✅ **Bell Icon** - Template-styled PNG icon that adapts to dark/light mode
✅ **Native Dialogs** - Uses macOS native dialogs via osascript
✅ **Lightweight** - Minimal dependencies (~10 packages)
✅ **Clean Architecture** - Channel-based event handling
✅ **Comprehensive Tests** - Full test coverage with TDD approach

---

## Building & Running

### Development
```bash
# Run tests
make test

# Build binary only
make build

# Build app bundle
make app

# Build and run
make run

# Clean build artifacts
make clean
```

### Running the App
```bash
open osxnotifier.app
```

To stop the app, click the bell icon and select "Quit".

---

## Project Structure

```
osxnotifier/
├── main.go              # App initialization and lifecycle
├── icon.go              # Bell icon (template-styled SVG)
├── menu.go              # Menu creation and items
├── about.go             # About dialog logic
├── main_test.go         # App initialization tests
├── icon_test.go         # Icon and system tray tests
├── menu_test.go         # Menulifecycle using systray.Run()
├── icon.go              # Bell icon (embedded PNG)
├── menu.go              # Channel-based menu handling
├── about.go             # Native osascript dialogs
├── bell-icon.png        # 22x22px template-styled icon
├── generate_icon.go     # Icon generator utility
├── main_test.go         # Lifecycle tests
├── icon_test.go         # Icon validation tests
├── menu_test.go         # Menu structure tests
├── about_test.go        # Dialog tests
├── Info.plist           # macOS app bundle config (LSUIElement=true)
├── Makefile             # Build automation
└── plans/               # Implementation documentation
**Current:**
- **Language:** Go 1.26
- **Menu Bar Library:** fyne.io/systray v1.12.0
- **Icon:** Template-styled PNG (22x22px, black on transparent)
- **Dialog:** Native macOS osascript
- **Dependencies:** Minimal (~10 packages total)
---

## Development History

### Completed Phases

1. ✅ **Phase 1:** Setup Dependencies and Basic App Structure
2. ✅ **Phase 2:** Implement Menu Bar Icon
3. Architecture

### Lifecycle
- Uses `systray.Run(onReady, onExit)` for clean lifecycle management
- No hidden windows or Dock icons
- Minimal resource footprint

### Menu System
- Channel-based event handling via `systray.AddMenuItem()`
- Goroutine listens to click channels using select statement
- Clean separation between UI and business logic

### Icon
- 22x22px PNG embedded with `//go:embed`
- Template-styled (black on transparent) for automatic dark/light mode adaptation
- Generated programmatically with `generate_icon.go`

### Dialogs
- Native macOS dialogs via `osascript` command
- No external dependencies for UI elements
- Consistent with system appearance
```bash
make test
# or
go test ./...
```

**Test Coverage:**
- App initialization (headless compatible)
- Icon resource validation and template styling
- System tray setup
- Menu structure and items
- About dialog content and execution
- Window lifecycle management

---

## Contributing

When making changes:

1. Write tests first (TDD approach)
2. Run `make test` before committing
3. Run `make build` to verify compilation
Run all tests with:
```bash
make test
# or
go test ./...
```

**Test Coverage:**
- Lifecycle functions (onReady, onExit)
- Icon validation (PNG format, embedded data)
- Menu creation and structure
- About dialog content and osascript formatting
- All tests pass without requiring UI interaction

**Test-Driven Development:**
This project was built using strict TDD principles - all tests were written before implementation.
Development History

This app was originally built with Fyne framework and later migrated to systray for true menu bar-only behavior. The migration eliminated the Dock icon issue and reduced the dependency footprint by ~70%.

**Migration Documentation:**
- [Initial Plan](plans/osx-menu-bar-app-plan.md)
- [Migration Plan](plans/systray-migration-plan.md)
- [Migration Complete](plans/systray-migration-complete.md)

## Contributing

When making changes:

1. Write tests first (TDD approach)
2. Run `make test` to verify all tests pass
3. Run `make build` to verify compilation
4. Test manually with `open osxnotifier.app`
5. Update documentation as needed

## License

[Add your license here]
