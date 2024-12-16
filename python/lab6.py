class plane:
    Name: str
    Flight_number: int
    Power: int
    Purchased_seats: int

    def __init__(self, name: str, flight_number: int, power: int,
                 purchased_seats: int):
        self.Name = name
        self.Flight_number = flight_number
        self.Power = power
        self.Purchased_seats = purchased_seats

    def speed_calculation(self, f_air: int):
        speed = self.Power / f_air
        return speed

    def revenue_calculation(self, price_a_place: int):
        money = self.Purchased_seats * price_a_place
        return money

    def get_routes(self):
        routes = {1: "Paris-Moscow", 2: "Moscow-Paris", 3: "Moscow-Seoul",
                  4: "Seoul-Moscow"}
        nomer = routes.get(self.Flight_number)
        return nomer


Plane = plane("Airobus999", 3, 200000, 150)
print("Скорось", Plane.Name, "равна", Plane.speed_calculation(5000), "м/с")
print("Прибыль с продажи белетов составляет:",
      Plane.revenue_calculation(23000), "рублей")
print(Plane.Name, "летит по маршруту", Plane.get_routes())
