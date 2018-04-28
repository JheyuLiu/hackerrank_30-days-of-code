<?php

class Person {
	protected $firstName, $lastName, $id;

	public function __construct($first_name, $last_name, $identification) {
		$this->firstName = $first_name;
		$this->lastName = $last_name;
		$this->id = $identification;
	}

	function printPerson() {
		print("Name: {$this->lastName}, {$this->firstName}\n");
		print("ID: {$this->id}\n");
	}
}

class Student extends Person {
	private $testScores;

    public function __construct($first_name, $last_name, $identification, $scores) {
    	parent::__construct($first_name, $last_name, $identification);
    	$this->testScores = $scores;
    }

	function calculate() {
		$average_score = 0;
        $sum = 0;
        $len = count($this->testScores);
		for($i = 0; $i < $len; $i++) {
			$sum += $this->testScores[$i];
		}

        $average_score = $sum / $len;
		if($average_score >= 90 && $average_score <= 100){
            return "O";
        }else if($average_score >= 80 && $average_score < 90){
            return "E";
        }else if($average_score >= 70 && $average_score < 80){
            return "A";
        }else if($average_score >= 55 && $average_score < 70){
            return "P";
        }else if($average_score >= 40 && $average_score < 55){
            return "D";
        }else if($average_score < 40){
            return "T";
        }
	}
}

$file = fopen("stdin", "r");

$name_id = explode(' ', fgets($file));

$first_name = $name_id[0];
$last_name = $name_id[1];
$id = (int)$name_id[2];

fgets($file);

$scores = array_map(intval, explode(' ', fgets($file)));

$student = new Student($first_name, $last_name, $id, $scores);

$student->printPerson();

print("Grade: {$student->calculate()}");