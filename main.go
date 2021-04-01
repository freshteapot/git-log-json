package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type ChangeLog struct {
	Hash string `json:"hash"`
	What string `json:"what"`
	When string `json:"when"`
	PR   string `json:"pr"`
}

func main() {
	path := os.Args[1]
	r, _ := git.PlainOpen(path)
	ref, _ := r.Head()
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})
	re := regexp.MustCompile(`(?s)\((.*)\)`)

	chanelog := make([]ChangeLog, 0)
	_ = cIter.ForEach(func(c *object.Commit) error {
		if !strings.HasPrefix(c.Message, "Change") {
			return nil
		}

		parts := strings.Split(c.Message, "\n")
		subject := parts[0]
		m := re.FindAllStringSubmatch(subject, -1)
		// I only care for the new format
		if len(m) == 0 {
			return nil
		}

		pr := m[0][1]
		pr = strings.TrimLeft(pr, "#")

		parts = parts[1:]
		message := strings.Join(parts, "\n")
		message = strings.TrimSpace(message)

		chanelog = append(chanelog, ChangeLog{
			What: message,
			When: c.Committer.When.Format(time.RFC3339),
			PR:   pr,
			Hash: c.Hash.String()[:6],
		})

		return nil
	})

	b, _ := json.Marshal(chanelog)
	fmt.Println(string(b))
}
