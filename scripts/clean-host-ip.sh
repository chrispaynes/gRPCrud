#!/bin/bash
set -x
HOSTIP=$(ip addr show | grep wlp | grep -Po 'inet \K[\d.]+')

sed -i "s/$HOSTIP/0.0.0.0/g" ./configs/app.toml