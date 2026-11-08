//go:build !windows

package main

import (
	"fmt"
	"os"
)

func openApp(url, profile string) (*os.Process, error) {
	fmt.Println("RAID//READY running at", url)
	return nil, nil
}

func openURL(u string) error {
	fmt.Println("Open in your browser:", u)
	return nil
}

func fatal(msg string) { fmt.Fprintln(os.Stderr, msg); os.Exit(1) }
