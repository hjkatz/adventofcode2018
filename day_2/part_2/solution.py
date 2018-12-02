import itertools

inputs = []
with open("../input.txt", "r") as file:
    for line in file:
        inputs.append(line)

def string_differ(a, b):
    differ = 0
    common = ""

    i = 0
    while i < len(a):
        if a[i] == b[i]:
            common += a[i]
        else:
            differ += 1

        # don't forget
        i += 1

    return (differ, common)

def find_common_id(inputs):
    for pair in itertools.product(inputs, repeat=2):
        amount, common = string_differ(*pair)

        if amount == 1:
            return common

print(find_common_id(inputs))
