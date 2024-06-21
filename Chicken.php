<?php
spl_autoload_register(function ($class_name) {
    include $class_name . '.php';
});

class Chicken extends Animal
{
    public $animalId;
    public $animalBreed = 'chickens';
    public $product = 'eggs';

    function __construct()
    {
        $this->animalId = parent::$id++;
    }

    public function getOutputProduct()
    {
        return rand(0, 1);
    }
}
