#!/usr/bin/env bash

for i in api_*.go; do
  sed -i '14i \       "time"' $i
done
