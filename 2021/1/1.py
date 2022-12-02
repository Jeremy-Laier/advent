with open('input.txt') as input:
    lines = input.readlines()

depths=[]
for line in lines:
    depths.append(int(line.strip()))

last = 999999

increases = 0

for i, depth in enumerate(depths):
    total = sum(depths[i:i+3])
    if (total > last):
        increases += 1
    last = total

print(increases)
