import regex
import collections

Claim = collections.namedtuple("Claim", ["id", "x", "y", "width", "height"])

def parse_claim(line):
    match = regex.match("#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", line)
    if match:
        return Claim(id=match.group(1),
                     x=int(match.group(2)),
                     y=int(match.group(3)),
                     width=int(match.group(4)),
                     height=int(match.group(5)))


inputs = []
with open("../input.txt", "r") as file:
    for line in file:
        inputs.append(parse_claim(line))

FABRIC = [[{} for x in xrange(1000)] for y in xrange(1000)]

def mark(claim):
    for row in FABRIC[claim.y:claim.y + claim.height]:
        for inch in row[claim.x:claim.x + claim.width]:
            inch[claim.id] = True

for claim in inputs:
    mark(claim)

def count_duplicate_claims():
    count = 0
    for row in FABRIC:
        for inch in row:
            if len(inch.keys()) > 1:
                count += 1

    return count

print(count_duplicate_claims())
