# OSX Notifier

A macOS menu bar application written in Go that displays a bell icon in the top menu bar with a dropdown menu containing "About" and "Quit" options.

## Current Status

**⚠️ Known Issue:** The app currently shows a Dock icon when running, which is not ideal for a menu bar-only application. This is due to Fyne framework limitations.

**✅ Working Features:**
- Bell icon in menu bar (adapts to dark/light mode)
- Dropdown menu with "About" and "Quit" options
- About dialog with app information
- Clean quit functionality
- Comprehensive test suite

**📋 Next Step:** Migrate from Fyne to `fyne.io/systray` library to achieve true menu bar-only behavior without Dock icon.

See: [Migration Plan](plans/migration-to-systray-plan.md)

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
├── menu_test.go         # Menu structure tests
├── about_test.go        # About dialog tests
├── Info.plist           # macOS app bundle configuration (LSUIElement)
├── Makefile             # Build automation
└── plans/               # Implementation and migration plans
    ├── osx-menu-bar-app-plan.md
    ├── osx-menu-bar-app-complete.md
    └── migration-to-systray-plan.md  ← NEXT STEP
```

---

## Technology Stack

**Current:**
- **Language:** Go 1.23
- **GUI Framework:** Fyne v2
- **Icon:** Template-styled SVG with currentColor

**Planned (Migration):**
- **Language:** Go 1.23
- **Menu Bar Library:** fyne.io/systray
- **Icon:** Template-styled PNG
- **Dialog:** Native macOS osascript

---

## Development History

### Completed Phases

1. ✅ **Phase 1:** Setup Dependencies and Basic App Structure
2. ✅ **Phase 2:** Implement Menu Bar Icon
3. ✅ **Phase 3:** Implement Menu with About and Exit Entries
4. ✅ **Phase 4:** Implement About Dialog
5. ✅ **Phase 5:** Implement Exit Functionality (later simplified to use Fyne's built-in Quit)

### Optimizations Made

- Removed duplicate "Exit" menu item (using Fyne's automatic "Quit" instead)
- Fixed orphan window issue in About dialog
- Fixed Fyne threading issues with proper `fyne.Do` usage
- Added `.app` bundle support with `LSUIElement=true`
- Created Makefile for build automation

### Known Limitations (Current Implementation)

- **Dock Icon Appears:** Fyne is a full GUI framework and shows Dock icon even with `LSUIElement=true`
- **Window Management:** Fyne needs windows internally, though hidden from user

---

## Next Steps

### Migration to systray Library

To achieve true menu bar-only behavior (no Dock icon), follow the migration plan:

1. **Read the plan:** [migration-to-systray-plan.md](plans/migration-to-systray-plan.md)
2. **Create checkpoint:** `git commit -am "Pre-migration checkpoint"`
3. **Create branch:** `git checkout -b migration/systray`
4. **Execute phases:** Follow the 8-phase migration plan
5. **Test thoroughly:** Verify no Dock icon appears
6. **Merge:** Once verified, merge back to main

**Estimated time:** ~2.5 hours

---

## Testing

All tests can be run with:
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
4. Test manually with `make run`
5. Update documentation if needed

---

## License

[Add your license here]

---

## Author

[Add your name/info here]
