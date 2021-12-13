with open('input.txt') as input:
    lines = input.readlines()

given = []
for line in lines:
    inpu = line.strip().split('|')
    given.append(inpu[0].strip().split(' ') + inpu[1].strip().split(' '))

# 2 segments = 1
# 4 segments = 4
# 3 segments = 7
# 7 segments = 8
nums = {
        2: 1,
        4: 4,
        3: 7,
        7: 8
        }

count = 0
for signals in given:
    for output in signals[-4:]:
        if (nums.get(len(output), -1) != -1):
            count+=1
print(count)
