#!/bin/bash
docker  pull quay.io/mchirico/scrape:test
kind load docker-image quay.io/mchirico/scrape:test
