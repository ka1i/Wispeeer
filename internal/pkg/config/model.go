package config

import "html/template"

// options
type Options struct {
	Site       `yaml:",inline"`
	Urls       `yaml:",inline"`
	Directory  `yaml:",inline"`
	Pagination `yaml:",inline"`
	Deployment `yaml:",inline"`
}

// Site ...
type Site struct {
	Title       string        `yaml:"title" default:"Welcome to wisper"`
	Subtitle    string        `yaml:"subtitle" default:"wispeeer wisper"`
	Description template.HTML `yaml:"description" default:"Generate by Wispeeer"`
	Keywords    string        `yaml:"keywords" default:"blog"`
	Author      string        `yaml:"author" default:"void"`
	Theme       string        `yaml:"theme" default:"wisper"`
	Language    string        `yaml:"language" default:"en"`
	Timezone    string        `yaml:"timezone" default:"Local"`
}

// Urls ...
type Urls struct {
	Root      string `yaml:"root" default:"http://localhost:4180"`
	Permalink string `yaml:"permalink" default:"website"`
}

// Directory ...
type Directory struct {
	ArticleDir string `yaml:"article_dir" default:"article"`
	SourceDir  string `yaml:"source_dir" default:"source"`
	PublicDir  string `yaml:"public_dir" default:"public"`
}

// Pagination ...
type Pagination struct {
	PerPage       string `yaml:"per_page" default:"9"`
	PaginationDir string `yaml:"pagination_dir" default:"page"`
}

// Deployment ...
type Deployment struct {
	Type   string `yaml:"type" default:"git"`
	Repo   string `yaml:"repo" default:"void"`
	Branch string `yaml:"branch" default:"master"`
}
