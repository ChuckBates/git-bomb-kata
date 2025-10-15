package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	changes := dirtyFiles(ctx)
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] git-bomb: executing git reset --hard \n", timestamp)
	cmd := exec.CommandContext(ctx, "git", "reset", "--hard")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(output))
	}

	_ = triggerFileSystemRefresh()
	_ = triggerFileChangesRefresh(changes)

	return nil
}

func dirtyFiles(ctx context.Context) []string {
	cmd := exec.CommandContext(ctx, "git", "status", "--porcelain", "-z"
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	return strings.Split(string(output), "\n")
}

func triggerFileChangesRefresh(paths []string) error {
	now := time.Now()
	for _, path := range paths {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			os.Chtimes(path, now, now)
		}
	}
	return nil
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
