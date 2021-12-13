with open('input.txt') as input:
    lines = input.readlines()

inpu = []
for line in lines:
    inpu = line.strip().split(',')

inpu = list(map(int, inpu))

cur = []

for i in range(len(inpu)):
    fuelCost = 0
    for j in range(len(inpu)):
        fuelCost += abs(inpu[i] - inpu[j])

    cur.append( (i, fuelCost))

cur.sort(key = lambda x: x[1])
print(cur[0][1])

print("part2 ----------")
cur = []

for i in range(len(inpu)):
    fuelCost = 0
    cost = 0
    for j in range(len(inpu)):
        cost = abs(inpu[j] - i)
        fuelCost += (cost * (cost + 1)) / 2

    cur.append( (i, fuelCost))


cur.sort(key = lambda x: x[1])
print(cur[0])
