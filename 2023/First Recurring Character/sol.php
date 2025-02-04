<?php

function solve(string $s){
	$freq_map=[];
	foreach(str_split($s) as $index=>$char){
	if(array_key_exists($char,$freq_map)){
	return $char;
	}else{
	$freq_map[$char] = 1;
	continue;
	}
}
}
print(solve("ABCD"));
?>
