# Lambda Golang


This project contains a Go application designed to be deployed as an AWS Lambda function. The function exposes an HTTP endpoint.

## Getting Started

### Prerequisites

* Docker

### Building and Running

To build and run the project locally, you can use the provided Docker Compose file.

```bash
mkdir -p ~/.aws-lambda-rie && \
    curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie && \
    chmod +x ~/.aws-lambda-rie/aws-lambda-rie```

```bash
docker-compose up --build
```

This will build the Docker image and start the Lambda function locally, making it accessible on port 9000.

## Usage

You can invoke the Lambda function using a `curl` command like the one below:

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" \
  -H "Content-Type: application/json" \
  -d \
'{
    "version": "2.0",
    "rawPath": "/test",
    "requestContext": {
      "http": {
        "method": "POST",
        "path": "/test"
      }
    },
    "body": "{\"name\":\"Edu\"}"
  }'
```

### Expected Output

The service will respond with the following output:

```
testPathEdu
```
