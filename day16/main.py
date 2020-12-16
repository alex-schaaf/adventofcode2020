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


def check_or(ranges: List[Tuple[int, int]], value: int) -> bool:
    p1, p2 = False, False
    (r1, r2), (r3, r4) = ranges[0], ranges[1]
    if r1 <= value <= r2:
        p1 = True
    if r3 <= value <= r4:
        p2 = True
    return p1 or p2


if __name__ == "__main__":
    rules, my_ticket, tickets = parse_input("./input")
    # print(rules)
    # print(my_ticket)
    # print(tickets)
    invalid_sum = 0
    for ticket in tickets:
        valid_values = []
        for value in ticket:
            passed_any = 0
            for rule, ranges in rules.items():
                passed_any += check_or(ranges, value)
            valid_values.append(bool(passed_any))
            if not bool(passed_any):
                invalid_sum += value
        # print(ticket)
        # print(valid_values)
    print(invalid_sum)