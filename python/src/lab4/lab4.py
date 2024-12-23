import math


def Calculate(b, x: float) -> float:
    return (1 + (math.sin(b ** 3 + x ** 3) ** 2)) / ((b ** 3 + x ** 3) ** (1 / 3))


def TaskA(b, Xn, Xk, delX: float) -> list:
    arr = []
    while Xn <= Xk:
        arr.append(Calculate(b, Xn))
        Xn += delX
    return arr


def TaskB(b: float, x: list) -> list:
    arr = []
    for elem in x:
        arr.append(Calculate(b, elem))
    return arr


def PrintValue(Values):
    for elem in Values:
        print(elem)


def RunLab4Tasks():
    b = 2.5
    arr = [1.1, 2.4, 3.6, 1.7, 3.9]
    PrintValue(TaskA(b, 1.28, 3.28, 0.4))
    PrintValue(TaskB(b, arr))
