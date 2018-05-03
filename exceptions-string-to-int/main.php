<?php

$handle = fopen("php://stdin","r");
fscanf($handle, "%s", $S);

function badString() {
  print "Bad String\n";
  die;
}

set_error_handler('badString');


$testS = ((int)$S + ((int)$S === "0"));
$result = 1 / $testS;
echo $testS;