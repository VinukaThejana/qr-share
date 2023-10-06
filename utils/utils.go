// Package utils contains utility functions
package utils

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// GetDirAndFile is a function that is used to get the directory
// and the file from the given file path
func GetDirAndFile(filePath string) (dir string, filename string, err error) {
	if !filepath.IsAbs(filePath) {
		filePath, err = filepath.Abs(filePath)
		if err != nil {
			return "", "", err
		}
	}

	var seperator string
	if runtime.GOOS == "windows" {
		seperator = `\`
	} else {
		seperator = "/"
	}

	path := filePath
	if runtime.GOOS != "windows" && strings.Contains(filePath, "~/") {
		currentUser, err := user.Current()
		data := strings.Split(filePath, "~/")

		if err != nil {
			return "", "", err
		}
		if runtime.GOOS == "darwin" {
			path = fmt.Sprintf("/Users/%s/%s", currentUser.Username, data[1])
		} else {
			path = fmt.Sprintf("/home/%s/%s", currentUser.Username, data[1])
		}
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", "", err
	}

	if info.IsDir() {
		return path, "", nil
	}

	data := strings.Split(path, seperator)
	if len(data) == 0 {
		return "", "", fmt.Errorf("unexpected error")
	}
	filename = strings.ReplaceAll(data[len(data)-1], " ", "%20")

	for i, value := range data {
		if i == len(data)-1 || i == 0 {
			continue
		}

		dir = dir + seperator + value
	}
	return dir, filename, nil
}

// GetOutBoundIP is a function that is used to get the IP address of the device that can
// connect to the internet
func GetOutBoundIP() (ip net.IP, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP, nil
}
