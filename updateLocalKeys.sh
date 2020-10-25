#!/bin/sh

if [[ ! -d data ]]
then
  echo "missing data folder"
  exit 1
fi

cd data
echo "delete existing keys"
rm -rf "2020-*"
echo "get current keys"
scp -r "user@mqtt.pantra.eu:./cwa/*" .
echo "done"
exit 0
