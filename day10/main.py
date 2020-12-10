# python 3.9
from collections import Counter


def parse_input(filepath: str, split: str = "\n") -> list[str]:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split(split)


if __name__ == "__main__":
    lines = parse_input("./input")
    joltages = [int(l) for l in lines]

    graph = [0]
    differences = []
    while joltages:
        indices = []
        diffs = []

        for i, j in enumerate(joltages):
            diff = j - graph[-1]
            if 1 <= diff <= 3:
                diffs.append(diff)
                indices.append(i)
        
        i = indices[diffs.index(min(diffs))]
        differences.append(min(diffs))
        graph.append(joltages[i])        
        joltages.pop(i)

    graph.append(graph[-1] + 3)
    differences.append(3)

    count = Counter(differences)
    for k, v in count.items():
        print(f"{k} -> {v}")

    print(f"Exercise 1: {count.get(1) * count.get(3)}")