package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	absPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		fmt.Println("Error resolving symlinks:", err)
		return
	}

	fmt.Println(" Absolute path to the executable:", absPath)
	time.Sleep(10 * time.Second)
}
