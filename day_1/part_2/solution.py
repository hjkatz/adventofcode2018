import sys

solution = 0
series = {}

frequencies = []
with open("../input.txt", "r") as file:
    for line in file:
        frequencies.append(int(line))

while True:
    for reading in frequencies:
        if solution in series:
            print(solution) # found duplicate
            sys.exit()
        else:
            series[solution] = True

        solution += reading
