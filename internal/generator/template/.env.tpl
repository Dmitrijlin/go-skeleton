{{ if .hasPostgres }}
POSTGRES_USER=dev
POSTGRES_PASSWORD=dev
POSTGRES_DB={{ .projectName }}
{{ end }}