#!/bin/bash
set -e
#set -x

MYAPP=network-exporter
MYAPP_USER=network_exporter
MYAPP_GROUP=network_exporter
MYAPP_SERVICE=${MYAPP}
MYAPP_BIN=/usr/bin/${MYAPP}
MYAPP_DESCRIPTION="Prometheus Exporter for Networking"
MYAPP_CONF_DIR="/etc/${MYAPP}"
MYAPP_SYSCONF="/etc/sysconfig/${MYAPP_SERVICE}"

systemctl stop ${MYAPP_SERVICE}
systemctl disable ${MYAPP_SERVICE}
if systemctl is-active --quiet ${MYAPP_SERVICE}; then
  printf "FAIL: ${MYAPP_SERVICE} service is running\n"
  exit 1
else
  printf "INFO: ${MYAPP_SERVICE} service is not running\n"
fi

rm -rf /usr/bin/network-exporter

if [ -e ${MYAPP_CONF_DIR}/vault.yml ]; then
  mv ${MYAPP_CONF_DIR}/vault.yml ${MYAPP_CONF_DIR}/vault.yml.`date +"%Y%m%d.%H%M%S"`
fi

if [ -e ${MYAPP_CONF_DIR}/vault.key ]; then
  mv ${MYAPP_CONF_DIR}/vault.key ${MYAPP_CONF_DIR}/vault.key.`date +"%Y%m%d.%H%M%S"`
fi

if [ -e ${MYAPP_CONF_DIR}/hosts ]; then
  mv ${MYAPP_CONF_DIR}/hosts ${MYAPP_CONF_DIR}/hosts.`date +"%Y%m%d.%H%M%S"`
fi

if [ -e ${MYAPP_SYSCONF} ]; then
  mv ${MYAPP_SYSCONF} ${MYAPP_SYSCONF}.`date +"%Y%m%d.%H%M%S"`
fi
