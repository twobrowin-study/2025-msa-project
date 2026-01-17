#!/bin/bash

echo "Initializing migrations"
go run ./migrate db init

echo "Performing migrations"
go run ./migrate db migrate