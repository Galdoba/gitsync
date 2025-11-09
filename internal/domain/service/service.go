package service

import (
	"fmt"

	"github.com/Galdoba/gitsync/internal/domain/values/syncer"
	"github.com/Galdoba/gitsync/internal/infrastructure/logger"
)

type SyncService struct {
	logger  logger.Logger
	repoMap map[string]string
}

func NewService(log logger.Logger, repoMap map[string]string) (*SyncService, error) {
	if repoMap == nil {
		return nil, fmt.Errorf("repository map not provided")
	}
	ss := SyncService{
		logger:  log,
		repoMap: repoMap,
	}
	return &ss, nil
}

func (ss *SyncService) Serve(command string, args ...string) error {
	errs := []error{}
	switch command {
	case "pull":
		for repo, path := range ss.repoMap {
			s := syncer.NewSync(repo, path)
			if err := s.Pull(); err != nil {
				ss.logger.Errorf("failed to pull %v: %v", repo, err)
				errs = append(errs, err)
			}
		}
	case "push":
		for repo, path := range ss.repoMap {
			s := syncer.NewSync(repo, path)
			if err := s.Push(); err != nil {
				ss.logger.Errorf("failed to push %v: %v", repo, err)
				errs = append(errs, err)
			}
		}

	}
	if len(errs) > 0 {
		return fmt.Errorf("errors occured: %v", len(errs))
	}
	return nil
}
