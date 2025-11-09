package syncer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Galdoba/gitsync/internal/domain/values/gitstatus"
)

type Syncer struct {
	RepoName string
	RepoRoot string
}

func NewSync(name, root string) Syncer {
	return Syncer{
		RepoName: name,
		RepoRoot: root,
	}
}

func (s Syncer) Pull() error {
	gs, err := s.GetStatus()
	if err != nil {
		return fmt.Errorf("%v failed to get status: %v", s.RepoName, err)
	}
	if err := s.pull(gs); err != nil {
		return err
	}
	return nil
}

func (s Syncer) Push() error {
	gs, err := s.GetStatus()
	if err != nil {
		return fmt.Errorf("%v failed to get status: %v", s.RepoName, err)
	}
	if err := s.push(gs); err != nil {
		return err
	}
	return nil
}

func (s Syncer) GetStatus() (gitstatus.GitStatus, error) {
	if err := os.Chdir(s.RepoRoot); err != nil {
		return gitstatus.GitStatus{}, fmt.Errorf("failed to change directory to %v repo: %v", s.RepoName, err)
	}
	cmd := exec.Command("git", "status")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return gitstatus.GitStatus{}, fmt.Errorf("failed to get git output: %v", err)
	}
	statusLines := strings.Split(string(data), "\n")
	gs := gitstatus.GitStatus{}
	for _, line := range statusLines {
		if strings.Contains(line, "Your branch is up to date") {
			gs.IsUpToDate = true
		}
		if strings.Contains(line, "Changes not staged for commit:") {
			gs.HasUnstagedChanges = true
		}
		if strings.Contains(line, "Untracked files:") {
			gs.HasUntrackedFiles = true
		}
	}
	return gs, nil
}

func (s Syncer) pull(gs gitstatus.GitStatus) error {
	if err := os.Chdir(s.RepoRoot); err != nil {
		return fmt.Errorf("failed to change directory to %v repo: %v", s.RepoName, err)
	}
	if gs.HasUntrackedFiles || gs.HasUnstagedChanges {
		return fmt.Errorf("refused: local changes detected for %v", s.RepoName)
	}
	cmd := exec.Command("git", "pull")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to get git output: %v", err)
	}
	statusLines := strings.SplitSeq(string(data), "\n")
	for line := range statusLines {
		if strings.Contains("Already up to date.", line) {
			return fmt.Errorf("refused: local repo is up to date")
		}
	}
	fmt.Fprintf(os.Stderr, string(data))
	return nil
}

func (s Syncer) push(gs gitstatus.GitStatus) error {
	if err := os.Chdir(s.RepoRoot); err != nil {
		return fmt.Errorf("failed to change directory to %v repo: %v", s.RepoName, err)
	}
	if !gs.HasUntrackedFiles && !gs.HasUnstagedChanges {
		return fmt.Errorf("refused: no local changes detected")
	}
	cmdAdd := exec.Command("git", "add", "--all")
	if err := cmdAdd.Run(); err != nil {
		return fmt.Errorf("failed to run 'git add' for %v", s.RepoName)
	}
	cmdCommit := exec.Command("git", "commit", "-m", time.Now().Format(time.DateTime))
	if err := cmdCommit.Run(); err != nil {
		return fmt.Errorf("failed to run 'git commit' for %v", s.RepoName)
	}
	cmdPush := exec.Command("git", "push")
	data, err := cmdPush.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to get git output: %v", err)
	}
	// statusLines := strings.Split(string(data), "\n")
	// for _, line := range statusLines {
	// 	if strings.Contains("Already up to date.", line) {
	// 		return fmt.Errorf("refused: local repo is up to date")
	// 	}
	// }
	fmt.Fprintf(os.Stderr, string(data))
	return nil
}
