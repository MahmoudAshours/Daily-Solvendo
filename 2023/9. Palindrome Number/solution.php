<?php

  /**
     * @param Integer $x
     * @return Boolean
     */
    function isPalindrome($x):bool{
	 $s = "$x";
	 if(strrev($s)===$s){
		return true;
	 }else{
		return false;
	 }

    }

print(isPalindrome(121));
?>
