<?php


$defaultdata = array( "showpassword"=>"no", "bgcolor"=>"#ffffff");
$plain  = json_encode($defaultdata);
$cipher = base64_decode("ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSEV4sFxFeaAw=");
echo "\nplain:". strlen($plain);
echo "\ncipher:". strlen($cipher);


$key = '';
for ($i=0; $i < strlen($plain); $i++) { 
	$key .= $plain[$i] ^$cipher[$i];
}

echo "\n".$key;
?>