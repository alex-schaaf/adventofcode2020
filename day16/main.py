from typing import List, Tuple


def parse_input(filepath: str) -> Tuple:
    with open(filepath, "r") as file:
        content = file.read()
    parts = content.split("\n\n")
    rules = parts[0].split("\n")
    my_ticket = parts[1].split("\n")[1]
    tickets = parts[2].split("\n")[1:]

    rules_parsed = {}
    for rule in rules:
        name, ranges = rule.split(": ")
        ranges = ranges.split(" or ")

        rs = []
        for r in ranges:
            r1, r2 = r.split("-")
            rs.append((int(r1), int(r2)))
        rules_parsed[name] = rs

    my_ticket = [int(n) for n in my_ticket.split(",")]
    tickets_parsed = []
    for ticket in tickets:
        tickets_parsed.append([int(n) for n in ticket.split(",")])
    return rules_parsed, my_ticket, tickets_parsed


if __name__ == "__main__":
    rules, my_ticket, tickets = parse_input("./test")
    print(rules)
    print(my_ticket)
    print(tickets)