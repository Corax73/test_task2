<?php
spl_autoload_register(function ($class_name) {
    include $class_name . '.php';
});

abstract class Animal{
    
    static $id = 0;
    
    public $animalId; 
    
    public abstract function getOutputProduct();
}
