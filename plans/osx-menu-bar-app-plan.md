## Plan: OSX Menu Bar Application with Icon and Menu

Create a macOS menu bar application using Go and the Fyne toolkit that displays an icon in the top menu bar with a dropdown menu containing "About" (shows app info window) and "Exit" (closes the application) entries.

**Phases: 5 phases**

1. **Phase 1: Setup Dependencies and Basic App Structure**
    - **Objective:** Initialize the Fyne dependency and create the basic application structure with proper initialization
    - **Files/Functions to Modify/Create:**
        - [go.mod](go.mod) - Add Fyne v2 dependency
        - [main.go](main.go) - Create with basic app initialization and main function
    - **Tests to Write:**
        - `TestAppInitialization` - Verify app can be created without panic
        - `TestModuleDependencies` - Verify Fyne dependency is properly loaded
    - **Steps:**
        1. Write test to verify basic app structure initializes without errors
        2. Run tests (expect failure - no implementation yet)
        3. Add Fyne v2 dependency to go.mod using `go get fyne.io/fyne/v2`
        4. Create main.go with app initialization (no UI yet)
        5. Run tests to confirm they pass
        6. Run `go mod tidy` and verify compilation

2. **Phase 2: Implement Menu Bar Icon**
    - **Objective:** Add a system tray icon to the macOS menu bar that is visible and clickable
    - **Files/Functions to Modify/Create:**
        - [main.go](main.go) - Add `setupSystemTray()` function
        - [icon.go](icon.go) - Create with embedded icon data or icon loading logic
    - **Tests to Write:**
        - `TestSystemTraySetup` - Verify system tray can be configured
        - `TestIconResource` - Verify icon resource exists and is valid
    - **Steps:**
        1. Write tests for system tray setup and icon resource validation
        2. Run tests (expect failure)
        3. Create or embed a simple icon (can use Fyne's default or a simple PNG)
        4. Implement `setupSystemTray()` to create the tray icon using Fyne's desktop.App interface
        5. Verify icon appears in menu bar (may require manual verification on macOS)
        6. Run tests to confirm they pass

3. **Phase 3: Implement Menu with About and Exit Entries**
    - **Objective:** Create a dropdown menu with "About" and "Exit" menu items that appear when the tray icon is clicked
    - **Files/Functions to Modify/Create:**
        - [main.go](main.go) - Add `createMenu()` function
        - [menu.go](menu.go) - Create with menu setup logic and menu item definitions
    - **Tests to Write:**
        - `TestMenuCreation` - Verify menu structure is created correctly
        - `TestMenuItems` - Verify "About" and "Exit" menu items exist
    - **Steps:**
        1. Write tests for menu creation and menu item existence
        2. Run tests (expect failure)
        3. Implement `createMenu()` function that creates a Fyne menu
        4. Add "About" and "Exit" as menu items (no actions yet)
        5. Wire menu to system tray icon using `SetSystemTrayMenu()`
        6. Run tests to confirm they pass
        7. Manually verify menu appears on click

4. **Phase 4: Implement About Dialog**
    - **Objective:** Wire the "About" menu item to display a dialog window with application information
    - **Files/Functions to Modify/Create:**
        - [main.go](main.go) or [menu.go](menu.go) - Add `showAboutDialog()` function
        - [about.go](about.go) - Create with about dialog logic and app metadata
    - **Tests to Write:**
        - `TestAboutDialogContent` - Verify about dialog contains correct app information
        - `TestAboutDialogCreation` - Verify about dialog can be created
    - **Steps:**
        1. Write tests for about dialog creation and content
        2. Run tests (expect failure)
        3. Define app metadata (name, version, description)
        4. Implement `showAboutDialog()` using Fyne's `dialog.ShowInformation()`
        5. Wire About menu item click to call `showAboutDialog()`
        6. Run tests to confirm they pass
        7. Manually verify About dialog displays correctly

5. **Phase 5: Implement Exit Functionality**
    - **Objective:** Wire the "Exit" menu item to cleanly terminate the application
    - **Files/Functions to Modify/Create:**
        - [main.go](main.go) or [menu.go](menu.go) - Add `exitApp()` function
    - **Tests to Write:**
        - `TestExitFunction` - Verify exit function exists and is callable
        - `TestExitMenuWiring` - Verify Exit menu item is wired to exit function
    - **Steps:**
        1. Write tests for exit functionality
        2. Run tests (expect failure)
        3. Implement `exitApp()` function that calls appropriate Fyne quit method
        4. Wire Exit menu item click to call `exitApp()`
        5. Run tests to confirm they pass
        6. Manually verify clicking Exit closes the application cleanly
        7. Run final smoke test of complete application

**Decisions Made:**
- **Icon:** Custom bell icon as template (adapts to dark/light mode)
- **About dialog:** App name + version + description
