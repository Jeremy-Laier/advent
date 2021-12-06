with open('input.txt') as input:
    lines = input.readlines()

inpu = []
for line in lines:
    inpu = line.strip().split(',')


inpu  = list(map(int, inpu))


def solutionPt1(inpu):
    day = 256
    while (day > 0):
        for i in range(len(inpu)):
            inpu[i] -= 1
            if (inpu[i] < 0):
                inpu[i] = 6
                inpu.appendleft(8)
            
        day -= 1

fishes = 9 * [0]

for fish in inpu:
    fishes[fish] += 1

print(fishes)

# grab num of zeroes left
# set 0th index to 0th
# reduce all indeces by 1
# add num of zeroes back to 8

day = 256

while (day > 0):
    zeroFishes = fishes[0]
    fishes.pop(0)
    fishes[6] += zeroFishes
    fishes.append(zeroFishes)
    day -= 1

print(sum(fishes))
