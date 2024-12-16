import math


def CalculateFunction(x, b):
    y = (1 + (math.sin(b**3 + x**3))**2) / ((b**3 + x**3)**(1/3))
    return y


def TaskA(xa_beginning, xa_end, xa_delta, b):
    y_a = []
    while xa_beginning != (xa_end + xa_delta):
        y_a.append(CalculateFunction(xa_beginning, b))
        xa_beginning += xa_delta
    return y_a


def TaskB(xs, b):
    y_b = []
    for i in xs:
        y_b.append(CalculateFunction(i, b))
    return y_b


print("Задача А", TaskA(1.28, 3.28, 0.4, 2.5))
xs = [1.1, 2.4, 3.6, 1.7, 3.9]
print("Задача В", TaskB(xs, 2.5))
