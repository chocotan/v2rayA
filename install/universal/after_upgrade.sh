#!/bin/sh

systemctl daemon-reload
systemctl restart v2raya

ECHOLEN=$(echo -e|awk '{print length($0)}')
if [ ${ECHOLEN} = '0' ]
then
    ECHO='echo -e'
else
    ECHO='echo'
fi;
$ECHO "\033[36m**************************************\033[0m"
$ECHO "\033[36m*         Congratulations!           *\033[0m"
$ECHO "\033[36m* HTTPS demo: https://v2raya.mzz.pub *\033[0m"
$ECHO "\033[36m* HTTP  demo: http://v.mzz.pub       *\033[0m"
$ECHO "\033[36m**************************************\033[0m"