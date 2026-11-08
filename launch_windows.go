//go:build windows

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

func browserCandidates() []string {
	pf, pf86, la := os.Getenv("ProgramFiles"), os.Getenv("ProgramFiles(x86)"), os.Getenv("LOCALAPPDATA")
	var out []string
	for _, base := range []string{pf86, pf, la} {
		if base == "" {
			continue
		}
		out = append(out,
			filepath.Join(base, `Microsoft\Edge\Application\msedge.exe`),
			filepath.Join(base, `Google\Chrome\Application\chrome.exe`),
			filepath.Join(base, `BraveSoftware\Brave-Browser\Application\brave.exe`),
		)
	}
	return out
}

func openApp(url, profile string) (*os.Process, error) {
	for _, exe := range browserCandidates() {
		if _, err := os.Stat(exe); err != nil {
			continue
		}
		cmd := exec.Command(exe,
			"--app="+url,
			"--user-data-dir="+profile,
			"--no-first-run", "--no-default-browser-check",
			"--disable-background-timer-throttling",
			"--disable-renderer-backgrounding",
			"--disable-backgrounding-occluded-windows",
			"--disable-session-crashed-bubble",
			"--autoplay-policy=no-user-gesture-required",
			"--disable-features=Translate",
			"--window-size=1600,900", "--start-maximized",
		)
		if err := cmd.Start(); err == nil {
			return cmd.Process, nil
		}
	}
	// No Chromium browser found: use the default browser.
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	return nil, err
}

// openURL opens a web page in the user's default browser.
func openURL(u string) error {
	return exec.Command("rundll32", "url.dll,FileProtocolHandler", u).Start()
}

func fatal(msg string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	mb := user32.NewProc("MessageBoxW")
	t, _ := syscall.UTF16PtrFromString("RAID//READY")
	m, _ := syscall.UTF16PtrFromString(msg)
	mb.Call(0, uintptr(unsafe.Pointer(m)), uintptr(unsafe.Pointer(t)), 0x10)
	os.Exit(1)
}
