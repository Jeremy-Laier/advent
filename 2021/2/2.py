# read file
with open('input.txt') as input:
    lines = input.readlines()

commands = []
for line in lines:
    command = line.strip().split(' ')
    command[1] = int(command[1])
    commands.append(command)

# solve
horizontal = 0
depth = 0
aim = 0

def parseCommand(command):
    global horizontal, aim, depth

    cmd = command[0]
    change = command[1]

    if (cmd == 'forward'):
        horizontal += change
        depth += change * aim
    elif(cmd == 'down'):
        aim += change
    else:
        aim -= change
    print(horizontal, depth, aim)

for command in commands:
    parseCommand(command)

print(horizontal, depth, aim, horizontal * depth)
