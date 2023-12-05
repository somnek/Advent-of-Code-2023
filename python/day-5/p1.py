items = [
    "seed",
    "soil",
    "fertilizer",
    "water",
    "light",
    "temperature",
    "humidity",
    "location",
]


def do():
    with open("input.txt", "r") as f:
        lines = f.readlines()
        lines = [l.strip() for l in lines]

        init_seeds = list(map(int, lines[0].split(": ")[1].strip().split(" ")))
        lines = lines[2:]

        mm = {v: 0 for v in items}
        locs = []
        name = ""

        for si, seed in enumerate(init_seeds):

            ii = 1  # item index
            mm["seed"] = seed
            last = mm[items[ii - 1]]

            for line in lines:

                if line.endswith(":"):
                    name = items[ii]
                    ii += 1
                    done = False

                elif line and not done:
                    nums = [int(d) for d in line.split()]
                    dst, src, ran = nums

                    if src <= last <= src + ran:
                        rs = dst + (last - src)
                        mm[name] = rs
                        done = True
                        last = rs
                    else:
                        mm[name] = last

            locs.append(last)
            ii = 0

    print(min(locs))

    
if __name__ == "__main__":
    do()
