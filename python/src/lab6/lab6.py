class Car:
    Name: str
    Color: str
    Speed: int
    
    def __init__(self, name, color: str, speed: int) -> None:
        self.Name = name
        self.Color = color
        self.Speed = speed
    
    def SetSpeed(self, speed: int) -> None:
        self.Speed = speed
    
    def GetSpeed(self) -> int:
        return self.Speed
    
    def GetColor(self) -> str:
        return self.Color
    
def CompleteLab6() -> None:
    chevrolet = Car("Шевроле", "синий", 100)
    chevrolet.SetSpeed(200)
    print(chevrolet.GetSpeed())
    print(chevrolet.GetColor())