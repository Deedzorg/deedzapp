#!/bin/sh
set -e

log() {
  echo "$(date +"%Y-%m-%d %H:%M:%S") $1"
}

log "Starting NGINX RTMP server..."
nginx
log "NGINX started"

cleanup() {
    log "Shutting down..."
    pkill -TERM nginx || true
    pkill -TERM main || true
}
trap cleanup INT TERM EXIT

sleep 1  # give NGINX a second to initialize

log "Starting Go app..."
./main || { log "Go app crashed with status $?"; exit 1; }
