
<p align="center">
  <img src="quick_logo.png" alt="Quick CLI Logo" width="400">
</p>

<p align="center">
  <strong>Create fast and reliable projects with the Quick framework</strong>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#commands">Commands</a> •
  <a href="#examples">Examples</a> •
  <a href="#project-structure">Project Structure</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

## ✨ Overview

Quick CLI is a powerful command-line tool designed to help you create and manage projects using the Quick framework. With simple and intuitive commands, you can quickly scaffold new projects, add controllers, and manage templates.

## 🔧 Installation

### Option 1 – Using go install

```bash
go install github.com/goquick-run/cli@latest
mv ~/go/bin/cli ~/go/bin/quick
```

### Option 2 – Clone and build manually

```bash
git clone https://github.com/goquick-run/cli
cd cli
go install
mv ~/go/bin/cli ~/go/bin/quick
```

### 🧠 Usage
Quick CLI provides a set of commands to help you create and manage Quick projects:

```bash
quick [command] [flags]
```

```bash
quick --help
```

```bash
$ quick
Quick CLI is a command-line tool to create
and manage projects using the Quick framework.

Create new projects, add controllers, and manage templates
with simple and intuitive commands.

Usage:
  quick [flags]
  quick [command]

Available Commands:
  addc        Adds a new controller to the project
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a new Quick project
  template    Manage templates for the Quick CLI
  ui          Launch an interactive UI to create a Quick project

Flags:
      --config string   config file (defaults to $HOME/.quick.yaml)
  -h, --help            help for quick
  -v, --version         Show the Quick CLI version

Use "quick [command] --help" for more information about a command.
```

### 🛠 Command Details
```bash
quick init
```

**Initialize a new Quick project with the basic structure.**

```bash
$ quick init --help
Initialize a new Quick project with the basic structure.
You can specify a name for the project or use the current directory.

Example:
  quick init my-project
  quick init --template api my-project

Usage:
  quick init [name-of-project] [flags]

Flags:
  -h, --help              help for init
  -t, --template string   template to use (default, api, web) (default "default")
```

**Example:**
```bash
$ quick init my-awesome-api --template api
Creating project... Done!

🎉 Project created successfully! 🎉

To start using your project:

  cd my-awesome-api
  go mod tidy
  go run cmd/server/main.go
```

For more information, check the documentation at: [https://github.com/goquick-run/docs](https://github.com/goquick-run/docs)

```bash
quick addc
```

**Add a new controller to an existing Quick project.**
```bash
$ quick addc --help
Add a new controller to the current Quick project.
Must be executed inside an existing Quick project.

Example:
  quick addc user
  quick addc product

Usage:
  quick addc [controller-name] [flags]

Flags:
  -h, --help   help for addc
```
**Example:**
```bash
$ quick addc user
Adding controller... Done!

Controller 'user' added successfully!
```
```bash
quick template
```

**Manage templates for the Quick CLI.**

```bash
$ quick template --help
Manage templates for the Quick CLI.
You can list, add, and use templates to create projects and components.

Usage:
  quick template [command]

Available Commands:
  add         Add a new template
  list        List all available templates
  use         Use a specific template

Flags:
  -h, --help   help for template
```

Use "quick template [command] --help" for more information about a command.

```bash
quick template list
```

**List all available templates.**
```bash
$ quick template list
Templates available:
----------------------
- default: Default template for Quick projects [web]
- api: Template for RESTful APIs [api]
- cli: Template for CLI applications [cli]
```

**Quick template add**
Add a new template to the Quick CLI.
```bash
$ quick template add --help
Add a new template to the Quick CLI.
You can add templates from a Git repository or a local directory.

Example:
  quick template add --source [https://github.com/user/template.git](https://github.com/user/template.git) --name my-template
  quick template add --source ./my-local-template --name my-template

Usage:
  quick template add [flags]

Flags:
      --category string      template category (api, web, cli, custom) (default "custom")
      --description string   template description
  -h, --help                 help for add
      --name string          template name
      --source string        template source (Git URL or local path)
```

**Example**
```bash
$ quick template add --source https://github.com/user/api-template.git --name custom-api --description "My custom API template" --category api
Adding template... Done!

Template 'custom-api' added successfully!
```

**Quick template use**
Use a specific template to create a new project or component.

```bash
$ quick template use --help
Use a specific template to create a new project or component.
Should be used in conjunction with other commands like init or addc.

Example:
  quick template use api-rest
  quick template use microservice

Usage:
  quick template use [template-name] [flags]

Flags:
  -h, --help   help for use
```

**Example**
```bash
$ quick template use custom-api
Checking template... Done!

Template 'custom-api' selected successfully!
Use 'quick init --template custom-api [project-name]' to create a project with this template.
```
```bash
quick ui
```

Launch an interactive UI to create a Quick project (coming soon).
```bash
$ quick ui
Launching interactive UI...

[Interactive UI will be displayed here]
```
**🌟 Examples**
Creating a new API project

```bash
$ quick init my-api --template api
Creating project... Done!

🎉 Project created successfully! 🎉

To start using your project:

  cd my-api
  go mod tidy
  go run main.go
```

For more information, check the documentation at:
  [https://github.com/goquick-run/docs](https://github.com/goquick-run/docs)

**Adding a controller to an existing project**

```bash
$ cd my-api
$ quick addc user
Adding controller... Done!

Controller 'user' added successfully!
```

**Creating a custom template and using it**
```bash
$ quick template add --source [https://github.com/user/microservice-template.git](https://github.com/user/microservice-template.git) --name microservice --description "Microservice template with gRPC" --category api
Adding template... Done!

Template 'microservice' added successfully!

$ quick init my-microservice --template microservice
Creating project... Done!

🎉 Project created successfully! 🎉

To start using your project:

  cd my-microservice
  go mod tidy
  go run main.go
```

For more information, check the documentation at:
  [https://github.com/goquick-run/docs](https://github.com/goquick-run/docs)

### 🧱 Project Structure

```bash
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── addc.go
│   ├── addc_test.go
│   ├── init.go
│   ├── init_test.go
│   ├── root.go
│   ├── root_test.go
│   ├── template
│   │   ├── add.go
│   │   ├── add_test.go
│   │   ├── list.go
│   │   ├── list_test.go
│   │   ├── template.go
│   │   ├── template_test.go
│   │   ├── use.go
│   │   └── use_test.go
│   ├── ui.go
│   └── ui_test.go
├── go.mod
├── go.sum
├── internal
│   ├── project
│   │   ├── creator.go
│   │   ├── creator_test.go
│   │   ├── validator.go
│   │   └── validator_test.go
│   ├── template
│   │   ├── manager.go
│   │   ├── manager_test.go
│   │   ├── renderer.go
│   │   └── renderer_test.go
│   └── ui
│       ├── prompt.go
│       ├── prompt_test.go
│       ├── quick.go
│       ├── quick_test.go
│       ├── spinner.go
│       └── spinner_test.go
├── main.go
└── quick_logo.png
```

### 🤝 Contributing

Contributions are welcome! Feel free to:
	•	[Open issues](https://github.com/goquick-run/cli/issues)
	•	[Submit pull requests](https://github.com/goquick-run/cli/pulls)
	•	Suggest new features

Please make sure to follow the established code structure and naming conventions.
