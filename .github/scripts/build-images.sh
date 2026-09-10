#!/usr/bin/env bash
set -euo pipefail

make docker-build && make docker-save
make signimage-build && make signimage-save
make webhookimage-build && make webhookimage-save
make workerimage-build && make workerimage-save
