from copy import deepcopy
from typing import Callable, List


def parse_input(filepath: str, split: str = "\n") -> list:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split(split)


coords = ((1, -1), (1, 0), (1, 1), (0, -1), (0, 1), (-1, -1), (-1, 0), (-1, 1))


def countNeighbors(seats: list, y: int, x: int) -> int:
    occupied = 0
    for dy, dx in coords:
        if y + dy < 0:
            continue
        elif y + dy >= len(seats):
            continue
        if x + dx < 0:
            continue
        elif x + dx >= len(seats[0]):
            continue
        neighbor_seat = seats[y + dy][x + dx]
        if neighbor_seat == "#":
            occupied += 1
    return occupied


def countNeighborsSight(seats: list, y: int, x: int) -> int:
    occupied = 0
    ny = len(seats)
    nx = len(seats[0])
    for dy, dx in coords:
        i = 1
        while 0 <= y + dy * i < ny and 0 <= x + dx * i < nx:
            seat = seats[y + dy * i][x + dx * i]
            if seat == "#":
                occupied += 1
                break
            elif seat == "L":
                break
            i += 1
    return occupied


def solve(seats: List[List[str]], count_func: Callable, n_occupied: int) -> int:
    while True:
        newSeats = deepcopy(seats)
        # print(f"Iteration {i}")
        changed = 0
        for y, row in enumerate(seats):
            # print("".join(row), end="\t")
            for x, seat in enumerate(row):
                n = count_func(seats, y, x)
                # print(n, end="")
                if seat == "L" and n == 0:
                    # -> occupied
                    newSeats[y][x] = "#"
                    changed += 1
                    continue
                if seat == "#" and n >= n_occupied:
                    # -> empty
                    newSeats[y][x] = "L"
                    changed += 1
                    continue
            # print("")
        seats = deepcopy(newSeats)

        if changed == 0:
            break

    occupied = 0
    for row in seats:
        for seat in row:
            if seat == "#":
                occupied += 1

    return occupied


if __name__ == "__main__":
    lines = parse_input("./input")
    seats = []
    for line in lines:
        seats.append([seat for seat in line])

    sol1 = solve(seats, countNeighbors, 4)
    print(f"Exercise 1: {sol1} occupied seats.")
    sol2 = solve(seats, countNeighborsSight, 5)
    print(f"Exercise 2: {sol2} occupied seats.")