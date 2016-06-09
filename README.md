# dynamock
This is a api to run a simple http server, which can serve up mock service responses. Responses can be JSON to simulate REST.

### Development
```bash
go run main.go -c templates/sample.json
``` 

### Configuration
On startup, config values are loaded from the config.json file.

* Services can be configured to return different responses, depending on a request parameter or request header.
* Content-type for a service response can be set for each service. If not set, content-type defaults to application/text.
* HTTP Status code can be set for each service.
* Latency (ms) can be set to simulate slow service responses. Latency can be set for a single service, or globally for all services.
* Additional headers can be defined for responses.

```json
{
    "port": "3010",
    "latency": 30,
    "contentType": "application/json",
    "mockDir": "templates/mock",
    "services": {
        "ping": {
            "contentType": "application/json",
            "responses": {
                "get": {
                    "status": 200,
                    "body": "json"
                }
            }
        },
        "api/:id/services": {
            "contentType": "application/json",
            "responses": {
                "get": {
                    "status": 200,
                    "body": "{\"id\": 1, \"name\": \"service\"}"
                },
                "post": {
                    "status": 200
                },
                "put": {
                    "status": 200
                },
                "delete": {
                    "status": 200
                }
            }
        },
        "api/:id/files": {
            "contentType": "application/json",
            "responses": {
                "get": {
                    "status": 200,
                    "bodyFile": "files.json"
                 }
            }
        },
        "dyn1": {
            "responses": {
                "get": {
                    "dynamic": {
                        "random": {
                            "status": [200, 201],
                            "body": ["ok1", "ok2", "ok3"]
                        }
                    }
                 }
            }
        },
        "dyn2": {
            "responses": {
                "get": {
                    "dynamic": {
                        "random": {
                            "status": [200, 201, 500],
                            "bodyFile": [
                                "dyn.1.json",
                                "dyn.2.json"
                            ]
                        }
                    }
                 }
            }
        },
        "headers": {
            "contentType": "application/json",
            "headers": {
                "x-requested-by": "a1334c7dh3a8",
                "authorization": "Bearer a1334c7dh3a8"
            },
            "responses": {
                "get": {
                    "status": 200,
                    "bodyFile": "files.json"
                 }
            }
        },
        "switch": {
            "contentType": "application/json",
            "headers": {
                "x-requested-by": "a1334c7dh3a8",
                "authorization": "Bearer a1334c7dh3a8"
            },
            "responses": {
                 "get": {
                    "dynamic": {
                        "switch": {
                            "id": [
                                {"if": "1", "when": {"status": 200, "body": "ok1"}},
                                {"if": "2", "when": {"status": 200, "body": "ok2"}},
                                {"if": "3", "when": {"status": 200, "body": "ok3"}}
                            ],
                            "name": [
                                {"if": "go", "when": {"status": 500, "bodyFile": "error.json"}}
                            ]
                        }
                    }
                 }
            }
        }
    }
}
```

### Usage
```bash
# compose
docker-compose up -d

# docker
docker build -t dynamock .
docker run -d -p 3010:3010 dynamock

# binary
curl -L https://github.com/helderfarias/dynamock/releases/download/v1.0/dynamock_darwin_osx.zip > dynamock.zip \
    && unzip dynamock.zip \
    && rm dynamock.zip
```


### Inspired
https://github.com/gstroup/apimocker 

