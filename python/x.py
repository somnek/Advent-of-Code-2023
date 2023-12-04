def do():
    with open('input.txt', 'r') as f:
        lines = f.readlines()

    ttl = 0
    for i, line in enumerate(lines):
        win_part, me_part = get_part(line)
        wc = len(set(win_part) & set(me_part)) # win count

        if wc > 0:
            ttl += 2 ** (wc - 1)

    print(f'ttl = {sum}')


def get_part(line):
    part_a, part_b = line.split('| ')
    win_part_raw = part_a.split(': ')[-1]
    win_part = [int(x) for x in win_part_raw.split(' ') if x]
    me_part = [int(x) for x in part_b.split(' ') if x]
    return win_part, me_part


if __name__ == '__main__':
    do()
