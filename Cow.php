<?php
spl_autoload_register(function ($class_name) {
    include $class_name . '.php';
});

class Cow extends Animal{

    public $animalId;
    public $animalBreed = 'cows';    
    public $product = 'milk';

    function __construct() {
        
        $this -> animalId = parent::$id++;
        
    }

    
    public function getOutputProduct(){

        return rand(8, 12);

    }
}
