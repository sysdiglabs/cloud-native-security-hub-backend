# Prometheus Catalog

![last commit](https://flat.badgen.net/github/last-commit/sysdiglabs/prometheus-hub-backend?icon=github) ![licence](https://flat.badgen.net/github/license/sysdiglabs/prometheus-hub-backend)

Prometheus Catalog is a platform for discovering and sharing rules and 
configurations for monitoring cloud native applications.

This repository contains the HTTP API and backend code that runs the
https://promcat.io site

## Usage

This code requires a recent golang version (1.13) and it uses modules to handle
the dependencies.

You also can use the docker compose file using the command:
```
docker-compose up -d
```

### Configuration

This project requires a PostgreSQL 11 server running. And you configure the app
to attack the database using the `DATABASE_URL` environment variable, which contains
the connection string for your server.

For example: `DATABASE_URL="postgres://username:password@127.0.0.1/db_name?sslmode=disable"`

### cmd/server

This is the HTTP API server and it will listen to requests on the `8080` port.

```
$ go run cmd/server/main.go
```

### Endpoints of the API

```
URL: /apps
RESPONSE: JSON with all the apps

URL: /apps/:app
RESPONSE: JSON with a specific 'app'

URL: /apps/:app/:appVersion/resources
RESPONSE: JSON with all the resources for a specific 'version' of an 'app'

URL: /resources
RESPONSE: JSON with all the resources

URL: /resources/:kind/:app/:appVersion
RESPONSE: JSON with the latest version of the 'kind' of resource for a specific 'version' of an 'app' 

URL: /resources/:kind/:app/:appVersion/:version
RESPONSE: JSON with the specific 'version' of the 'kind' of resource for a specific 'version' of an 'app' 
```

Notes: 
* The name of the app is passed in slug format (i.e. "AWS Fargate" will be passed as "aws-fargate").
* In case of error or not resources or apps found, the API returns 404 error code. 


### cmd/dbimport

You need to setup a couple of environment variables previously to import any
data in the database:

* `RESOURCES_PATH`: Path to prometheus-hub-resources/resources directory
* `APPS_PATH`: Path to prometheus-hub-resources/apps directory


These directories can be found in the [Prometheus Catalog Resources repository](https://github.com/sysdiglabs/promcat-resources).


Then with the `DATABASE_URL` set, execute:

```
$ go run cmd/dbimport/main.go
```

And voila!

## Contributing

Contributors are welcome! You will need a quick package overview to understand
some design decisions:

* `pkg/usecases`: You will find the entry points in the `pkg/usecases` directory.
  One action per file, modeled like a command.
* `pkg/resource` and `pkg/apps`: This is the domain code for resources
  and apps. You will find the repositories, entities and value objects.
* `test`: All our code is test driven, in this directory we have some fixtures
  to avoid repeating test data in the test code.
* `web`: The web is just a delivery mechanism, it is separated from the backend code
  and can be used as a library if you need to. Is responsible to JSON
  marshalling and HTTP communications.
* `db`: Contains the migration files for the database. For every change
  in the schema, you will need to create the corresponding migration file.
