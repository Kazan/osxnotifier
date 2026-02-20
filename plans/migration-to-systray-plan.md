## Migration Plan: Switch from Fyne to fyne.io/systray Library

### Why This Migration?

**Current Issue:**
- Fyne is a full GUI framework that creates application windows by default
- Even with `LSUIElement=true`, Fyne shows a Dock icon because it needs window management
- Fyne was designed for full desktop apps, not menu bar-only applications

**Solution:**
- Switch to `fyne.io/systray` - a lightweight library specifically designed for system tray/menu bar applications
- No Dock icon, no windows, no app activation - pure menu bar functionality
- Used by many macOS menu bar apps and maintains the same underlying macOS APIs

---

## Implementation Plan

### Phase 1: Update Dependencies

**Objective:** Replace Fyne with systray library

**Files to Modify:**
- `go.mod` - Remove Fyne, add systray

**Steps:**
1. Remove Fyne dependency:
   ```bash
   go mod edit -droprequire fyne.io/fyne/v2
   ```

2. Add systray dependency:
   ```bash
   go get fyne.io/systray
   ```

3. Run `go mod tidy`

**Expected Changes:**
- `go.mod` will have `fyne.io/systray` instead of `fyne.io/fyne/v2`
- Much smaller dependency tree (systray is lightweight)

---

### Phase 2: Rewrite Main Application Structure

**Objective:** Convert from Fyne's app.Run() to systray.Run()

**Files to Modify:**
- `main.go` - Complete rewrite of initialization and lifecycle

**Current Structure:**
```go
func main() {
    application := initializeApp()
    hiddenWindow := application.NewWindow("OSX Notifier")
    setupSystemTray(application, hiddenWindow)
    application.Run()
}
```

**New Structure:**
```go
func main() {
    systray.Run(onReady, onExit)
}

func onReady() {
    // Set icon
    systray.SetTemplateIcon(getIconData(), getIconData())
    systray.SetTitle("OSX Notifier")
    systray.SetTooltip("OSX Notifier")
    
    // Create menu items
    createMenu()
}

func onExit() {
    // Cleanup code if needed
}
```

**Key Differences:**
- No app object, no windows
- `systray.Run()` takes two functions: onReady (setup) and onExit (cleanup)
- Icon is set directly with `systray.SetTemplateIcon()` or `systray.SetIcon()`
- Menu items are created with `systray.AddMenuItem()`

**Implementation Details:**
- `getIconData()` should return `[]byte` of the icon (PNG format works best)
- Use `systray.SetTemplateIcon()` for dark/light mode adaptation
- Menu creation is imperative, not declarative like Fyne

---

### Phase 3: Rewrite Icon Handling

**Objective:** Convert icon from Fyne Resource to raw byte array

**Files to Modify:**
- `icon.go` - Change from SVG string to byte array function

**Current Approach:**
```go
func getIconResource() fyne.Resource {
    return fyne.NewStaticResource("bell-icon", []byte(svgContent))
}
```

**New Approach:**
```go
func getIconData() []byte {
    // Return PNG byte array for the bell icon
    // Can embed a PNG file or generate programmatically
    return []byte{...}
}
```

**Options for Icon:**
1. **Embed PNG file:** Use `//go:embed` to embed a PNG file
2. **Inline PNG data:** Base64 decode or literal byte array
3. **Generate programmatically:** Use image/draw to create icon

**Recommendation:** Create a simple 22x22px PNG icon and embed it:
```go
import _ "embed"

//go:embed bell-icon.png
var iconData []byte

func getIconData() []byte {
    return iconData
}
```

**Template Icon:**
- For dark/light adaptation, use a monochrome PNG (black icon on transparent background)
- Call `systray.SetTemplateIcon(iconData, iconData)` (same data for both normal and dark mode)

---

### Phase 4: Rewrite Menu System

**Objective:** Convert from Fyne's declarative menu to systray's imperative menu

**Files to Modify:**
- `menu.go` - Rewrite menu creation logic

**Current Fyne Approach:**
```go
func createMenu(application fyne.App) *fyne.Menu {
    aboutItem := fyne.NewMenuItem("About", func() {
        showAboutDialog(application)
    })
    return fyne.NewMenu("", aboutItem)
}
```

**New systray Approach:**
```go
func createMenu() {
    // Add About menu item
    mAbout := systray.AddMenuItem("About", "Show application information")
    
    // Add separator
    systray.AddSeparator()
    
    // Quit is automatically added by systray on macOS
    // Or manually add:
    mQuit := systray.AddMenuItem("Quit", "Quit the application")
    
    // Handle menu item clicks in goroutines
    go func() {
        for {
            select {
            case <-mAbout.ClickedCh:
                showAboutDialog()
            case <-mQuit.ClickedCh:
                systray.Quit()
            }
        }
    }()
}
```

**Key Differences:**
- Menu items return a `*systray.MenuItem` with a `ClickedCh` channel
- Must listen to click channels in goroutines
- Separator is a separate function call
- Quit menu item is often auto-added on macOS, but can be added manually

**Important:** systray automatically adds a "Quit" menu item on macOS, so you don't need to add it manually unless you want custom behavior.

---

### Phase 5: Rewrite About Dialog

**Objective:** Create About dialog without Fyne's dialog system

**Files to Modify:**
- `about.go` - Complete rewrite using native dialogs or alternative approach

**Challenge:** systray has no built-in dialog system since it's menu bar only.

**Options:**

**Option 1: Native macOS Dialog (Recommended)**
Use `osascript` to show native macOS dialogs:
```go
func showAboutDialog() {
    script := fmt.Sprintf(`display dialog "%s\nVersion: %s\n%s" with title "%s" buttons {"OK"} default button "OK"`,
        appDescription, appVersion, "", appName)
    
    cmd := exec.Command("osascript", "-e", script)
    cmd.Run()
}
```

**Option 2: Open Browser with About Page**
```go
func showAboutDialog() {
    // Open a simple HTML page with about info
    url := "https://yourdomain.com/about"
    cmd := exec.Command("open", url)
    cmd.Run()
}
```

**Option 3: Terminal Notification**
```go
func showAboutDialog() {
    script := fmt.Sprintf(`display notification "%s\nVersion: %s" with title "%s"`,
        appDescription, appVersion, appName)
    
    cmd := exec.Command("osascript", "-e", script)
    cmd.Run()
}
```

**Recommendation:** Use Option 1 (osascript dialog) for a native macOS experience without dependencies.

---

### Phase 6: Update Tests

**Objective:** Rewrite tests for systray architecture

**Files to Modify:**
- `main_test.go` - Update for systray initialization
- `menu_test.go` - Update for channel-based menu items
- `about_test.go` - Update for osascript-based dialogs
- `icon_test.go` - Update for byte array icon

**Testing Challenges:**
- systray.Run() is blocking, so testing the full lifecycle is harder
- Many systray functions interact with the OS, making unit tests difficult

**Testing Strategy:**

1. **Test Icon Data:**
```go
func TestIconData(t *testing.T) {
    data := getIconData()
    if len(data) == 0 {
        t.Fatal("icon data should not be empty")
    }
    // Verify it's valid PNG
    _, err := png.Decode(bytes.NewReader(data))
    if err != nil {
        t.Fatalf("icon data is not valid PNG: %v", err)
    }
}
```

2. **Test Menu Item Creation (using test seams):**
```go
var addMenuItem = systray.AddMenuItem

func TestMenuItemCreation(t *testing.T) {
    // Test that menu items would be created correctly
    // This is more of an integration test
}
```

3. **Test About Dialog Command:**
```go
func TestAboutDialogCommand(t *testing.T) {
    // Test that the osascript command is correctly formatted
    // Don't actually execute it
    script := getAboutDialogScript()
    if !strings.Contains(script, appName) {
        t.Error("dialog script should contain app name")
    }
}
```

**Note:** Full integration testing will require actually running the app and manually verifying behavior.

---

### Phase 7: Update Build System

**Objective:** Ensure Makefile and Info.plist still work correctly

**Files to Verify:**
- `Makefile` - Should work as-is (no changes needed)
- `Info.plist` - Keep `LSUIElement=true` (still needed and will now work properly)
- `.gitignore` - Already correct

**No Changes Needed:**
- Build process remains the same
- App bundle structure remains the same
- `LSUIElement=true` will now work correctly with systray

---

### Phase 8: Remove Unused Files

**Objective:** Clean up Fyne-specific test files and code

**Files to Remove:**
- Any Fyne-specific mock or test doubles
- Fyne window lifecycle tests

**Files to Review:**
- Check if any test helpers are Fyne-specific and need removal

---

## Complete File-by-File Changes

### go.mod
**Before:**
```
module osxnotifier

go 1.23

require fyne.io/fyne/v2 v2.x.x
```

**After:**
```
module osxnotifier

go 1.23

require fyne.io/systray v1.11.0
```

---

### main.go
**Before:**
```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/driver/desktop"
    "fyne.io/fyne/v2/theme"
)

func initializeApp() fyne.App {
    return app.NewWithID("com.osxnotifier.app")
}

func main() {
    application := initializeApp()
    hiddenWindow := application.NewWindow("OSX Notifier")
    // ... window setup
    setupSystemTray(application, hiddenWindow)
    application.Run()
}
```

**After:**
```go
package main

import (
    "fyne.io/systray"
)

func main() {
    systray.Run(onReady, onExit)
}

func onReady() {
    // Set the icon (template icon for dark/light mode)
    systray.SetTemplateIcon(getIconData(), getIconData())
    systray.SetTitle("OSX Notifier")
    systray.SetTooltip("OSX Notifier - Click for menu")
    
    // Create menu
    createMenu()
}

func onExit() {
    // Cleanup code if needed
}
```

---

### icon.go
**Before:**
```go
package main

import "fyne.io/fyne/v2"

func getIconResource() fyne.Resource {
    // SVG content
    return fyne.NewStaticResource("bell-icon", []byte(svgContent))
}
```

**After:**
```go
package main

import _ "embed"

//go:embed bell-icon.png
var iconData []byte

func getIconData() []byte {
    return iconData
}
```

**Also Create:** `bell-icon.png` - A 22x22px PNG icon (monochrome black bell on transparent background)

---

### menu.go
**Before:**
```go
package main

import "fyne.io/fyne/v2"

func createMenu(application fyne.App) *fyne.Menu {
    aboutItem := fyne.NewMenuItem("About", func() {
        showAboutDialog(application)
    })
    return fyne.NewMenu("", aboutItem)
}
```

**After:**
```go
package main

import "fyne.io/systray"

func createMenu() {
    mAbout := systray.AddMenuItem("About", "Show application information")
    systray.AddSeparator()
    mQuit := systray.AddMenuItem("Quit", "Quit OSX Notifier")
    
    // Handle menu clicks
    go func() {
        for {
            select {
            case <-mAbout.ClickedCh:
                showAboutDialog()
            case <-mQuit.ClickedCh:
                systray.Quit()
            }
        }
    }()
}
```

---

### about.go
**Before:**
```go
package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/dialog"
)

func showAboutDialog(application fyne.App) {
    // Fyne dialog code
}
```

**After:**
```go
package main

import (
    "fmt"
    "os/exec"
)

const (
    appName        = "OSX Notifier"
    appVersion     = "1.0.0"
    appDescription = "A macOS menu bar application"
)

func showAboutDialog() {
    message := fmt.Sprintf("%s\\n\\nVersion: %s\\n\\n%s", 
        appName, appVersion, appDescription)
    
    script := fmt.Sprintf(`display dialog "%s" with title "About %s" buttons {"OK"} default button "OK"`,
        message, appName)
    
    cmd := exec.Command("osascript", "-e", script)
    cmd.Run()
}
```

---

### Icon File Creation

**Create:** `bell-icon.png`

You'll need to create a 22x22px PNG icon. Here are options:

1. **Use an icon generator:** Create a simple bell icon
2. **Use SF Symbols:** macOS includes SF Symbols that can be exported as images
3. **Template format:** Monochrome (black) on transparent background

**Quick option for testing:**
Use a simple ASCII art converted to image, or download a free bell icon from:
- SF Symbols (macOS built-in)
- The Noun Project
- Icons8

**Important:** For template icon (adapts to dark/light mode):
- Use pure black (#000000) for the icon
- Use transparent background
- Keep it simple and recognizable at small sizes

---

## Testing & Verification Checklist

After migration, verify:

- [ ] No Dock icon appears when app is running
- [ ] Bell icon appears in menu bar
- [ ] Icon adapts to dark/light mode (if template icon used)
- [ ] Clicking icon shows dropdown menu
- [ ] "About" menu item works and shows native dialog
- [ ] "Quit" menu item closes the app
- [ ] No app window or menu bar activation occurs
- [ ] App can be launched with `open osxnotifier.app`
- [ ] App runs silently in background
- [ ] `LSUIElement=true` is properly respected

---

## Rollback Plan

If migration fails or has issues:

1. Revert `go.mod` changes
2. Restore original Fyne-based code from git
3. Run `go mod tidy`
4. Rebuild

**Git commands:**
```bash
git checkout go.mod go.sum
git checkout main.go icon.go menu.go about.go
go mod tidy
make clean && make run
```

---

## Expected Benefits

After migration:

✅ **No Dock icon** - True menu bar-only app  
✅ **Lighter weight** - Much smaller binary size  
✅ **No window management** - Simpler architecture  
✅ **Native macOS behavior** - Works like other menu bar apps  
✅ **LSUIElement works** - Properly respected by OS  
✅ **Faster startup** - Less framework overhead  

---

## Resources

- **systray documentation:** https://pkg.go.dev/fyne.io/systray
- **systray GitHub:** https://github.com/fyne-io/systray
- **Example apps:** Many open-source menu bar apps use systray
- **LSUIElement docs:** https://developer.apple.com/documentation/bundleresources/information_property_list/lsuielement
- **osascript docs:** `man osascript` in Terminal

---

## Estimated Effort

- **Phase 1 (Dependencies):** 5 minutes
- **Phase 2 (Main rewrite):** 30 minutes
- **Phase 3 (Icon):** 20 minutes (including PNG creation)
- **Phase 4 (Menu):** 20 minutes
- **Phase 5 (About dialog):** 15 minutes
- **Phase 6 (Tests):** 45 minutes
- **Phase 7-8 (Cleanup):** 10 minutes

**Total:** ~2.5 hours

---

## Notes for Future Session

When starting the migration:

1. **Commit current state first:** `git commit -am "Pre-migration checkpoint"`
2. **Create migration branch:** `git checkout -b migration/systray`
3. **Follow phases in order** - Don't skip ahead
4. **Test after each phase** - Verify builds before proceeding
5. **Keep Info.plist and Makefile** - They're still needed and correct

Good luck with the migration! 🚀
