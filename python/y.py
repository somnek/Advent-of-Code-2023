def do():
    with open('input.txt', 'r') as f:
        lines = f.readlines()

    size = len(lines)
    m = {k:1 for k in range(size)} 

    for i, line in enumerate(lines):
        win_part, me_part = get_part(line)
        wc = len(set(win_part) & set(me_part)) # win count

        for j in range(wc):
            m[i+j+1] += m[i] # plus how many copies

    print(f"ttl = {sum([x for x in m.values()])}")


def get_part(line):
    part_a, part_b = line.split('| ')
    win_part_raw = part_a.split(': ')[-1]
    win_part = [int(x) for x in win_part_raw.split(' ') if x]
    me_part = [int(x) for x in part_b.split(' ') if x]
    return win_part, me_part


if __name__ == '__main__':
    do()
