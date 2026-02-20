.PHONY: build clean test app run

# Build the binary
build:
	go build -o osxnotifier .

# Build as a macOS app bundle (menu bar only, no Dock icon)
app: build
	mkdir -p osxnotifier.app/Contents/MacOS
	mkdir -p osxnotifier.app/Contents/Resources
	cp osxnotifier osxnotifier.app/Contents/MacOS/
	cp Info.plist osxnotifier.app/Contents/Info.plist

# Run the app bundle
run: app
	open osxnotifier.app

clean:
	rm -f osxnotifier
	rm -rf osxnotifier.app

test:
	go test ./...
