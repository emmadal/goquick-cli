# 🚀 `quick init` Command

The `init` command is used to create a new Quick project with the basic structure.

## Usage

```bash
quick init [name-of-project] [flags]
```

## Flags

- `-t, --template string`: Template to use (default, api, web) (default "default")
- `-h, --help`: Help for init

## Examples

### Create a new project with the default template

```bash
quick init my-project
```

### Create a new project with the API template

```bash
quick init my-api --template api
```

### Create a new project in the current directory

```bash
quick init .
```

## Project Structure

When you create a new project with the `init` command, it will generate a project with the following structure (for the default template):

```
my-project/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   └── home.go
│   ├── models/
│   │   └── user.go
│   └── routes/
│       └── routes.go
├── pkg/
│   └── utils/
│       └── utils.go
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Next Steps

After creating your project, you can:

1. Navigate to your project directory:
   ```bash
   cd my-project
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run your project:
   ```bash
   go run cmd/server/main.go
   ```

4. Add controllers to your project:
   ```bash
   quick addc user
   ```

For more information, see the [Quick CLI Documentation](../README.md).
