inputs = []
with open("../input.txt", "r") as file:
    for line in file:
        inputs.append(line)


def count_chars(string):
    has_2 = 0
    has_3 = 0

    seen = {}
    for c in string:
        if c not in seen:
            seen[c] = 1
        else:
            seen[c] += 1

    if 2 in seen.values():
        has_2 = 1

    if 3 in seen.values():
        has_3 = 1

    return has_2, has_3

def checksum(inputs):
    count_2 = 0
    count_3 = 0

    for i in inputs:
        a, b = count_chars(i)
        count_2 += a
        count_3 += b

    return count_2 * count_3

print(checksum(inputs))
