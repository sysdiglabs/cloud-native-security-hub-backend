#!/bin/bash
#Usage: sh concatenator.sh aws-alb/1.0.0/alerts concat.json
PATHTOFILES=$1/*
echo "[" > $2
for file in $PATHTOFILES
do
  echo "making file $1"
  cat $file >> $2
  echo "
  ," >> $2
done
sed -i '$d' $2
echo "]" >> $2