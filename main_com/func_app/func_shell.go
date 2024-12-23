package func_app

import (
	"fmt"
	config_main "head/main_com/config"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func StartShellWeb(port int) *exec.Cmd {
	originalDir, _ := os.Getwd()
	os.Chdir("shell")

	var cmd *exec.Cmd

	os.Chdir("NM1")

	htmlContent := fmt.Sprintf(`%s/main`, strconv.Itoa(port))

	args := []string{
		config_main.Name,
		config_main.Window_h,
		config_main.Window_w,
		htmlContent,
	}

	cmd = exec.Command(config_main.Core_web, args...)

	defer func() {
		os.Chdir(originalDir)
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	doneChan := make(chan error, 1)
	go func() {
		doneChan <- cmd.Wait()
	}()

	go func() {
		select {
		case sig := <-sigChan:
			fmt.Printf("error", sig)
			os.Exit(0)
		case err := <-doneChan:
			if err != nil {
				fmt.Printf("error %v\n", err)
			} else {
				fmt.Println("error")
			}
			os.Exit(0)
		}
	}()

	return cmd
}
