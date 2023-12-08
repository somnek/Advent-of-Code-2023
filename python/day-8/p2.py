import math


def do():
    init, rules = parse()
    at = [k for k in rules if k.endswith("A")]

    ll = []
    for a in at:
        c = 0
        while not a.endswith("Z"):
            move = init[c % len(init)]
            if move == "L":
                a = rules[a][0]
            else:
                a = rules[a][1]
            c += 1
        ll.append(c)
    ans = math.lcm(*ll)
    print(f"ans: {ans}")


def parse() -> (list, dict):
    with open("input.txt") as f:
        lines = f.readlines()

    init = lines[0].strip()
    rules = {}

    for line in lines[2:]:
        split = line.strip().split(" = ")
        col2 = split[1][1:-1].split(", ")
        rules[split[0]] = col2
    return init, rules


if __name__ == "__main__":
    do()
