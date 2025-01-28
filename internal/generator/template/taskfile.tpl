version: 3

vars:
  DEPENDENCY: ""

tasks:
  default:
    cmds:
      - task --list-all

  tools:install-toolset:
    cmd: go install github.com/kazhuravlev/toolset/cmd/toolset@latest

  tools:init:
    run: once
    desc: Init toolset config
    deps:
      - "tools:install-toolset"
    status:
      - test -f ./.toolset.json
      - test -f ./.toolset.lock.json
    cmds:
      - toolset init .

  tools:add-default-dependencies:
    desc: Adding default dependencies
    run: once
    deps:
      - "tools:init"
    cmds:
      - toolset add go github.com/golangci/golangci-lint/cmd/golangci-lint

  tools:install:
    run: once
    deps:
      - tools:add-default-dependencies
    cmds:
      - toolset sync

  tools:add-dependency:
    desc: Adding custom dependency
    deps:
      - "tools:add-default-dependencies"
    cmds:
      - toolset add {{.DEPENDENCY}}

  lint:
    cmds:
      - toolset run golangci-lint run --config=./configs/.golangci.yml ./
