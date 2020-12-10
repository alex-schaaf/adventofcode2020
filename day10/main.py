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

    # ----------------------------------------------
    # exercise 2
    # ----------------------------------------------
    def differences(arr: list[int]) -> list[int]:
        results = []
        for i,j in zip(arr[1:], arr[:-1]):
            results.append(i-j)
        return results

    joltages = [int(l) for l in lines] 
    joltages = [0] + sorted(joltages) + [max(joltages)]

    print(differences(joltages))
    string = "".join(str(d) for d in differences(joltages))
    
    count = {"1111": 0, "111": 0, "11": 0, "1": 0}
    while string:
        if string[:4] == "1111":
            count["1111"] += 1
            string = string[4:]
            continue
        if string[:3] == "111":
            count["111"] += 1
            string = string[3:]
            continue
        if string[:2] == "11":
            count["11"] += 1
            string = string[2:]
            continue
        if string[:1] == "1":
            count["1"] += 1
            string = string[1:]
        else:
            string = string[1:]

    # "1111" can be 
    # (4)
    # (3,1)
    # (2,2)
    # (1,3)
    # (1,2,1)
    # (2,1,1)
    # (1,1,2)

    # "111" can be
    # (3)
    # (1,1,1)
    # (2, 1)
    # (1, 2)
    print(7**count.get("1111") * 4**count.get("111") * 2**count.get("11") )