#!/bin/bash
set -e

echo "Running migrations..."
DATABASE_URL=${DATABASE_URL} npx tsx migrate.ts

echo "Starting production server..."
node server.js