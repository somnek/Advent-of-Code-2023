def do():
    with open("input.txt") as f:
        lines = f.read().strip().split("\n")

    ans = 0
    for line in lines:
        line = clean_line(line)
        tail = solve(line[::-1])
        ans += tail

    print(f"ans: {ans}")


def solve(l: list) -> int:
    rows = [l]
    while not all([x == 0 for x in rows[-1]]):
        rows.append(diff(rows[-1]))

    rows[-1].append(0)
    for i in range(len(rows) - 2, -1, -1):  # omit the last row, going backward
        above = rows[i + 1]
        s = above[-1] + rows[i][-1]
        rows[i].append(s)
    return rows[0][-1]


def diff(l: list) -> list:
    return [l[i] - l[i - 1] for i in range(1, len(l))]


def clean_line(line: str):
    return list(map(int, [x for x in line.split(" ") if x]))


if __name__ == "__main__":
    do()
