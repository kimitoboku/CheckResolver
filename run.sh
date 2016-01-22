#!/bin/sh

if [ -e ./checked ]; then
    rm ./checked
fi

filename=$1
cat ${filename} | while read line
do
    ./CheckResolver ${line} >> checked
done
