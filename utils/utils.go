package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Command(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func CommandOutput(args ...string) string {
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func CheckPackage(pack string) bool {
	switch OS {
	case Arch:
		if !strings.Contains(CommandOutput("pacman", "-Q"), pack) {
			return false
		} else {
			return true
		}
	case Debian:
		if !strings.Contains(CommandOutput("apt", "list", "--installed"), pack) {
			return false
		} else {
			return true
		}
	case Fedora:
		if !strings.Contains(CommandOutput("dnf", "--installed", "output"), pack) {
			return false
		} else {
			return true
		}
	case Alpine:
		if !strings.Contains(CommandOutput("apk", "list", "-i"), pack) {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}

func InstallPackage(pack string) {
	switch OS {
	case Arch:
		Command("pacman", "-S", pack)
	case Debian:
		Command("apt-get", "install", pack)
	case Fedora:
		Command("dnf", "install", pack)
	case Alpine:
		Command("apk", "add", pack)
	}
}

const (
	Arch   = "arch"
	Debian = "debian"
	Alpine = "alpine"
	Fedora = "fedora"
)

var OS string
var KernelVer string
var MinecraftFolder string

func FindDirectory(dirName string) (string, error) {
	var foundPath string
	root := "/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == dirName {
			foundPath = path
			return fmt.Errorf("directory trovata")
		}
		return nil
	})

	if err != nil && err.Error() != "directory trovata" {
		return "", err
	}

	return foundPath, nil
}

func SetAndGetInfo() {
	uid := os.Geteuid()
	if uid != 0 {
		fmt.Println("Start OpenSS as root to continue")
		os.Exit(11)
	}

	OSgetOutput := CommandOutput("uname", "-a")

	switch {
	case strings.Contains(OSgetOutput, Arch):
		OS = Arch
	case strings.Contains(OSgetOutput, Debian) || strings.Contains(OSgetOutput, "ubuntu"):
		OS = Debian
	case strings.Contains(OSgetOutput, Alpine):
		OS = Alpine
	case strings.Contains(OSgetOutput, Fedora):
		OS = Fedora
	}

	KernelVer = CommandOutput("uname", "-r")

	output, error := FindDirectory(".minecraft")
	if error != nil {
		log.Fatal(error)
	}
	MinecraftFolder = output
}
