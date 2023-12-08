def do():
    init, rules = parse()
    lost, at, c = True, "AAA", 0

    while lost:
        for d in init:
            c += 1
            if d == "L":
                at = rules[at][0]
            elif d == "R":
                at = rules[at][1]

            if at == "ZZZ":
                lost = False

    print(f"ans: {c}")


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
