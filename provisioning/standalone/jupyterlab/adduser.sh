#!/usr/bin/env bash

useradd -ms /bin/bash crownlabs

PUID=1010
PGID=1010

groupmod -o -g "$PGID" crownlabs
usermod -o -u "$PUID" crownlabs
