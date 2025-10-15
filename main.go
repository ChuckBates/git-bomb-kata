package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	interval := flag.Duration("interval", 5*time.Minute, "interval between resets")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *interval)
	defer cancel()

	fmt.Printf("git-bomb: resetting repo every %v\n", *interval)

	resetErr := reset(ctx)
	if resetErr != nil {
		fmt.Printf("git-bomb: error resetting repo %v\n", resetErr)
		os.Exit(1)
	}

	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("git-bomb: Shutting down")
			return
		case <-ticker.C:
			resetErr = reset(ctx)
			if resetErr != nil {
				fmt.Printf("git-bomb: error resetting repo %v\n", resetErr)
				os.Exit(1)
			}
		}
	}
}

func reset(ctx context.Context) error {
	timestamp := time.Now()
	fmt.Printf("git-bomb: executing git reset --hard %v\n", timestamp)
	cmd := exec.CommandContext(ctx, "git", "reset", "--hard")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(output))
	}
	return err
}
