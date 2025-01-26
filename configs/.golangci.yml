# Options for analysis running.
run:
  # Timeout for analysis, e.g. 30s, 5m.
  # If the value is lower or equal to 0, the timeout is disabled.
  # Default: 1m
  timeout: 5m
  # Exit code when at least one issue was found.
  # Default: 1
  issues-exit-code: 10
  # Include test files or not.
  # Default: true
  tests: true
  # List of build tags, all linters use it.
  # Default: []
  build-tags: []
  # If set, we pass it to "go list -mod={option}". From "go help modules":
  # If invoked with -mod=readonly, the go command is disallowed from the implicit
  # automatic updating of go.mod described above. Instead, it fails when any changes
  # to go.mod are needed. This setting is most useful to check that go.mod does
  # not need updates, such as in a continuous integration and testing system.
  # If invoked with -mod=vendor, the go command assumes that the vendor
  # directory holds the correct copies of dependencies and ignores
  # the dependency descriptions in go.mod.
  #
  # Allowed values: readonly|vendor|mod
  # Default: ""
  modules-download-mode: mod
  # Allow multiple parallel golangci-lint instances running.
  # If false, golangci-lint acquires file lock on start.
  # Default: false
  allow-parallel-runners: false
  # Allow multiple golangci-lint instances running, but serialize them around a lock.
  # If false, golangci-lint exits with an error if it fails to acquire file lock on start.
  # Default: false
  allow-serial-runners: false
  # Define the Go version limit.
  # Mainly related to generics support since go1.18.
  # Default: use Go version from the go.mod file, fallback on the env var `GOVERSION`, fallback on 1.17
  #  go: '1.19'
  # Number of operating system threads (`GOMAXPROCS`) that can execute golangci-lint simultaneously.
  # If it is explicitly set to 0 (i.e. not the default) then golangci-lint will automatically set the value to match Linux container CPU quota.
  # Default: the number of logical CPUs in the machine
  concurrency: 4

# output configuration options
output:
  # The formats used to render issues.
  # Formats:
  # - `colored-line-number`
  # - `line-number`
  # - `json`
  # - `colored-tab`
  # - `tab`
  # - `html`
  # - `checkstyle`
  # - `code-climate`
  # - `junit-xml`
  # - `junit-xml-extended`
  # - `github-actions`
  # - `teamcity`
  # - `sarif`
  # Output path can be either `stdout`, `stderr` or path to the file to write to.
  #
  # For the CLI flag (`--out-format`), multiple formats can be specified by separating them by comma.
  # The output can be specified for each of them by separating format name and path by colon symbol.
  # Example: "--out-format=checkstyle:report.xml,json:stdout,colored-line-number"
  # The CLI flag (`--out-format`) override the configuration file.
  #
  # Default:
  #   formats:
  #     - format: colored-line-number
  #       path: stdout
  formats:
    - format: colored-line-number
  # Print lines of code with issue.
  # Default: true
  print-issued-lines: true
  # Print linter name in the end of issue text.
  # Default: true
  print-linter-name: true
  # Add a prefix to the output file references.
  # Default: ""
  path-prefix: ""
  # Sort results by the order defined in `sort-order`.
  # Default: false
  sort-results: true
  # Order to use when sorting results.
  # Require `sort-results` to `true`.
  # Possible values: `file`, `linter`, and `severity`.
  #
  # If the severity values are inside the following list, they are ordered in this order:
  #   1. error
  #   2. warning
  #   3. high
  #   4. medium
  #   5. low
  # Either they are sorted alphabetically.
  #
  # Default: ["file"]
  sort-order:
    - linter
    - severity
    - file # filepath, line, and column.
  # Show statistics per linter.
  # Default: false
  show-stats: true

severity:
  # Set the default severity for issues.
  #
  # If severity rules are defined and the issues do not match or no severity is provided to the rule
  # this will be the default severity applied.
  # Severities should match the supported severity names of the selected out format.
  # - Code climate: https://docs.codeclimate.com/docs/issues#issue-severity
  # - Checkstyle: https://checkstyle.sourceforge.io/property_types.html#SeverityLevel
  # - GitHub: https://help.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
  # - TeamCity: https://www.jetbrains.com/help/teamcity/service-messages.html#Inspection+Instance
  #
  # `@linter` can be used as severity value to keep the severity from linters (e.g. revive, gosec, ...)
  #
  # Default: ""
  default-severity: ""
  # If set to true `severity-rules` regular expressions become case-sensitive.
  # Default: false
  case-sensitive: false
  # When a list of severity rules are provided, severity information will be added to lint issues.
  # Severity rules have the same filtering capability as exclude rules
  # except you are allowed to specify one matcher per severity rule.
  #
  # `@linter` can be used as severity value to keep the severity from linters (e.g. revive, gosec, ...)
  #
  # Only affects out formats that support setting severity information.
  #
  # Default: []
  rules: []

linters:
  # Disable all linters.
  # Default: false
  disable-all: true
  # Enable specific linter
  # https://golangci-lint.run/usage/linters/#enabled-by-default
  enable:
    - asasalint
    - asciicheck
    - bidichk
    - bodyclose
    - canonicalheader
    - containedctx
    - contextcheck
    - copyloopvar
    - cyclop
    - decorder
    - depguard
    - dogsled
    - dupl
    - dupword
    - durationcheck
    - err113
    - errcheck
    - errchkjson
    - errname
    - errorlint
    - execinquery
    - exhaustive
    - exhaustruct
    - exportloopref
    - exptostd
    - fatcontext
    - forbidigo
    - forcetypeassert
    - funlen
    - gci
    - ginkgolinter
    - gocheckcompilerdirectives
    - gochecknoglobals
    - gochecknoinits
    - gochecksumtype
    - gocognit
    - goconst
    - gocritic
    - gocyclo
    - godot
    - godox
    - gofmt
    - gofumpt
    - goheader
    - goimports
    - gomoddirectives
    - gomodguard
    - goprintffuncname
    - gosec
    - gosimple
    - gosmopolitan
    - govet
    - grouper
    - iface
    - importas
    - inamedparam
    - ineffassign
    - interfacebloat
    - intrange
    - ireturn
    - lll
    - loggercheck
    - maintidx
    - makezero
    - mirror
    - misspell
    - mnd
    - musttag
    - nakedret
    - nestif
    - nilerr
    - nilnesserr
    - nilnil
    - nlreturn
    - noctx
    - nolintlint
    - nonamedreturns
    - nosprintfhostport
    - paralleltest
    - perfsprint
    - prealloc
    - predeclared
    - promlinter
    - protogetter
    - reassign
    - recvcheck
    - revive
    - rowserrcheck
    - sloglint
    - spancheck
    - sqlclosecheck
    - staticcheck
    - stylecheck
    - tagalign
    - tagliatelle
    - tenv
    - testableexamples
    - testifylint
    - testpackage
    - thelper
    - tparallel
    - unconvert
    - unparam
    - unused
    - usestdlibvars
    - usetesting
    - varnamelen
    - wastedassign
    - whitespace
    - wrapcheck
    - wsl
    - zerologlint
  # Enable all available linters.
  # Default: false
  enable-all: false
  # Disable specific linter
  # https://golangci-lint.run/usage/linters/#disabled-by-default
  disable: []
  # Enable presets.
  # https://golangci-lint.run/usage/linters
  # Default: []
  presets:
    - bugs
    - comment
    - complexity
    - error
    - format
    - import
    - metalinter
    - module
    - performance
    - sql
    - style
    - test
    - unused
  # Enable only fast linters from enabled linters set (first run won't be fast)
  # Default: false
  fast: false
linters-settings:
  asasalint:
    # To specify a set of function names to exclude.
    # The values are merged with the builtin exclusions.
    # The builtin exclusions can be disabled by setting `use-builtin-exclusions` to `false`.
    # Default: ["^(fmt|log|logger|t|)\.(Print|Fprint|Sprint|Fatal|Panic|Error|Warn|Warning|Info|Debug|Log)(|f|ln)$"]
    # exclude:
    #   - Append
    #   - \.Wrapf
    # To enable/disable the asasalint builtin exclusions of function names.
    # See the default value of `exclude` to get the builtin exclusions.
    # Default: true
    use-builtin-exclusions: true
    # Ignore *_test.go files.
    # Default: false
    ignore-test: false
  decorder:
    # Required order of `type`, `const`, `var` and `func` declarations inside a file.
    # Default: types before constants before variables before functions.
    dec-order:
      - const
      - type
      - var
      - func
    # If true, underscore vars (vars with "_" as the name) will be ignored at all checks
    # Default: false (underscore vars are not ignored)
    ignore-underscore-vars: false
    # If true, order of declarations is not checked at all.
    # Default: true (disabled)
    disable-dec-order-check: false
    # If true, `init` func can be anywhere in file (does not have to be declared before all other functions).
    # Default: true (disabled)
    disable-init-func-first-check: false
    # If true, multiple global `type`, `const` and `var` declarations are allowed.
    # Default: true (disabled)
    disable-dec-num-check: true
    # If true, type declarations will be ignored for dec num check
    # Default: false (type statements are not ignored)
    disable-type-dec-num-check: false
    # If true, const declarations will be ignored for dec num check
    # Default: false (const statements are not ignored)
    disable-const-dec-num-check: false
    # If true, var declarations will be ignored for dec num check
    # Default: false (var statements are not ignored)
    disable-var-dec-num-check: false
  depguard:
    # Rules to apply.
    #
    # Variables:
    # - File Variables
    #   Use an exclamation mark `!` to negate a variable.
    #   Example: `!$test` matches any file that is not a go test file.
    #
    #   `$all` - matches all go files
    #   `$test` - matches all go test files
    #
    # - Package Variables
    #
    #   `$gostd` - matches all of go's standard library (Pulled from `GOROOT`)
    #
    # Default (applies if no custom rules are defined): Only allow $gostd in all files.
    rules:
      # Name of a rule.
      main:
        # Defines package matching behavior. Available modes:
        # - `original`: allowed if it doesn't match the deny list and either matches the allow list or the allow list is empty.
        # - `strict`: allowed only if it matches the allow list and either doesn't match the deny list or the allow rule is more specific (longer) than the deny rule.
        # - `lax`: allowed if it doesn't match the deny list or the allow rule is more specific (longer) than the deny rule.
        # Default: "original"
        list-mode: original
        # List of file globs that will match this list of settings to compare against.
        # Default: $all
        files:
          - "$all"
        # List of allowed packages.
        # Entries can be a variable (starting with $), a string prefix, or an exact match (if ending with $).
        # Default: []
        allow:
          - $gostd
        # List of packages that are not allowed.
        # Entries can be a variable (starting with $), a string prefix, or an exact match (if ending with $).
        # Default: []
        deny:
          - pkg: "github.com/pkg/errors"
            desc: Should be replaced by standard lib errors package
issues:
  # Excluding configuration per-path, per-linter, per-text and per-source
  exclude-rules:
    # Exclude some linters from running on tests files.
    - path: _test\.go
      linters:
        - gocyclo
        - errcheck
        - dupl
        - gosec
    # Run some linter only for test files by excluding its issues for everything else.
    - path-except: _test\.go
      linters:
        - forbidigo
    # Exclude some `staticcheck` messages.
    - linters:
        - staticcheck
      text: "SA9003:"
    # Exclude `lll` issues for long lines with `go:generate`.
    - linters:
        - lll
      source: "^//go:generate "
  # Maximum issues count per one linter.
  # Set to 0 to disable.
  # Default: 50
  max-issues-per-linter: 0
  # Maximum count of issues with the same text.
  # Set to 0 to disable.
  # Default: 3
  max-same-issues: 0
