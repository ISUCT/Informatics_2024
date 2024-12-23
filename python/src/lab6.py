class Cat:
    def __init__(self, name, age, breed):
        self.name = name
        self.age = age
        self.breed = breed

    def UpdateAge(self, age):
        self.age = age

    def UpdateStruct(self, name, age, breed):
        self.name = name
        self.age = age
        self.breed = breed
    
    def printName(self):
        return self.name
    
    def printAge(self):
        return self.age
    
    def printBreed(self):
        return self.breed
    

def RunLab6Tasks():
    Tom = Cat("Tom", 9, "Scottish Fold")
    print("У нас есть кот по имени", Tom.printName(), "возрастом", Tom.printAge(), "лет и породы", Tom.printBreed())
    Tom.UpdateAge(13)
    print("Через 4 года коту", Tom.printName(), "будет", Tom.printAge(), "лет")
    Tom.UpdateStruct("Misa", 5, "Maine coon")
    print("Также у нас есть кошка по имени", Tom.printName(), "возрастом", Tom.printAge(), "лет и породы", Tom.printBreed())
