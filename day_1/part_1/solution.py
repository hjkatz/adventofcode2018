solution = 0
with open("../input.txt", "r") as file:
    for line in file:
        solution += int(line)

print(solution)
