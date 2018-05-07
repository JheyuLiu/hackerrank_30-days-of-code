<?php

class Node {
	public $left, $right;
	public $data;

	function __construct($data)
	{
		$this->left = $this->right = null;
		$this->data = $data;
	}
}

class Solution {
	public function insert($root, $data)
	{
		if($root == null) {
			return new Node($data);
		} else {
			if($data <= $root->data) {
				$curr = $this->insert($root->left, $data);
				$root->left = $curr;
			} else {
				$curr = $this->insert($root->right, $data);
			}
			return $root;
		}
	}

	public function getHeight($root) {
		if($root == null){
            return 0;
        }
        
        $max_left_depth = $this->depth_search($root->left);
        $max_right_depth = $this->depth_search($root->right);
        $max_depth = $max_left_depth > $max_right_depth ? $max_left_depth : $max_right_depth;
        
        return $max_depth;
	}

	private function depth_search($root) {
        if($root == null) {
            return 0;
        }
        
        $left_count = $this->depth_search($root->left) + 1;
        $right_count = $this->depth_search($root->right) + 1;
        
        return $left_count > $right_count ? $left_count : $right_count;
    }
}

$myTree = new Solution();
$root = null;
$T = intval(fgets(STDIN));
while($T-- > 0) {
	$data = intval(fgets(STDIN));
	$root = $myTree->insert($root, $data);
}
$height = $myTree->getHeight($root);
echo $height;