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

if [ ! -e ${MYAPP_CONF_DIR}/hosts ]; then
  cat << EOF > ${MYAPP_CONF_DIR}/hosts
myswitch os=cisco_nxos host_overwrite=127.0.0.1
EOF
fi

if [ ! -e ${MYAPP_CONF_DIR}/vault.yml ]; then
  cat << 'EOF' > ${MYAPP_CONF_DIR}/vault.yml
$ANSIBLE_VAULT;1.1;AES256
32326463356266366337623135303463636434316263643961616531303139666333623463636665
3561383561623735303533373734626234303165383538620a626637353365613633303439306136
62336638326562336538396165396662356465396633646664313263636164636437386137356461
6165316437396133640a303261373437356539666638613333373035663961393135613939316165
35653032316630383935653963366265313136663933333732653233373633383138643661616538
36626461633266626561333437623132633461666537303833313931306366633164623339363237
31666133326431383536643639333562333938356535363330613565616230623031323830626564
35346533643066653231356662303134623266356631386331303939663964316639303362343739
30613862663935623036396536363739633966353136633161313738376662653239643964316631
61616137666438643363363836393663396130343562316638613331623464613532336333306639
63343938313862623765303838383633646135313866656664616336623734343461653361333531
36663564363364663332613030383161366666383830383730613763396665353633623863633439
63336632333566646138656637636265626531373536376330393161313131656261653439633338
33313534323733323739613537656462323031393731383864323965353933643262366239626362
39656630336563616463643135333065303232633363373932376539316562323934633939383864
36363136366636663165303837653138383932653436343238333066313663636463396163613464
30333133663537303339373236326438396638373962646133623561383663646436373763383836
65633263326565373834626532656666633837363236666137663730396161303965663532323036
61313131656332393939333761333535653162353538376332626665343538633865313139613836
63643163633338613438643935316335396535626363623737633534663632366330613135643639
36633765363430663036656263393162333336383266366234396661373130363939373862656139
64643862306431396533663439633166626132343530363739646331666235313337633237613638
66386237633935343264643832313432636361316366313232663236393736323862373034376264
33353865376161636265666235626636666532623335386531313266393532653138636532373339
38303332633933633263323862303136356365636330373636346162353361356163303861386165
62316361393533363239323262666434346338393437643533303138343134396463303631386433
34353463303835376632333238386665356637383131373733613830613266386362323463393737
36323537646432306131616530313566323236663430643561343465333861363930633563343135
38633562643234353764663461303131656230633638373034303436356630633066613563333637
30636338346539313536346431373037636634353465663732363536623664303230666565376434
30356132363731646534626662373039643665313532346531396166333065643562623439313434
34386238363336663563633337643834633639653035303134303233333939323632333532333365
3831
EOF
  cat << EOF > ${MYAPP_CONF_DIR}/vault.key
7f017fde-e88b-42c5-89df-a7c8f9de981d
EOF
fi

cat << EOF > ${MYAPP_SYSCONF}
OPTIONS="-log.level info \
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
