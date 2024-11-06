#!/usr/bin/env bash

FILES="api_circuits.go api_dcim.go api_ipam.go api_tenancy.go api_virtualization.go api_vpn.go api_wireless.go"

for i in $FILES; do
  sed -i '14i \       "time"' $i
done
