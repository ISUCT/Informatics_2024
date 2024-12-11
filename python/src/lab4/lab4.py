import math

def CalculateFunction(a: float, b: float, x: float) -> float:
    if math.sin(a+b*x) < 0:
        return float('NaN')
    
    return (math.pow(math.sin(a+b*x), 3.5)) / (1 + math.cos(math.log10(a+b*x)))

def CompleteTaskA(a: float, b: float, xMin: float, xMax: float, xDelta: float) -> list[float]:
    result: list[float] = []
    x = xMin
    while x < xMax:
        result.append(CalculateFunction(a, b, x))
        x += xDelta
    return result
        

def CompleteTaskB(a: float, b: float, x: list[float]) -> list[float]:
    result: list[float] = []
    for x in x:
        result.append(CalculateFunction(a, b, x))
    return result

def CompleteLab4() -> None:
    a, b = 2.5, 4.6
    xMin = 1.15
    xMax = 3.05
    xDelta = 0.38
    x = [1.2, 1.36, 1.57, 1.93, 2.25]

    print(CompleteTaskA(a, b, xMin, xMax, xDelta))
    print(CompleteTaskB(a, b, x))
 