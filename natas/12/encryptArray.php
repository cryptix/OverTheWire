<?php

function xor_encrypt($in) {
    $key = 'qw8J';
    $text = $in;
    $outText = '';

    // Iterate through each character
    for($i=0;$i<strlen($text);$i++) {
    $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }

    return $outText;
}
$data2 = array( "showpassword"=>"yes", "bgcolor"=>"#ff0000");
var_dump($data2);
var_dump(json_encode($data2));

$cook = base64_encode(xor_encrypt(json_encode($data2)));

echo $cook;
?>