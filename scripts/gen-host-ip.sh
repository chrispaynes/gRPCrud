#!/bin/bash
sed -i "s/0.0.0.0/$(ip addr show | grep wlp | grep -Po 'inet \K[\d.]+')/g" ./configs/app.toml
