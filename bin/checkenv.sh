#!/bin/sh

if [ -z "${BAMBOO_USERNAME}" ] || [ -z "${BAMBOO_PASSWORD}" ];then
  echo "Must set BAMBOO_USERNAME and BAMBOO_PASSWORD."
  exit 1
fi