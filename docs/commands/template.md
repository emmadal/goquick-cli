# 📝 `quick template` Command

The `template` command is used to manage templates for the Quick CLI.

## Usage

```bash
quick template [command]
```

## Subcommands

- `list`: List all available templates
- `add`: Add a new template
- `use`: Use a specific template

## template list

The `list` subcommand lists all available templates.

```bash
quick template list
```

Example output:

```
Templates available:
----------------------
- default: Default template for Quick projects [web]
- api: Template for RESTful APIs [api]
- cli: Template for CLI applications [cli]
```

## template add

The `add` subcommand adds a new template to the Quick CLI.

```bash
quick template add [flags]
```

### Flags

- `--source string`: Template source (Git URL or local path)
- `--name string`: Template name
- `--description string`: Template description
- `--category string`: Template category (api, web, cli, custom) (default "custom")
- `-h, --help`: Help for add

### Examples

#### Add a template from a Git repository

```bash
quick template add --source https://github.com/user/template.git --name my-template
```

#### Add a template from a local directory

```bash
quick template add --source ./my-local-template --name my-template --description "My custom template" --category api
```

## template use

The `use` subcommand selects a specific template for use with other commands.

```bash
quick template use [template-name] [flags]
```

### Flags

- `-h, --help`: Help for use

### Example

```bash
quick template use api-rest
```

## Template Structure

A template is a directory with a specific structure that Quick CLI can use to generate projects and components. A typical template structure looks like this:

```
my-template/
├── template.yaml        # Template metadata
├── project/             # Project template files
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── controllers/
│   │   │   └── home.go
│   │   ├── models/
│   │   │   └── user.go
│   │   └── routes/
│   │       └── routes.go
│   ├── pkg/
│   │   └── utils/
│   │       └── utils.go
│   ├── .gitignore
│   ├── go.mod
│   └── README.md
└── controller/          # Controller template files
    └── controller.go
```

The `template.yaml` file contains metadata about the template:

```yaml
name: my-template
description: My custom template
category: api
version: 1.0.0
author: Your Name
```

## Creating Custom Templates

You can create custom templates for your projects and components. To create a custom template:

1. Create a directory with your template files
2. Add a `template.yaml` file with metadata about your template
3. Add your template to Quick CLI using the `template add` command

Templates can use variables that will be replaced with actual values when the template is used. Variables are enclosed in double curly braces, like `{{.ProjectName}}`.

Available variables:

- `{{.ProjectName}}`: The name of the project
- `{{.PackageName}}`: The package name (usually the same as the project name)
- `{{.ControllerName}}`: The name of the controller (when using `addc`)

For more information, see the [Quick CLI Documentation](../README.md).
