package main

import (
	"fmt"
	"os"

	"github.com/nick-vanduijn/replicate/pkg/replicate"
)

func usage() {
	fmt.Println("Usage: replicate <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  check       Check whether Homebrew is installed")
	fmt.Println("  install     Run the bundled install script (macOS only)")
	fmt.Println("  --version   Print version")
	fmt.Println("  help        Show this help")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	cmd := os.Args[1]
	switch cmd {
	case "check":
		code, err := replicate.CheckHomebrew()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(code)
	case "install":
		code, err := replicate.RunInstallScript()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(code)
	case "help":
		usage()
	case "--version", "-v":
		fmt.Println(replicate.Version)
	default:
		usage()
		os.Exit(1)
	}
}
