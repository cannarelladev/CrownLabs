#!/usr/bin/env bash

useradd -s /bin/bash crownlabs

PUID=1010
PGID=1010

groupmod -o -g "$PGID" crownlabs
usermod -o -u "$PUID" crownlabs

usermod -d /config crownlabs

mkdir /app
mkdir /config
mkdir /defaults

chown crownlabs:crownlabs /app
chown crownlabs:crownlabs /config
chown crownlabs:crownlabs /defaults
