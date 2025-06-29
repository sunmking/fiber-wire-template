# Go Web Service Template

A template for building web services in Go using Fiber, Wire, and other common libraries.

## Features

* **Fiber:** An Express inspired web framework built on Fasthttp.
* **Wire:** A code generation tool for dependency injection.
* **Zap logger:** A blazing fast, structured, leveled logger.
* **Viper config:** A complete configuration solution for Go applications.
* **JWT:** JSON Web Tokens for authentication.
* **Cron jobs:** Scheduled task execution.
* **ozzo-dbx:** A powerful data access layer for Go.
* **Redis:** In-memory data structure store, used as a cache or message broker.

## Prerequisites

* Go 1.21 or higher

## Getting Started

### Clone the repository

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

### Set up environment variables

This project uses Viper for configuration management. You can set the `APP_CONF` environment variable to specify which configuration file to use. For example:

```bash
export APP_CONF=local # Uses config/local.yaml
# or
export APP_CONF=prod # Uses config/prod.yaml
```

If `APP_CONF` is not set, it defaults to `local`.

### Install dependencies

```bash
go mod tidy
```

### Run the application

```bash
go run cmd/app/main.go
```

## Configuration

Configuration files are located in the `config/` directory.

* `config/local.yaml`: Configuration for local development.
* `config/prod.yaml`: Configuration for production.

The `APP_CONF` environment variable determines which configuration file is loaded. See the "Set up environment variables" section for more details.

## Project Structure

* `cmd/`: Main applications for the project.
    * `cmd/app/main.go`: The entry point for the web service.
* `internal/`: Private application and library code. This is where the core business logic resides.
    * `internal/bootstrap`: Application initialization (config, logger, database, etc.).
    * `internal/controller`: HTTP handlers.
    * `internal/core`: Core business logic and services.
    * `internal/middleware`: HTTP request middleware.
    * `internal/model`: Data structures and database models.
    * `internal/repository`: Data access layer.
* `pkg/`: Library code that's safe to use by external applications.
* `config/`: Configuration files.
* `route/`: API route definitions.

## API Endpoints

Specific API endpoints should be documented here. Include information such as:

* HTTP Method (GET, POST, PUT, DELETE, etc.)
* URL Path
* Request Parameters (if any)
* Request Body (if any)
* Success Response (status code, body)
* Error Responses (status codes, body)

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

Standard placeholder text:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## License

This project is licensed under the MIT License.
