<?php

class Difference{
    private $elements=array();
    public $maximumDifference;

    function __construct($arr) {
        $this->elements = $arr;
    }
 
    function ComputeDifference() {
        sort($this->elements);

        $first = current($this->elements);
        $this->maximumDifference =  end($this->elements) - $first;
    }
}

$N=intval(fgets(STDIN));
$a =array_map('intval', explode(' ', fgets(STDIN)));
$d=new Difference($a);
$d->ComputeDifference();
print ($d->maximumDifference);
