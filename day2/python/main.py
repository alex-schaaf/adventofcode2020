def exercise1(filepath: str) -> int:
    with open(filepath, "r") as file:
        lines = file.readlines()
    lines = [line.rstrip() for line in lines]

    valid_passwords = 0
    for line in lines:
        minmax, char, password = line.split()
        min_, max_ = minmax.split("-")
        min_, max_ = int(min_), int(max_)
        char = char[0]

        count = countOccurance(char, password, 0)
        if min_ <= count <= max_:
            valid_passwords += 1

    return valid_passwords


def countOccurance(char: str, string: str, count: int) -> int:
    if len(string) == 0:
        return count
    if string[-1] == char:
        count += 1
    return countOccurance(char, string[:-1], count)


if __name__ == "__main__":
    filepath = "../input"
    sol1 = exercise1(filepath)
    print(f"Exercise 1: Found {sol1} valid passwords.")