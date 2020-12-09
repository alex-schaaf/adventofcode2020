import os
from datetime import datetime

template = """package main

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
    lines := readFileLines("./test", "\\n")
}
"""

day = datetime.now().strftime("%d")
folder = f"day{day}"
os.mkdir(folder)

with open(folder + "/main.go", "w") as file:
    file.write(template)

open(folder + "/test", "w").close()
open(folder + "/input", "w").close()