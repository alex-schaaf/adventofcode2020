# python 3.9


def parse_input(filepath: str, split: str = "\n") -> list[str]:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split(split)


def parse_bags(lines: list[str]) -> dict[str, list[tuple[int, str]]]:
    bags = {}

    for line in lines:
        bag_name, children = line.split(" bags contain ")

        bags[bag_name] = []
        if children == "no other bags.":
            bags[bag_name].append((0, None))
        else:
            for child in children.split(", "):
                split = child.split()
                n = int(child[0])
                name = split[1] + " " + split[2]
                bags[bag_name].append((n, name))

    return bags

def contains(bags_all: dict, bag: str, target: str = "shiny gold"):
    contained_bags = []
    for n, child_name in bags_all[bag]:
        contained_bags.append(child_name)
    if target in contained_bags:
        return True
    elif contained_bags[0] is None:
        return False
    else:
        return max([contains(bags_all, n) for n in contained_bags])


def count_containing(bags_all: dict, bag: str) -> int:
    n_bags, contained_bags = zip(*bags_all.get(bag))
    if sum(n_bags) == 0:
        return 0
    else:
        count = 0
        for i, (n, contained_bag) in enumerate(zip(n_bags, contained_bags)):
            n_child = count_containing(bags_all, contained_bag)
            count += n + n * n_child
        return count


if __name__ == "__main__":
    lines = parse_input("./input")
    bags = parse_bags(lines)

    counter = 0
    for bag in bags:
        counter += contains(bags, bag)

    print(f"Exercise 1: {counter}")
    counter2 = count_containing(bags, "shiny gold")
    print(f"Exercise 2: {counter2}")