## Phase 4 Complete: Replace About Dialog with Native macOS Dialog

Successfully converted from Fyne dialog system to native macOS osascript-based dialogs. The About dialog now uses the native macOS dialog API through osascript, eliminating all Fyne dialog dependencies.

**Files created/changed:**
- about.go
- about_test.go
- menu.go

**Functions created/changed:**
- showAboutDialog() - REWRITTEN: No parameters, uses osascript for native macOS dialog
- REMOVED: aboutDialogContent() - Inlined into showAboutDialog()
- REMOVED: runOnUIThread variable (no longer needed)
- REMOVED: newInformationDialog variable (no longer needed)
- ADDED: runOsaScript() - Mockable osascript runner for testability

**Implementation Details:**
- Uses `exec.Command("osascript", "-e", script)` for native dialog
- Message format: "App Name\n\nVersion: X.X.X\n\nDescription"
- Dialog has "OK" button and proper title
- Safe string escaping for AppleScript content
- Synchronous execution (blocks until user clicks OK)

**Tests created/changed:**
- TestAboutDialogContent - UPDATED: Still validates formatting expectations
- TestShowAboutDialogNoParams - NEW: Verifies no-parameter call and osascript invocation
- REMOVED: TestAboutDialogCallable (Fyne-specific)
- REMOVED: TestAboutDialogCallableWithTestApp (Fyne-specific)
- REMOVED: TestAboutDialogClosesParentWindowWhenDialogCloses (Fyne-specific)
- REMOVED: All Fyne dialog mocks and test doubles

**Review Status:** APPROVED

**Git Commit Message:**
```
refactor: Replace Fyne dialog with native macOS osascript

- Rewrite showAboutDialog() to use native osascript dialogs
- Remove all Fyne dialog dependencies from about.go
- Remove showAboutDialog() parameter (no longer needs Fyne App)
- Update menu.go to call showAboutDialog() without parameters
- Add runOsaScript() mockable runner for testability
- Implement safe string escaping for AppleScript content
- Remove Fyne-specific tests and mocks from about_test.go
- Keep content validation tests
```
