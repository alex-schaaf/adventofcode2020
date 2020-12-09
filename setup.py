import os

template = """package main

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
    lines := readFileLines("./test", "\\n")
}
"""

for n in range(11, 26):
    print(n)
    folder = f"day{n}"
    os.mkdir(folder)

    with open(folder + "/main.go", "w") as file:
        file.write(template)

    f = open(folder + "/test", "w")
    f.close()
    f = open(folder + "/input", "w")
    f.close()