#!/bin/bash

### BEGIN INIT INFO
# Required-Start:    $all
# Default-Start:     2 3 4 5
### END INIT INFO

DIALTIMEOUT=10

DR=`route -n | egrep '^0.0.0.0'| grep -v ppp | sed 's/^[^ ]*  *\([^ ]*\) .*/default gw \1/'` ;
 if [ -n "$DR" ] ; then
   trap "echo route add $DR ; route add $DR ; exit"  2 3 9 15
    route delete $DR
    echo route delete $DR
 fi

while  true ; do
    wvdial
    sleep $DIALTIMEOUT
done