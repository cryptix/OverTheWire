<?php
$ciph="ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSFlorExZaaAw";



function xor_encrypt($in) {
    $key = '<censored>';
    $text = $in;
    $outText = '';

    // Iterate through each character
    for($i=0;$i<strlen($text);$i++) {
    $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }

    return $outText;
}



// $tempdata = json_decode(xor_encrypt(base64_decode($_COOKIE["data"])), true);

echo "data:".$data;
$b64 = base64_decode($data);
echo "\n\nb64".$b64;

$ucrypt = xor_encrypt($b64);
echo "\n\nuncrypt:".$ucrypt;
$tempdata = json_decode($ucrypt, true);
var_dump($tempdata);
?>