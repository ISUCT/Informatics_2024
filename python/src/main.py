from lab4 import task_a, task_b
def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))
       
      
    resultA = task_a(0.05, 0.06, 0.2, 0.95, 0.15)
    arr = [0.15, 0.26, 0.37, 0.48, 0.56]
    for result in resultA:
        print(f"x: {result[0]:.2f}, y: {result[1]:.2f}")
       
    resultB = task_b(0.2, 0.95, arr)
    for result in resultB:
        print(f"x: {result[0]:.2f}, y: {result[1]:.2f}")
   
