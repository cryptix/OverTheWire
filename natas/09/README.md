# Level 9

Reverse the $encodedSecret. I stay with PHP for this:

> $encodedSecret = "3d3d516343746d4d6d6c315669563362";
> echo "Secret:".base64_decode(strrev(hex2bin($encodedSecret))). "\n";

Entering the Secret gives you:

> next pw: sQ6DKR8ICwqDMTd48lQlJfbF1q9B3edT

