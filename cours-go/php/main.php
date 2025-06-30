

<?php

// $function = "welcome";
// FUNCTION test() {
//     $name = 9;
//     global $function;
//     echo "<h1>" . $function .  $name ."</h1>\n";
//     print ($function . $name);
// }
// test();
// echo phpversion();

// $x = 5;
// $y = 10;

// function myTest() {
//     $GLOBALS['y'] = $GLOBALS['x'] + $GLOBALS['x'];
// }
// myTest();
// echo $y;


// function Test1() {
//      static $x = 0;
//      $x++;
//     echo $x;
// }

// Test1();
// Test1();
// Test1();

// $x = 5;
// $s = "hello";
// $f = 665.0090;
// $a = array($x, $s, $f);
// $b =  true;

// var_dump($x);
// var_dump($s);
// var_dump($f);
// var_dump($b);
// var_dump($a);


// class Car {
//     public $color;
//     public $model;
//     public function __construct($color, $model) {
//         $this -> color = $color;
//         $this -> model = $model;
//     }
//     public function message() {
//         return "My car is a " . $this->$color . " " . $this->model . "!";
//     }
// }


// $myCar = new Car("red", "volvo");

// var_dump($myCar);
// $x = null;
// var_dump($x);

// $b = (object) $b;
// var_dump($b)

// echo strlen("omar");
// echo "\n";
// echo str_word_count("omar el haouch");
// echo "\n";
// echo strpos("omar el haouch", "el"); // search for index of start word
// echo "\n";
// echo strtoupper("omar");
// echo "\n";
// echo strtolower("OMAR");

// echo "\n";
// echo str_replace("haouch", "sony", "omar el haouch" );
// echo "\n";
// echo strrev("OMAR");
// echo "\n";
// echo trim("OMAR                ");

// // Convert string to array
// print_r (explode(" ", "omar el"));
// echo "omar" . "el haouch"; // concat string
// // slicing
// echo substr("omar el haouch", 0, -4);
// echo "\n";
// echo substr("omar el \"\" haouch", 6);



$a = 5;
$b = 5.34;
$c = "25";
var_dump($a);
var_dump($b);
var_dump($c);

// the some
var_dump(is_int(9.9));
var_dump(is_integer(9));
var_dump(is_long(9));

ECHO PHP_INT_MAX;
ECHO "\n";
ECHO PHP_INT_MIN;
ECHO "\n";
echo PHP_INT_SIZE;
ECHO "\n";

// the some 
var_dump(is_float(9.9));
var_dump(is_double(9.9));

echo PHP_FLOAT_DIG;
ECHO "\n";
echo PHP_FLOAT_MAX;
ECHO "\n";
echo PHP_FLOAT_MIN;
ECHO "\n";
echo PHP_FLOAT_EPSILON;

echo is_finite(2);
ECHO "\n";
echo is_finite(log(0));
ECHO "\n";
echo is_finite(2000);
ECHO "\n";

echo (int) 9.9;
ECHO "\n";
// cat to  null no supported in this version
// echo (unset) 9;

 
?>
