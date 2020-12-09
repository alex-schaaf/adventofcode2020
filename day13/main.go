package main

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
    lines := readFileLines("./test", "\n")
}
