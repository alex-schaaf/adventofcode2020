from typing import List
from dataclasses import dataclass


def parse_input(filepath: str, split: str = "\n") -> List[str]:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split(split)


directions = {
    0: "N",
    90: "E",
    180: "S",
    270: "W",
}


@dataclass
class Ship:
    x: int
    y: int
    facing: int
    debug: bool = False

    def move(self, action: str, value: int):
        if action == "N":
            self.y += value
        elif action == "S":
            self.y -= value
        elif action == "E":
            self.x += value
        elif action == "W":
            self.x -= value
        elif action == "F":
            self.move(directions[self.facing], value)

    def turn(self, action: str, value: int):
        if action == "R":
            to = (self.facing + value) % 360
        else:  # L
            to = (self.facing - value) % 360
        if self.debug:
            print(f"Turning ship {action} from {self.facing} to {to}.")
        self.facing = to


@dataclass
class Waypoint:
    x: int
    y: int


if __name__ == "__main__":
    lines = parse_input("./input")

    ship = Ship(0, 0, 90, debug=False)
    for instruction in lines:
        action = instruction[0]
        value = int(instruction[1:])

        if ship.debug:
            print(action, value)
        if action in "NESWF":
            ship.move(action, value)
        else:
            ship.turn(action, value)
        if ship.debug:
            print(f"({ship.x},{ship.y})")

    manhatten = abs(ship.x) + abs(ship.y)
    print(f"Exercise 1: Manhatten distance is {manhatten}.")

    ship = Ship(0, 0, 90, debug=True)
    waypoint = Waypoint(10, 1)
    for instruction in lines:
        action = instruction[0]
        value = int(instruction[1:])
        if ship.debug:
            print(action, value)
        if action in "NESW":
            if action == "N":
                waypoint.y += value
            elif action == "E":
                waypoint.x += value
            elif action == "S":
                waypoint.y -= value
            elif action == "W":
                waypoint.x -= value
        elif action in "RL":
            if action == "R":  # clockwise
                for _ in range(value // 90):
                    x, y = waypoint.x, waypoint.y
                    waypoint.x = y
                    waypoint.y = -1 * x
            else:  # L counter-clockwise
                for _ in range(value // 90):
                    x, y = waypoint.x, waypoint.y
                    waypoint.x = -1 * y
                    waypoint.y = x
        elif action in "F":
            ship.x += waypoint.x * value
            ship.y += waypoint.y * value
        if ship.debug:
            print(f"ship: ({ship.x},{ship.y}), wp: ({waypoint.x},{waypoint.y})")

    manhatten = abs(ship.x) + abs(ship.y)
    print(f"Exercise 2: Manhatten distance is {manhatten}.")