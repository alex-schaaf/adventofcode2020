from typing import List
import re


def parse_input(filepath: str, split: str = "\n") -> List[str]:
    with open(filepath, "r") as file:
        content = file.read()
    return content.split(split)


def bti(bytecode: str) -> int:
    return int(bytecode, base=2)


def itb(integer: int) -> str:
    return "{0:b}".format(integer).zfill(36)


def apply(value: str, mask: str) -> str:
    result = ""
    for m, val in zip(mask, value):
        if m == "X":
            result += val
        else:
            result += m
    return result


def write_to_memory(value: str, pos: int, memory: List[int]) -> List[int]:
    if pos >= len(memory):
        for _ in range(len(memory), pos + 1):
            memory.append(0)
    memory[pos] = bti(value)
    return memory


if __name__ == "__main__":
    lines = parse_input("./input")

    # exercise 1
    memory: List[int] = []
    mask = ""
    for line in lines:
        arg, val = line.split(" = ")

        if "mem" in arg:
            arg = int(arg[4:].rstrip("]"))
            val = itb(int(val))
        else:
            mask = val
            continue

        # print(arg, "\t", val)
        # print("\t", mask)
        result = apply(val, mask)
        # print("\t", result)
        memory = write_to_memory(result, arg, memory)
        # print("")

    print(f"Exercise 1: The sum of all memory values is {sum(memory)}.")

    # exercise 2