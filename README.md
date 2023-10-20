# golang-gin-boilerplate
Example boilerplate to create Golang Restful API with Gin Framework.

# Example
```sh
make # build
./build/golang-gin-boilerplate # run

make docker # build docker image

# go get
go env -w GOPRIVATE=github.com/harunalbayrak
go get github.com/harunalbayrak/golang-gin-boilerplate
```

# Configuration
This project uses the viper configuration tool to use yaml files as a configuration file. These configuration files are located on config directory.

# Dependencies
- zap
- gin
- viper

# API Table

| Type | Request                                    | Information                           |
|------|--------------------------------------------|---------------------------------------|
| GET  | /health                                    | Returns the status of the service     |
| POST | /projects/create                           | Creates project                       |
| GET  | /projects                                  | Retrieves all projects                |
| GET  | /projects/:project_id                      | Retrieve a project                    |
| POST | /projects/:project_id/articles/create      | Creates an article in the project     |
| GET  | /projects/:project_id/articles             | Retrieves all articles in the project |
| GET  | /projects/:project_id/articles/:article_id | Retrieves an article in the project   |