package tag

type Tag string

const (
	MainGoTag        Tag = "main.go"
	EditorConfigTag  Tag = "editorconfig"
	GitIgnoreTag     Tag = "gitignore"
	ReadmeTag        Tag = "readme"
	TaskfileTag      Tag = "taskfile"
	GolangCiTag      Tag = "golangci"
	DockerComposeTag Tag = "docker-compose"
	DotEnvTag        Tag = "dotenv"
)

var DefaultTagTemplates = map[Tag]string{
	MainGoTag:        "main.go.tpl",
	EditorConfigTag:  "editorconfig.tpl",
	GitIgnoreTag:     "gitignore.tpl",
	ReadmeTag:        "readme.tpl",
	TaskfileTag:      "taskfile.tpl",
	GolangCiTag:      "golangci.tpl",
	DockerComposeTag: "docker-compose.yml.tpl",
	DotEnvTag:        ".env",
}
