package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//go:embed web
var webFS embed.FS

var (
	mu       sync.Mutex
	lastPing = time.Now()
	gotPing  bool
)

func dataDir() string {
	base, err := os.UserConfigDir()
	if err != nil {
		base = os.TempDir()
	}
	d := filepath.Join(base, "RaidReady")
	_ = os.MkdirAll(d, 0o755)
	return d
}

func main() {
	dir := dataDir()
	dataFile := filepath.Join(dir, "data.json")

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fatal("Could not start the local game server:\n" + err.Error())
	}
	port := ln.Addr().(*net.TCPAddr).Port

	sub, _ := fs.Sub(webFS, "web")
	mux := http.NewServeMux()
	static := http.FileServer(http.FS(sub))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		static.ServeHTTP(w, r)
	})
	mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		mu.Lock()
		defer mu.Unlock()
		switch r.Method {
		case http.MethodGet:
			b, err := os.ReadFile(dataFile)
			if err != nil || !json.Valid(b) {
				b = []byte("{}")
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		case http.MethodPost:
			b, err := io.ReadAll(io.LimitReader(r.Body, 4<<20))
			if err != nil || !json.Valid(b) {
				http.Error(w, "bad json", 400)
				return
			}
			tmp := dataFile + ".tmp"
			if err := os.WriteFile(tmp, b, 0o644); err == nil {
				_ = os.Rename(tmp, dataFile)
			}
			w.WriteHeader(204)
		default:
			w.WriteHeader(405)
		}
	})
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		lastPing, gotPing = time.Now(), true
		mu.Unlock()
		w.WriteHeader(204)
	})
	go http.Serve(ln, mux)

	url := fmt.Sprintf("http://127.0.0.1:%d/", port)
	proc, err := openApp(url, filepath.Join(dir, "window"))
	if err != nil {
		fatal("Could not open the game window:\n" + err.Error())
	}

	// If we launched our own app window and it stays alive, quit when it closes.
	start := time.Now()
	if proc != nil {
		done := make(chan struct{})
		go func() { proc.Wait(); close(done) }()
		select {
		case <-done:
			// handed off to an existing browser process -> fall back to heartbeat
		case <-time.After(6 * time.Second):
			<-done
			return
		}
	}
	for {
		time.Sleep(2 * time.Second)
		mu.Lock()
		idle := time.Since(lastPing)
		ok := gotPing
		mu.Unlock()
		if (ok && idle > 150*time.Second) || (!ok && time.Since(start) > 180*time.Second) {
			return
		}
	}
}
