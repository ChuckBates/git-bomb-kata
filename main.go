package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	interval := flag.Duration("interval", 5*time.Minute, "interval between resets")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] git-bomb: resetting repo every %v\n", timestamp, *interval)

	timestamp = time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] git-bomb: performing initial reset\n", timestamp)

	resetErr := reset(ctx)
	if resetErr != nil {
		timestamp = time.Now().Format(time.RFC3339)
		fmt.Printf("[%s] git-bomb: error resetting repo %v\n", timestamp, resetErr)
		os.Exit(1)
	}

	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			timestamp = time.Now().Format(time.RFC3339)
			fmt.Printf("[%s] git-bomb: Shutting down \n", timestamp)
			return
		case <-ticker.C:
			resetErr = reset(ctx)
			if resetErr != nil {
				timestamp = time.Now().Format(time.RFC3339)
				fmt.Printf("[%s] git-bomb: error resetting repo %v\n", timestamp, resetErr)
				os.Exit(1)
			}
		}
	}
}

func reset(ctx context.Context) error {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] git-bomb: executing git reset --hard \n", timestamp)
	cmd := exec.CommandContext(ctx, "git", "reset", "--hard")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(output))
	}

	return triggerFileSystemRefresh()
}

func triggerFileSystemRefresh() error {
	absolutePath, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	file, err := os.CreateTemp(absolutePath, ".git-bomb-kata-temp")
	if err != nil {
		return err
	}

	name := file.Name()
	file.Close()

	time.Sleep(25 * time.Millisecond)
	return os.Remove(name)
}
