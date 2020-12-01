import strutils, sequtils

# read and parse file into sequence of int
let fileContents = readFile("input")
let lines = fileContents.splitLines()
var numbers: seq[int] = @[]
for i, line in lines:
    numbers.add(line.parseInt())

# exercise 1
for num in numbers:
    if (2020 - num in numbers):
        echo num, " * ", 2020 - num, " = ", num * (2020 - num)
        break