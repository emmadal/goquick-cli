# 🚀 Getting Started with Quick CLI

This tutorial will guide you through the process of creating your first Quick project using the Quick CLI.

## Prerequisites

Before you begin, make sure you have:

- Go 1.18 or later installed
- Quick CLI installed (see [Installation](#installation))

## Installation

### Option 1 – Using go install

```bash
go install github.com/goquick-run/cli@latest
```

### Option 2 – Clone and build manually

```bash
git clone https://github.com/goquick-run/cli.git
cd cli
go install
```

## Creating Your First Project

Let's create a simple API project using Quick CLI.

### Step 1: Initialize a new project

```bash
quick init my-first-api --template api
```

You should see output similar to:

```
Creating project... Done!

🎉 Project created successfully! 🎉

To start using your project:

  cd my-first-api
  go mod tidy
  go run main.go
```

### Step 2: Navigate to your project directory

```bash
cd my-first-api
```

### Step 3: Install dependencies

```bash
go mod tidy
```

### Step 4: Run your project

```bash
go run cmd/server/main.go
```

Your API should now be running at `http://localhost:8080`.

## Adding a Controller

Now, let's add a user controller to our API.

### Step 1: Add a user controller

```bash
quick addc user
```

You should see output similar to:

```
Adding controller... Done!

Controller 'user' added successfully!
```

### Step 2: Implement the controller methods

Open the generated controller file at `internal/controllers/user.go` and implement the methods according to your needs.

### Step 3: Restart your application

Stop your running application (Ctrl+C) and start it again:

```bash
go run cmd/server/main.go
```

Your API now has user endpoints at `http://localhost:8080/api/users`.

## Using Custom Templates

If you have specific requirements for your projects, you can create and use custom templates.

### Step 1: Create a custom template

You can add a custom template from a Git repository or a local directory:

```bash
quick template add --source https://github.com/user/my-template.git --name my-template --description "My custom template" --category api
```

### Step 2: Use your custom template

```bash
quick init my-custom-project --template my-template
```

## Next Steps

Now that you've created your first Quick project, you can:

- Add more controllers
- Implement your business logic
- Add authentication and authorization
- Deploy your application

For more information, check out the [Quick CLI Documentation](../README.md).

Happy coding! 🎉
