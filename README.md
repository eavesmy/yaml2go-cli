# yaml2go-cli

**Forked from `github.com/Icemap/yaml2go-cli`**

# Change
1. Yaml parser `github.com/PrasadG193/yaml2go` -> `github.com/fengzxu/yaml2go`.
2. Default package name is the base of output path.

## Install

```bash
go install github.com/eavesmy/yaml2go-cli@latest
```

## Show Help

```bash
./yaml2go-cli -h                                    
yaml2go-cli is a cli-tool for yaml to go struct

Usage:
  yaml2go-cli [flags]

Flags:
  -h, --help             help for yaml2go-cli
  -i, --input string     input yaml file path
  -o, --output string    output go file path
  -p, --package string   package name (default "main")
  -s, --struct string    struct name (default "Default")
```

## Example
```
./yaml2go-cli -i example/conf.yaml -o example/conf.go
```

# Thx [Icemap](https://github.com/Icemap/yaml2go-cli) and [fengzxu](github.com/fengzxu/yaml2go)
