#!/bin/bash

./dockerbuild.sh
docker build -t arschles/badhttp .
docker push arschles/badhttp
