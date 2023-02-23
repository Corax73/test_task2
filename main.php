<?php
spl_autoload_register(function ($class_name) {
    include $class_name . '.php';
});

//namespace farm;

/*require 'Animal.php';
require 'Chicken.php';
require 'Cow.php';
require 'Farm.php';

use Animal;
use Chicken;
use Cow;
use Farm;*/



$numberCow = 10;
$numberChicken = 20;

$farm = New Farm();

$farm -> addAnimal('Cow', $numberCow);
$farm -> addAnimal('Chicken', $numberChicken);

$farm -> getProduction();

$farm -> getCountAnimals();

print 'Went to the market, bought animals' . "\n";

$numberCow = 1;
$numberChicken = 5;

$farm -> addAnimal('Cow', $numberCow);
$farm -> addAnimal('Chicken', $numberChicken);

$farm -> getProduction();

$farm -> getCountAnimals();

