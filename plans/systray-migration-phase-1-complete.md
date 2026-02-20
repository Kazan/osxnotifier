## Phase 1 Complete: Update Dependencies and Create Icon Asset

Successfully migrated from Fyne to systray library by updating dependencies and creating a PNG icon asset for menu bar use. All tests pass and the foundation is set for the remaining migration phases.

**Files created/changed:**
- bell-icon.png (NEW)
- generate_icon.go (NEW - icon generator utility)
- icon.go
- icon_test.go
- main_test.go
- go.mod
- go.sum

**Functions created/changed:**
- getIconData() - NEW: Returns embedded PNG icon as byte array
- TestIconDataNotEmpty() - NEW: Validates icon data is loaded
- TestIconIsValidPNG() - NEW: Validates PNG format
- TestModuleDependenciesUpdated() - NEW: Verifies systray dependency

**Tests created/changed:**
- TestIconDataNotEmpty - Validates icon data is not empty
- TestIconIsValidPNG - Validates PNG signature and format
- TestModuleDependenciesUpdated - Verifies fyne.io/systray is direct dependency
- Removed: TestModuleDependencies (old Fyne check)

**Review Status:** APPROVED

**Git Commit Message:**
```
feat: Add systray dependency and PNG icon for migration

- Add fyne.io/systray v1.12.0 as direct dependency
- Create 22x22px bell-icon.png for menu bar (template-styled)
- Add getIconData() function with //go:embed support
- Add PNG validation tests (TestIconDataNotEmpty, TestIconIsValidPNG)
- Update module dependency tests to verify systray
- Generate bell icon programmatically with black on transparent
```
