package projectstruct

type Tag string
type TagParams func() map[string]any

const (
	MainGoTag        Tag = "main.go"
	EditorConfigTag  Tag = "editorconfig"
	GitIgnoreTag     Tag = "gitignore"
	ReadmeTag        Tag = "readme"
	TaskfileTag      Tag = "taskfile"
	GolangCiTag      Tag = "golangci"
	DockerComposeTag Tag = "docker-compose"
	DotEnvTag        Tag = "dotenv"
	PostgresqlTag    Tag = "postgresql"
)

var DefaultTagTemplates = map[Tag]struct {
	Template        string
	ParamsGenerator TagParams
}{
	MainGoTag: {
		Template:        "main.go.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
	EditorConfigTag: {
		Template:        "editorconfig.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
	GitIgnoreTag: {
		Template:        "gitignore.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
	ReadmeTag: {
		Template:        "readme.tpl",
		ParamsGenerator: readmeTagParamsGenerator,
	},
	TaskfileTag: {
		Template:        "taskfile.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
	GolangCiTag: {
		Template:        "golangci.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
	DockerComposeTag: {
		Template:        "docker-compose.yml.tpl",
		ParamsGenerator: dockerComposeTagParamsGenerator,
	},
	DotEnvTag: {
		Template:        ".env.tpl",
		ParamsGenerator: dotEnvTagParamsGenerator,
	},
	PostgresqlTag: {
		Template:        "postgresql.yml.tpl",
		ParamsGenerator: emptyTagParamsGenerator,
	},
}

var UsedTags map[Tag]bool

func CollectTags(config []ProjectStruct) {
	if UsedTags == nil {
		UsedTags = make(map[Tag]bool, len(DefaultTagTemplates))
	}

	for _, projStruct := range config {
		if projStruct.Type == File && projStruct.Tag == "" {
			UsedTags[projStruct.Tag] = true
		}

		if projStruct.Type == Dir {
			CollectTags(projStruct.Children)
		}
	}
}

func emptyTagParamsGenerator() map[string]any {
	return make(map[string]any)
}

func dotEnvTagParamsGenerator() map[string]any {
	params := make(map[string]any)

	if _, ok := UsedTags[PostgresqlTag]; ok {
		params["hasPostgres"] = true
	} else {
		params["hasPostgres"] = false
	}

	return params
}

func dockerComposeTagParamsGenerator() map[string]any {
	params := make(map[string]any)

	if _, ok := UsedTags[PostgresqlTag]; ok {
		params["hasPostgres"] = true
	} else {
		params["hasPostgres"] = false
	}

	return params
}

func readmeTagParamsGenerator() map[string]any {
	return map[string]any{
		"projectName":        "Go Skeleton",
		"projectDescription": "Service generated by Dmitrijlin\\go-skeleton",
	}
}
