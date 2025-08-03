#!bin/bash

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)

    echo "Loaded environment variables:"
    for key in $(grep -v '^#' .env_base | sed 's/=.*//'); do
        if ! grep -q "^$key=" .env; then
            echo "Error: $key is missing in .env but exists in .env_base"
            exit 1
        fi
    done
    for key in $(grep -v '^#' .env | sed 's/=.*//'); do
        echo "$key"
    done
else
    echo ".env file not found. Using default environment variables."
    exit 1
fi

go build -o atylab-api ./main.go
