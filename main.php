<?php
spl_autoload_register(function ($class_name) {
    include $class_name . '.php';
});

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

