#!/bin/sh

# file: spp_web_stop.sh
#
# description...
#

# make sure this script is run as root
#
if (test ${EUID} -ne 0); then
    printf "$0: Error: Permission denied.\n"
    printf "Make sure this script is run as root.\n"
    exit 1
fi

# stop webserver
#
sudo docker stop spp_webserver

# exit normally
#
exit 0

#
# end of file
