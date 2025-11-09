package config

type Config struct {
	TrackRepos map[string]string `toml:"track_repos"`
}

func Default() Config {
	cfg := Config{}
	cfg.TrackRepos = make(map[string]string)
	cfg.TrackRepos["gitsync"] = "~/go/src/github.com/Galdoba/gitsync"
	return cfg
}
