#!/usr/bin/env sh

lighthouse-ci \
    --config-path=config.json \
    http://host.docker.internal:3000 \
    --best-practices=100