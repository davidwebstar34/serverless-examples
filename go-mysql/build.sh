#!/bin/bash

# Build
GOOS=linux go2 build -o bin/hello

# Deploy
sls deploy
