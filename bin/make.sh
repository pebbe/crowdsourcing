#!/bin/sh

echo Content-type: text/plain
echo 

make 2>&1
make -C ../db new 2>&1
