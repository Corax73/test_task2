<?php
class Farm
{
    public $animals;
    public $productionPerWeek = [];

    public function addAnimal(string $nameClassAnimal, int $number): void
    {
        for ($i = 0; $i < $number; $i++) {
            $typeAnimal = $this->createAnimal($nameClassAnimal)->animalBreed;
            if (!isset($this->animals[$typeAnimal])) {
                $this->animals[$typeAnimal] = [];
                $this->animals[$typeAnimal][] = $this->createAnimal($nameClassAnimal);
            } else {
                $this->animals[$typeAnimal][] = $this->createAnimal($nameClassAnimal);
            }
        }
    }

    public function createAnimal(string $nameClassAnimal): Animal
    {
        if (class_exists($nameClassAnimal)) {
            return new $nameClassAnimal;
        }
    }

    public function getProduction(): void
    {
        $this->productionPerWeek = [];
        for ($i = 0; $i <= 6; $i++) {
            foreach ($this->animals as $typeAnimal) {
                foreach ($typeAnimal as $anim) {
                    $typeProduct = $anim->product;
                    if (!isset($this->productionPerWeek[$typeProduct])) {
                        $this->productionPerWeek[$typeProduct] = 0;
                        $this->productionPerWeek[$typeProduct] += $anim->getOutputProduct();
                    } else {
                        $this->productionPerWeek[$typeProduct] += $anim->getOutputProduct();
                    }
                }
            }
        }
        $typeProduct = array_keys($this->productionPerWeek);
        foreach ($typeProduct as $product) {
            print $product . ' collected ' . $this->productionPerWeek[$product] . "\n";
        }
    }

    public function getCountAnimals(): void
    {
        $typeAnimals = array_keys($this->animals);
        foreach ($typeAnimals as $typeAnimal) {
            $count = count($this->animals[$typeAnimal]);
            print 'Count ' . $this->animals[$typeAnimal][0]->animalBreed . ' ' . $count . "\n";
        }
    }
}
