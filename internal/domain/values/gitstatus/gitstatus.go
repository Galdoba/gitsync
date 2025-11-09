package gitstatus

type GitStatus struct {
	IsUpToDate         bool
	HasUnstagedChanges bool
	HasUntrackedFiles  bool
}
