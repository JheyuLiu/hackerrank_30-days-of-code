<?php

class Solution {
    private $stacks = array(), $queue = array();
    private $slen = 0, $qstart = 0, $qlast = 0;
    
    public function pushCharacter($ch)
    {
        $this->stacks[$this->slen++] = $ch;
    }
    
    public function popCharacter()
    {
        if ($this->slen < 1) {
            return '';
        }
        return $this->stacks[--$this->slen];
    }
    
    public function enqueueCharacter($ch)
    {
        $this->queue[$this->qlast++] = $ch;
    }
    
    public function dequeueCharacter()
    {
        if ($this->qstart == $this->qlast) {
            $this->qstart = 0;
            $this->qlast = 0;
        }
        return $this->queue[$this->qstart++];
    }
}

// read the string s
$s = fgets(STDIN);

// create the Solution class object p
$obj = new Solution();

$len = strlen($s);
$isPalindrome = true;

// push/enqueue all the characters of string s to stack
for ($i = 0; $i < $len; $i++) {
    $obj->pushCharacter($s{$i});
    $obj->enqueueCharacter($s{$i});
}

/*
pop the top character from stack
dequeue the first character from queue
compare both the characters
*/
for ($i = 0; $i < $len / 2; $i++) {
    if($obj->popCharacter() != $obj->dequeueCharacter()){
        $isPalindrome = false;
        
        break;
    }
}

//finally print whether string s is palindrome or not.
if ($isPalindrome)
    echo "The word, ".$s.", is a palindrome.";
else
    echo "The word, ".$s.", is not a palindrome.";
