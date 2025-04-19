# 🧩 Creating and Using Custom Templates

This tutorial will guide you through the process of creating and using custom templates with Quick CLI.

## What are Templates?

Templates are the building blocks of Quick projects. They define the structure and functionality of your projects and components. Quick CLI comes with several built-in templates, but you can also create your own custom templates to suit your specific needs.

## Creating a Custom Template

### Step 1: Create a template directory structure

First, create a directory for your template:

```bash
mkdir -p my-custom-template/project
mkdir -p my-custom-template/controller
```

### Step 2: Add template metadata

Create a `template.yaml` file in the root of your template directory:

```yaml
name: my-custom-template
description: My custom template for microservices
category: api
version: 1.0.0
author: Your Name
```

### Step 3: Add project template files

Add your project template files to the `project` directory. These files will be used when creating a new project with your template.

For example, create a basic project structure:

```bash
mkdir -p my-custom-template/project/cmd/server
mkdir -p my-custom-template/project/internal/config
mkdir -p my-custom-template/project/internal/controllers
mkdir -p my-custom-template/project/internal/models
mkdir -p my-custom-template/project/internal/routes
mkdir -p my-custom-template/project/pkg/utils
```

### Step 4: Add template files

Create template files with placeholders for variables. For example, create a `main.go` file in the `project/cmd/server` directory:

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"{{.PackageName}}/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	fmt.Println("Starting {{.ProjectName}} server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

### Step 5: Add controller template files

Add your controller template files to the `controller` directory. These files will be used when adding a new controller with your template.

For example, create a `controller.go` file in the `controller` directory:

```go
package controllers

import (
	"net/http"
)

// {{.ControllerName | title}}Controller handles {{.ControllerName}}-related requests
type {{.ControllerName | title}}Controller struct {
}

// New{{.ControllerName | title}}Controller creates a new {{.ControllerName | title}}Controller
func New{{.ControllerName | title}}Controller() *{{.ControllerName | title}}Controller {
	return &{{.ControllerName | title}}Controller{}
}

// GetAll handles GET requests to retrieve all {{.ControllerName}}s
func (c *{{.ControllerName | title}}Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

// GetByID handles GET requests to retrieve a {{.ControllerName}} by ID
func (c *{{.ControllerName | title}}Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

// Create handles POST requests to create a new {{.ControllerName}}
func (c *{{.ControllerName | title}}Controller) Create(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

// Update handles PUT requests to update a {{.ControllerName}}
func (c *{{.ControllerName | title}}Controller) Update(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

// Delete handles DELETE requests to delete a {{.ControllerName}}
func (c *{{.ControllerName | title}}Controller) Delete(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}
```

## Adding Your Custom Template to Quick CLI

Once you've created your custom template, you can add it to Quick CLI using the `template add` command:

```bash
quick template add --source ./my-custom-template --name my-custom-template --description "My custom template for microservices" --category api
```

If your template is in a Git repository, you can add it directly from there:

```bash
quick template add --source https://github.com/user/my-custom-template.git --name my-custom-template --description "My custom template for microservices" --category api
```

## Using Your Custom Template

### Creating a new project with your custom template

```bash
quick init my-microservice --template my-custom-template
```

### Adding a controller with your custom template

```bash
cd my-microservice
quick addc user
```

## Template Variables

Quick CLI supports several variables that you can use in your templates:

- `{{.ProjectName}}`: The name of the project
- `{{.PackageName}}`: The package name (usually the same as the project name)
- `{{.ControllerName}}`: The name of the controller (when using `addc`)

You can also use template functions like:

- `{{.ControllerName | title}}`: Capitalize the first letter of the controller name
- `{{.ControllerName | lower}}`: Convert the controller name to lowercase
- `{{.ControllerName | upper}}`: Convert the controller name to uppercase

## Advanced Template Features

### Conditional Statements

You can use conditional statements in your templates:

```go
package main

import (
	"fmt"
	{{if .UseHTTP}}"net/http"{{end}}
)

func main() {
	{{if .UseHTTP}}
	http.ListenAndServe(":8080", nil)
	{{else}}
	fmt.Println("HTTP server disabled")
	{{end}}
}
```

### Loops

You can use loops in your templates:

```go
package main

import (
	"fmt"
)

func main() {
	{{range .Features}}
	fmt.Println("Feature: {{.}}")
	{{end}}
}
```

## Best Practices

When creating custom templates, follow these best practices:

1. Use meaningful variable names
2. Document your template with comments
3. Test your template with different project names and configurations
4. Keep your template files organized
5. Use conditional statements to handle different scenarios
6. Provide sensible defaults

## Conclusion

Custom templates are a powerful way to extend Quick CLI and tailor it to your specific needs. By creating your own templates, you can standardize project structures, enforce best practices, and save time when creating new projects.

For more information, check out the [Quick CLI Documentation](../README.md).
