#!/bin/bash

set -e

cat > /etc/resolv.conf <<EOF
nameserver 8.8.8.8
nameserver 8.8.4.4
EOF

apt-get update
apt-get install usb-modeswitch wvdial smstools wget

sed -i 's/^auto wlan0/# auto wlan0/' /etc/network/interfaces

cp bin/sms-sqs /usr/local/bin
cp bin/smstools-watcher.sh /usr/local/bin
cp bin/wvdial-watcher.sh /usr/local/bin

cp etc/sms-sqs.conf /etc
cp etc/smsd.conf /etc
cp etc/wvdial.conf /etc
cp etc/usb_modeswitch.d/05c6:1000:uMa=Qualcomm /etc/usb_modeswitch.d

cp etc/systemd/system/wvdial-watcher.service /etc/systemd/system
cp etc/systemd/system/smstools-watcher.service /etc/systemd/system
cp etc/systemd/system/sms-sqs.service /etc/systemd/system

systemctl daemon-reload
systemctl enable smstools-watcher.service
systemctl start smstools-watcher.service
systemctl enable wvdial-watcher.service
systemctl start wvdial-watcher.service
systemctl enable sms-sqs.service
systemctl start sms-sqs.service
