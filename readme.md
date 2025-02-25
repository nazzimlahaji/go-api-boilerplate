# Go API Boilerplate

This Go API Boilerplate provides a structured starting point for developing applications in Go Language. It includes essential configurations, guidelines, and dependencies to help you set up a Go project quickly and efficiently. By following the steps outlined in this Go API Boilerplate, you can ensure that your Go application adheres to best practices and is ready for development and deployment. Whether you are a beginner or an experienced developer, this Go API Boilerplate aims to streamline your workflow and enhance productivity.

## Features

### Authentication with Firebase

This Go API Boilerplate integrates Firebase for authentication, providing a secure and scalable way to manage user authentication. By using Firebase, you can leverage its robust features such as email/password authentication, social media logins, and more.

### Monitoring with Sentry

Sentry is included for monitoring and error tracking. It helps you identify and fix issues in your application by providing real-time error tracking and performance monitoring. With Sentry, you can gain insights into the health of your application and improve its reliability.

### ORM with GORM

GORM is used as the Object-Relational Mapping (ORM) library in this Go API Boilerplate. It simplifies database interactions by allowing you to work with Go structs instead of raw SQL queries. GORM supports various databases and provides features like auto-migrations, associations, and more.

### CLI with Cobra

Cobra is used for building the command-line interface (CLI) of your application. It provides a simple and powerful way to create CLI commands, flags, and subcommands. With Cobra, you can easily extend your application with custom commands and improve its usability.

### File Storage with Minio

Minio is included for file storage, providing a high-performance, S3-compatible object storage solution. By using Minio, you can store and manage files efficiently within your Go API Boilerplate. It supports various storage backends and offers features like data protection, encryption, and scalability.

## Guide

### Initial Setup (First-Time Users Only)

1. **Set the GOPATH:**
   Add the following lines to your `.bashrc` or `.zshrc` file to set the GOPATH:

   ```sh
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

2. **Install GolangCI-Lint:**
   GolangCI-Lint helps detect and fix style errors, convention errors, and potential vulnerabilities. Install it using the following command:

   ```sh
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
   ```

3. **Configure GolangCI-Lint in VSCode:**
   After installation, configure GolangCI-Lint in VSCode by adding the following settings to your `User Settings JSON`:
   ```json
   {
     "go.lintTool": "golangci-lint",
     "go.lintFlags": ["--enable-all"]
   }
   ```
   Enabling all lint flags will generate many warnings. It's acceptable as long as there are no errors. You can choose to fix or ignore the warnings as needed.

### Step-by-Step Guideline

1. Copy the `.env.example` file and rename it to `.env`. Fill in the required key values in the `.env` file.
2. This API uses Firebase for authentication and requires a service account key. Generate the private key via the `Firebase Console`, download the file, rename it to `serviceAccountKey.json`, and place it in the `<root>` directory of your project.
3. To install a new package for your project, ensure it is used and then vendor it using the following command:
   ```sh
   go mod vendor
   ```

## How to Start the Server

1. Ensure you are in the `<root>` directory of your project.
2. Execute the following command to start the server:
   ```sh
   go run main.go server
   ```
3. This command will initialize and run the server.

## Adding Commands with Cobra-CLI

To add a new command to your project using Cobra-CLI, follow these steps:

1. Open your terminal and navigate to the root directory of your project.
2. Use the following command to create a new command file:
   ```sh
   cobra-cli add <command_name>
   ```
   Replace `<command_name>` with the desired name of your new command.
3. This command will generate a new file in the `cmd` directory with the specified command name. The file will contain a basic template for your new command.
4. Open the newly created file in your preferred code editor. You will see a function named `init()` where you can define the behavior of your command.
5. Implement the logic for your command within the `Run` function. You can also add flags and arguments as needed.
6. Once you have defined your command, save the file and rebuild your application to include the new command.
7. You can now run your new command using the following syntax:
   ```sh
   go run main.go <command_name>
   ```
   By following these steps, you can easily extend your application's CLI with custom commands using Cobra-CLI.

# Folder Directory Explanation

This section explains the purpose of each directory in the project structure.

### cmd

The `cmd` directory contains the main entry point for your application. It includes the primary executable files and the CLI commands created using Cobra.

### pkg

The `pkg` directory holds the core packages of your application. It is organized into several subdirectories:

- `common`: Contains shared utilities and helper functions that can be used across different parts of the application.
- `config`: Manages configuration settings and environment variables for the application.
- `entity`: Defines the data models and entities used within the application.
- `middleware`: Includes middleware components for handling requests and responses, such as authentication and logging.
- `model`: Contains the business logic and data access layer, including interactions with the database using GORM.
- `server`: Manages the server setup, routing, and initialization of the application.

By organizing your project in this way, you can maintain a clean and modular structure, making it easier to manage and scale your Go application.

# Dependencies

- Admin SDK Firebase
- Sentry
- GORM (https://github.com/go-gorm/gorm)
- Cobra (https://github.com/spf13/cobra)
- Minio (https://github.com/minio/minio-go)

## Additional Dependencies (Might needed for your project)

- Validator v10 (https://github.com/go-playground/validator)
- JSON Schema (https://github.com/santhosh-tekuri/jsonschema)
- GeoJSON Orb (https://github.com/paulmach/go.geojson)
