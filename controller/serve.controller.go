package controller

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/VinukaThejana/qr-share/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/eiannone/keyboard"
	"github.com/mdp/qrterminal/v3"
)

func listner() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	return l, func() {
		l.Close()
	}
}

func getPort(l net.Listener) int {
	return l.Addr().(*net.TCPAddr).Port
}

func quit() {
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			utils.Text{}.Error(err.Error())
			os.Exit(1)
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}
	}
}

// Serve is a function that is used to serve the file or the directory
func Serve(path string) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		dir, filename, err := utils.GetDirAndFile(path)
		if err != nil {
			utils.Text{}.Error(err.Error())
			return
		}

		ip, err := utils.GetOutBoundIP()
		if err != nil {
			utils.Text{}.Error(err.Error())
			os.Exit(1)
		}

		l, close := listner()
		defer close()

		fs := http.FileServer(http.Dir(dir))
		http.Handle("/", fs)

		var url string
		if filename != "" {
			url = fmt.Sprintf("http://%s:%d/%s", ip.String(), getPort(l), filename)
		} else {
			url = fmt.Sprintf("http://%s:%d/", ip.String(), getPort(l))
		}

		fmt.Println(utils.Text{}.H(utils.Style{
			Color: lipgloss.Color("#ffffff"),
			Align: lipgloss.Center,
			Bold:  true,
			Padding: utils.P{
				Top:    1,
				Bottom: 1,
				Left:   1,
			},
		}, fmt.Sprintf("Visit : %s", url)))
		config := qrterminal.Config{
			Level:     qrterminal.M,
			Writer:    os.Stdout,
			BlackChar: qrterminal.WHITE,
			WhiteChar: qrterminal.BLACK,
			QuietZone: 1,
		}
		qrterminal.GenerateWithConfig(url, config)

		fmt.Println(utils.Text{}.P(utils.Style{
			Color: lipgloss.Color("#30363D"),
			Padding: utils.P{
				Left:   1,
				Top:    1,
				Bottom: 1,
			},
			Align: lipgloss.Left,
			Bold:  true,
		}, "Press ESC to quit"),
		)

		if err := http.Serve(l, nil); err != http.ErrServerClosed {
			utils.Text{}.Error(err.Error())
		} else {
			return
		}
	}()

	if err := keyboard.Open(); err != nil {
		utils.Text{}.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	quit()

	os.Exit(0)
	wg.Wait()
}
