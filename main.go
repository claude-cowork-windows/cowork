package main

/*
#cgo LDFLAGS: -L./fs_agent/target/release -lfs_agent
#include <stdlib.h>

// Declare external C signatures for Rust functions
int safe_write_file(char* root, char* filename, char* content);
char* safe_read_file(char* root, char* filename);
*/
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"unsafe"
)

// RustSafeWrite is a Go wrapper around the Rust `safe_write_file` function.
// It handles C-string memory allocation and cleanup.
func RustSafeWrite(rootDir, filename, content string) error {
	// Convert Go strings to C strings
	cRoot := C.CString(rootDir)
	cFile := C.CString(filename)
	cContent := C.CString(content)
	
	// Free memory when function exits
	defer C.free(unsafe.Pointer(cRoot))
	defer C.free(unsafe.Pointer(cFile))
	defer C.free(unsafe.Pointer(cContent))

	// Call the Rust library
	result := C.safe_write_file(cRoot, cFile, cContent)
	if result != 0 {
		return fmt.Errorf("rust fs_agent error code: %d", result)
	}
	return nil
}

func main() {
	fmt.Println("🤖 Claude Cowork (Go + Rust Hybrid)")
	fmt.Println("-----------------------------------")

	// 1. Define the Sandbox directory
	cwd, _ := os.Getwd()
	sandboxDir := filepath.Join(cwd, "sandbox")
	// Ensure the directory exists
	os.MkdirAll(sandboxDir, 0755) 

	fmt.Printf("🔒 Sandbox locked to: %s\n", sandboxDir)

	// --- AI Agent Simulation (Mock) ---
	// In a production environment, this is where the Anthropic API call would happen.
	
	task := "Create a greeting file and save it."
	fmt.Printf("\n👤 User Task: %s\n", task)
	fmt.Println("🧠 Claude is thinking...")

	// Mocking a Tool Use response from Claude:
	aiFileName := "hello_world.txt"
	aiContent := "Hello! This file was created by Rust, controlled by Go, prompted by AI."
	
	fmt.Printf("⚡ Executing Tool: WriteFile('%s')\n", aiFileName)

	// Delegate the file operation to Rust for safety
	err := RustSafeWrite(sandboxDir, aiFileName, aiContent)
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
	} else {
		fmt.Println("✨ Task Completed Successfully.")
	}

	// --- Security Test (Simulation) ---
	fmt.Println("\n🕵️ Attempting simulated hack (Path Traversal)...")
	hackFile := "../system_hack.txt"
	
	// Try to write outside the sandbox
	err = RustSafeWrite(sandboxDir, hackFile, "Hacked!")
	if err != nil {
		fmt.Println("🛡️ Security blocked the attack successfully.")
	}
}
