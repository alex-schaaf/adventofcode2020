from typing import List
import re

fields = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
required_fields = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}


def parse_passports(filepath: str) -> List[str]:
    with open(filepath, "r") as file:
        content = file.read()
    passport_strings = content.split("\n\n")
    passport_strings = [pp.replace("\n", " ") for pp in passport_strings]
    return passport_strings


def exercise1(filepath: str) -> int:
    passport_strings = parse_passports(filepath)

    valid_passports = 0
    for passport in passport_strings:
        score = 0
        for field in required_fields:
            if field in passport:
                score += 1
        if score == 7:
            valid_passports += 1
    return valid_passports


def exercise2(filepath: str):
    passport_strings = parse_passports(filepath)

    passports = []
    for i, passport_str in enumerate(passport_strings):
        passport = {}
        entries = passport_str.split()
        for entry in entries:
            key, value = entry.split(":")
            passport[key] = value
        passports.append(passport)

    valid_passports = 0
    for i, passport in enumerate(passports):
        print(i, passport_strings[i])
        # ---
        byr = passport.get("byr", -1)
        if not (1920 <= int(byr) <= 2002):
            print("Invalid BYR", byr)
            continue
        # ---
        iyr = passport.get("iyr", -1)
        if not (2010 <= int(iyr) <= 2020):
            print("Invalid IYR", iyr)
            continue
        # ---
        eyr = passport.get("eyr", -1)
        if not (2020 <= int(eyr) <= 2030):
            print("Invalid EYR", eyr)
            continue
        # ---
        hgt = passport.get("hgt")
        if hgt is None:
            print("Invalid HGT", hgt)
            continue

        if len(hgt) <= 2:
            print("Invalid HGT", hgt)
            continue
        hgt_unit = hgt[-2:]
        hgt_val = int(hgt[:-2])
        if hgt_unit == "cm":
            if not 150 <= hgt_val <= 193:
                print("Invalid HGT", hgt)
                continue
        elif hgt_unit == "in":
            if not 59 <= hgt_val <= 76:
                print("Invalid HGT", hgt)
                continue
        else:
            print("Invalid HGT", hgt)
            continue
        # ---
        hcl = passport.get("hcl", None)
        if hcl is None:
            print("Invalid HCL", hcl)
            continue

        pattern = "#[a-z0-9]{6}"
        x = re.search(pattern, hcl)
        if x is None:
            print("Invalid HCL", hcl)
            continue
        # ---
        ecl = passport.get("ecl")
        if ecl not in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]:
            print("Invalid ECL", ecl)
            continue
        # ---
        pid = passport.get("pid")
        if pid is None:
            print("Invalid PID", pid)
            continue
        pattern = "[0-9]{9}"
        x = re.search(pattern, pid)
        if x is None:
            print("Invalid PID", pid)
            continue

        valid_passports += 1

    return valid_passports - 1


if __name__ == "__main__":
    sol1 = exercise1("./input")
    sol2 = exercise2("./input")
    print(f"Exercise 1: {sol1} valid passports found.")
    print(f"Exercise 2: {sol2} valid passports found.")
