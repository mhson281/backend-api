# Using Air for Live Reloading

Quick guide on what air is and how to use the included `.air.toml` file to live reload the application

## Step 1: Install Air

If you haven't already, install Air by running the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

```

Alternatively, you can install it via go install:

```bash
go install github.com/cosmtrek/air@latest

```

- Make sure Air is available in your PATH. You can check this by running:

```bash
air -v
```

If it outputs the version of Air, you’re ready to go.

## Step 2: Create an .air.toml Configuration File

In the root directory of your project, create an .air.toml file to configure Air. This file specifies how Air should watch and reload your application.

Here’s a sample .air.toml configuration file for this project:

```plaintext
# .air.toml
# Config file for air live reloading

# Root directory
root = "."

# Binary output name
bin = "tmp/main"

# Watch files with specific extensions
include_ext = ["go", "tpl", "tmpl", "html"]

# Exclude specific directories from being watched
exclude_dir = ["tmp", "vendor", "pkg", "configs"]

# Commands to run before reloading the application
prebuild = "go build -o tmp/main cmd/api/main.go"

# Command to run when starting the application
cmd = "./tmp/main"

# Customize logging
log_color = "true"
log_time = "true"

# Watch .env files for changes (useful for environment variable updates)
include_ext = ["go", "tpl", "tmpl", "html", "env"]

```

Explanation of the Configuration

- root: The root directory of your project.
- bin: The output path for the compiled binary. Here, it’s set to tmp/main.
- include_ext: The file extensions to watch. Here, we’re watching .go, .tpl, .tmpl, .html, and .env files.
- exclude_dir: Directories to exclude from being watched (e.g., vendor and tmp).
- prebuild: A command to build the project before starting it. This command compiles the Go program to tmp/main.
- cmd: The command to run the application. Here, it runs the compiled binary from tmp/main.
- log_color and log_time: Enable colored logging and timestamps in the Air output.

## Step 3: Run the Application with Air

To start the application with live reloading, simply run:

```bash
air
```

This will start the server and watch for file changes. Any time you make a change to your Go files (or other specified files in include_ext), Air will rebuild and restart the application automatically.

## Step 4: Testing Changes

With Air running, you can now make changes to your code, and the application will automatically reload. You don’t need to stop and restart the server each time you make an update, which speeds up the development process.

Example Workflow

1. Start Air: Run air in your terminal.
2. Make Code Changes: Edit any .go, .env, or other specified files.
3. Watch for Reload: Air will detect changes, rebuild the project, and restart the server.
4. Test Changes: Open Postman or use curl to test the endpoints as you make changes.

Notes

- The .air.toml configuration file can be customized further to suit your development needs.
- The tmp directory is used to store the binary so that it doesn’t interfere with your main project files.
- Make sure .air.toml is added to .gitignore if you don’t want this configuration file to be part of your version control.
