services:
{{ if .hasPostgres }}
  db:
    image: postgres:17.2-alpine3.21
    restart: always
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
{{ end }}
volumes:
  postgres_data: