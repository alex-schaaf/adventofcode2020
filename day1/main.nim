import strutils, sequtils

# read and parse file into sequence of int
let fileContents = readFile("input")
let lines = fileContents.splitLines()
var numbers: seq[int] = @[]
for i, line in lines:
    numbers.add(line.parseInt())

# exercise 1
proc solve(numbers: seq[int], sum: int): seq[int] =
    var sol: seq[int] = @[]
    for num in numbers:
        if (sum - num in numbers):
            sol.add(num)
            sol.add(sum - num)
            break
    return sol

let sol1 = solve(numbers, 2020)
echo sol1, sol1[0] * sol1[1]

# exercise 2
var num2: int
var num3: int
var sum: int

for i, num1 in numbers:
    for j in i+1 .. numbers.len() - 1:
        for k in j+1 .. numbers.len() - 1:
            num2 = numbers[j]
            num3 = numbers[k]
            sum = num1 + num2 + num3
            if sum == 2020:
                let prod = num1 * num2 * num3
                echo num1, " * ", num2," * ", num3," = ", prod
                break
