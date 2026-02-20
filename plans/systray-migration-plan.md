## Plan: Migrate from Fyne to systray for True Menu Bar-Only Behavior

**TL;DR:** Replace Fyne GUI framework with lightweight systray library to eliminate the Dock icon issue and achieve true menu bar-only behavior. The migration involves updating dependencies, rewriting core app structure to use systray's lifecycle, converting icon from SVG to PNG, implementing channel-based menu handling, using native macOS dialogs, and updating all tests to match the new architecture.

**Phases: 5 phases**

### Phase 1: Update Dependencies and Create Icon Asset
- **Objective:** Replace Fyne with systray dependency and create PNG icon asset for menu bar use
- **Files/Functions to Modify/Create:**
  - `go.mod` - Replace Fyne with systray
  - `bell-icon.png` - NEW: 22x22px template-styled PNG icon
  - `icon.go` - Convert from Fyne Resource to byte array with embed
- **Tests to Write:**
  - `TestIconDataNotEmpty` - Verify icon data is loaded
  - `TestIconIsValidPNG` - Verify PNG format validity
  - `TestModuleDependenciesUpdated` - Verify systray in go.mod
- **Steps:**
  1. Write test for icon PNG data validation (`icon_test.go`)
  2. Run tests (should fail - no PNG exists yet)
  3. Create 22x22px bell icon PNG file
  4. Update `icon.go` to use `//go:embed` and return `[]byte`
  5. Run tests (should pass)
  6. Update `go.mod` to remove Fyne and use systray directly
  7. Update module dependency test in `main_test.go`
  8. Run `go mod tidy`
  9. Run all tests (should pass)

### Phase 2: Rewrite Main Application Lifecycle
- **Objective:** Convert from Fyne's app.Run() to systray.Run() with onReady/onExit callbacks
- **Files/Functions to Modify/Create:**
  - `main.go` - Complete rewrite: remove Fyne app, implement systray.Run
  - `main_test.go` - Update app initialization tests
- **Tests to Write:**
  - `TestOnReadyFunction` - Verify onReady sets up menu bar correctly
  - `TestOnExitFunction` - Verify onExit cleanup
- **Steps:**
  1. Write tests for onReady and onExit functions
  2. Run tests (should fail - functions don't exist)
  3. Rewrite `main.go`: remove `initializeApp`, `setupSystemTray`, hidden window logic
  4. Implement `main()` with `systray.Run(onReady, onExit)`
  5. Implement `onReady()`: SetTemplateIcon, SetTitle, SetTooltip, call createMenu
  6. Implement `onExit()`: cleanup code
  7. Run tests (should pass)
  8. Build and manually verify (no Dock icon should appear)

### Phase 3: Rewrite Menu System with Channel-Based Handling
- **Objective:** Convert from Fyne's declarative menu to systray's imperative channel-based menu
- **Files/Functions to Modify/Create:**
  - `menu.go` - Complete rewrite: AddMenuItem, click channel handling
  - `menu_test.go` - Update menu structure tests
- **Tests to Write:**
  - `TestMenuItemsCreated` - Verify menu items are added (mock systray)
  - `TestMenuClickHandling` - Verify click channel goroutine structure
- **Steps:**
  1. Write tests for menu item creation and click handling
  2. Run tests (should fail - new menu structure doesn't exist)
  3. Rewrite `createMenu()`: remove Fyne menu, use `systray.AddMenuItem`
  4. Add About menu item with tooltip
  5. Add separator
  6. Add Quit menu item
  7. Implement goroutine with select statement for click channel handling
  8. Handle `mAbout.ClickedCh` → call `showAboutDialog()`
  9. Handle `mQuit.ClickedCh` → call `systray.Quit()`
  10. Update tests to match new menu structure
  11. Run tests (should pass)

### Phase 4: Replace About Dialog with Native macOS Dialog
- **Objective:** Convert from Fyne dialog to native osascript-based macOS dialog
- **Files/Functions to Modify/Create:**
  - `about.go` - Complete rewrite: remove Fyne dialog, use osascript
  - `about_test.go` - Update dialog tests for osascript approach
- **Tests to Write:**
  - `TestAboutDialogScriptFormat` - Verify osascript command format
  - `TestAboutDialogContent` - Verify dialog message contains app info
  - `TestAboutDialogNoWindow` - Verify no window creation
- **Steps:**
  1. Write tests for osascript command formatting
  2. Run tests (should fail - osascript approach doesn't exist)
  3. Remove Fyne imports and window creation from `showAboutDialog`
  4. Implement dialog message formatting with app info
  5. Implement osascript command construction
  6. Execute osascript with `exec.Command("osascript", "-e", script)`
  7. Remove `runOnUIThread` and `newInformationDialog` vars (no longer needed)
  8. Update tests: remove Fyne dialog mocks, test osascript format
  9. Run tests (should pass)
  10. Manually test dialog appearance

### Phase 5: Clean Up and Verify Complete Migration
- **Objective:** Remove all Fyne remnants, verify tests pass, and ensure no Dock icon appears
- **Files/Functions to Modify/Create:**
  - `icon_test.go` - Remove Fyne mocks and interfaces
  - All test files - Remove Fyne test imports and environment setup
- **Tests to Write:**
  - No new tests, but verify all existing tests pass
- **Steps:**
  1. Remove Fyne-specific test code (mockDesktopApp, Fyne test driver setup)
  2. Remove all `t.Setenv("FYNE_DRIVER", "test")` calls
  3. Run all tests (should pass)
  4. Run `go mod tidy` to clean up unused dependencies
  5. Build app: `make build`
  6. Build app bundle: `make app`
  7. Run app: `open osxnotifier.app`
  8. **Verify**: No Dock icon appears
  9. **Verify**: Bell icon appears in menu bar
  10. **Verify**: Clicking shows menu with About and Quit
  11. **Verify**: About shows native macOS dialog
  12. **Verify**: Quit closes the app
  13. Run full test suite one final time

---

**Open Questions:**
1. **Icon Asset Creation**: Should I generate a simple programmatic bell icon PNG, or would you prefer to provide a specific icon file? (I can create a basic 22x22px bell icon)
2. **Native Dialog Style**: The osascript dialog is functional but basic. Would you like to explore alternatives like opening a simple HTML page in browser, or is the native dialog acceptable?
3. **Test Coverage for systray APIs**: Many systray functions interact directly with macOS. Should we focus on unit tests with mocks, or accept that some integration testing will be manual?
