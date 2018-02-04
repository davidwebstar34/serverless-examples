#!/bin/bash

# Build
GOOS=linux go2 build -o bin/gorm-hello

# Deploy
sls deploy
