package config

type GitRepo struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

func (gr *GitRepo) GetUrl() string {
	return gr.URL
}

func (gr *GitRepo) GetPath() string {
	return gr.Path
}

func (gr *GitRepo) CanFetch() bool {
	return gr.URL != ""
}

type GitRepoConfiger interface {
	GetUrl() string
	GetPath() string
	CanFetch() bool
}
