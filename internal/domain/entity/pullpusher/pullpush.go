package pullpusher

import (
	"fmt"
	"os"
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
	if err := os.Chdir(s.RepoRoot); err != nil {
		return fmt.Errorf("failed to change directory to %v repo: %v", s.RepoName, err)
	}
	// git status
	// if up to date
	//
	//	git pull
	return nil
}

func (s Syncer) Push() error {
	if err := os.Chdir(s.RepoRoot); err != nil {
		return fmt.Errorf("failed to change directory to %v repo: %v", s.RepoName, err)
	}
	// git status
	// if not up to date
	//	    git add --all
	//      git -commit -m "today timedate"
	//      git push
	return nil
}
