<?php
class Node{
    public $left,$right;
    public $data;
    function __construct($data)
    {
        $this->left=$this->right=null;
        $this->data = $data;
    }
}
class Solution{
    public function insert($root,$data){
        if($root==null){
            return new Node($data);
        }
        else{            
            if($data<=$root->data){
                $cur=$this->insert($root->left,$data);
                $root->left=$cur;
            }
            else{
                $cur=$this->insert($root->right,$data);
                $root->right=$cur;
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

    public function levelOrder($root) {
        //Write your code here
        if($root == null) {
            return;
        }
        
        $height = $this->getHeight($root);
        for($i=0; $i<=$height; $i++) {
            $this->printLevelOrder($root, $i);
        }
    }

    private function printLevelOrder($root, $i) {
        if($root == null) {
            return;
        }
        
        if($i == 0) {
            print($root->data.' ');
        } else if($i) {
            $this->printLevelOrder($root->left, $i-1);
            $this->printLevelOrder($root->right, $i-1);
        }
    }
}//End of Solution
$myTree=new Solution();
$root=null;
$T=intval(fgets(STDIN));
while($T-->0){
    $data=intval(fgets(STDIN));
    $root=$myTree->insert($root,$data);
}
$myTree->levelOrder($root);