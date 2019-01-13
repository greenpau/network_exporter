#!/bin/bash
set -e
set -x

MYAPP=network-exporter
MYAPP_USER=network_exporter
MYAPP_GROUP=network_exporter
MYAPP_SERVICE=${MYAPP}
MYAPP_BIN=/usr/bin/${MYAPP}
MYAPP_DESCRIPTION="Prometheus Exporter for Networking"
MYAPP_CONF_DIR="/etc/${MYAPP}"
MYAPP_SYSCONF="/etc/sysconfig/${MYAPP_SERVICE}"

if [ -f "./${MYAPP}" ]; then
  rm -rf $MYAPP_BIN
  cp ./${MYAPP} ${MYAPP_BIN}
fi

if getent group ${MYAPP_GROUP}  >/dev/null; then
  printf "INFO: ${MYAPP_GROUP} group already exists\n"
else
  printf "INFO: ${MYAPP_GROUP} group does not exist, creating ...\n"
  groupadd --system ${MYAPP_GROUP}
fi

if getent passwd ${MYAPP_USER} >/dev/null; then
  printf "INFO: ${MYAPP_USER} user already exists\n"
else
  printf "INFO: ${MYAPP_USER} group does not exist, creating ...\n"
  useradd --system -d /var/lib/${MYAPP} -s /bin/bash -g ${MYAPP_GROUP} ${MYAPP_USER}
fi

mkdir -p /etc/${MYAPP}
touch ${MYAPP_CONF_DIR}/hosts
touch ${MYAPP_CONF_DIR}/vault.yml
touch ${MYAPP_CONF_DIR}/vault.key


cat << EOF > ${MYAPP_SYSCONF}
OPTIONS="-log.level debug \
-api.poll-interval 15 \
-api.inventory ${MYAPP_CONF_DIR}/hosts \
-api.vault ${MYAPP_CONF_DIR}/vault.yml \
-api.vault.key ${MYAPP_CONF_DIR}/vault.key"
EOF

chown -R ${MYAPP_USER}:${MYAPP_GROUP} /etc/${MYAPP}
mkdir -p /var/lib/${MYAPP}
chown -R ${MYAPP_USER}:${MYAPP_GROUP} /var/lib/${MYAPP}

cat << EOF > /usr/lib/systemd/system/${MYAPP_SERVICE}.service
[Unit]
Description=$MYAPP_DESCRIPTION
After=network.target

[Service]
User=${MYAPP_USER}
Group=${MYAPP_GROUP}
EnvironmentFile=-${MYAPP_SYSCONF}
ExecStart=${MYAPP_BIN} \$OPTIONS
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl is-active --quiet ${MYAPP_SERVICE} && systemctl stop ${MYAPP_SERVICE}
systemctl enable ${MYAPP_SERVICE}
systemctl start ${MYAPP_SERVICE}
if systemctl is-active --quiet ${MYAPP_SERVICE}; then
  printf "INFO: ${MYAPP_SERVICE} service is running\n"
else
  printf "FAIL: ${MYAPP_SERVICE} service is not running\n"
  systemctl status ${MYAPP_SERVICE}
  exit 1
fi
