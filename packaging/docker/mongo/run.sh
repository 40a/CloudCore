#!/bin/bash
set -m

mongod &

mongo admin --eval "help" > /dev/null 2>&1
RET=$?

while [[ RET -ne 0 ]]; do
  echo "Waiting for MongoDB to start..."
  mongo admin --eval "help" > /dev/null 2>&1
  RET=$?
  sleep 1
done

echo "Setting up users..."
# create root user
mongo admin --eval "db.createUser({user: '$MONGO_ROOT_USER', pwd: '$MONGO_ROOT_PASSWORD', roles:[{ role: 'root', db: 'admin' }]});"
# create app user/database
mongo $MONGO_APP_DATABASE --eval "db.createUser({ user: '$MONGO_APP_USER', pwd: '$MONGO_APP_PASSWORD', roles: [{ role: 'readWrite', db: '$MONGO_APP_DATABASE' }, { role: 'read', db: 'local' }]});"
echo "Shutting down"
mongo admin --eval "db.shutdownServer();"

sleep 3

mongod --auth