import numpy as np


def parse_input(filepath: str) -> list:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split("\n")


def count_neighbors(y: int, x: int, z: int, w: int, world: np.ndarray) -> int:
    neighbors = 0
    for dy in range(-1, 2):
        for dx in range(-1, 2):
            for dz in range(-1, 2):
                for dw in range(-1, 2):
                    if dy == 0 and dx == 0 and dz == 0 and dw == 0:
                        continue
                    try:
                        neighbors += world[y + dy, x + dx, z + dz, w + dw]
                    except IndexError:
                        continue
    return neighbors


if __name__ == "__main__":
    lines = parse_input("./input")

    ny = len(lines)
    nx = len(lines[0])

    # initialize world
    world = np.zeros((ny, nx, 1, 1))
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            if char == "#":
                world[y, x, 0] = 1

    # cycle world
    for cycle in range(6):
        world = np.pad(world, 1, "constant", constant_values=0)

        ny, nx, nz, nw = world.shape

        new_world = np.copy(world)

        for y in range(ny):
            for x in range(nx):
                for z in range(nz):
                    for w in range(nw):
                        active = bool(world[y, x, z, w])
                        n = count_neighbors(y, x, z, w, world)
                        if active:
                            if 2 <= n <= 3:
                                continue
                            else:
                                new_world[y, x, z, w] = 0
                        else:
                            if n == 3:
                                new_world[y, x, z, w] = 1
        world = new_world

    print(np.count_nonzero(world))