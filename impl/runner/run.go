package runner

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Run(args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, args[0], strings.Join(args[1:], " "))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("%s", string(out))
}
