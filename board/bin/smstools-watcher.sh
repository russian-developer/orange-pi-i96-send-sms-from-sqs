#!/bin/bash

### BEGIN INIT INFO
# Required-Start:    $all
# Default-Start:     2 3 4 5
### END INIT INFO

USB_DEVICE=/dev/ttyUSB1

while true ; do
  if [ ! -f "$USB_DEVICE" ]; then
    echo "Dongle disconnected"
    while [ ! -f "$USB_DEVICE" ]; do
      echo "Dongle still disconnected..."
      sleep 1
    done
    echo "Dongle connected. Restart smstools"
    /etc/init.d/smstools restart
  fi
  echo "Dongle connected"
  sleep 1
done