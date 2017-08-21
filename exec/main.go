package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "upload-pack", "..")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				log.Printf("exit status: %d", status.ExitStatus())
			}
			log.Print(err)
		}
	} else {
		log.Println("Well done")
	}
}
