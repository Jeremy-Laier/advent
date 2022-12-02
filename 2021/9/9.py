with open('input.txt') as input:
    lines = input.readlines()

input = []
# 2d array of nums
for line in lines:
    parsed = list(line.strip())
    parsed = list(map(int, parsed))

    input.append(parsed)

# if this is smallest element, save to list
lowPoints = []
for row in range(len(input)):
    for col in range(len(input[0])):
        
        # elements in left col
        # check if 0
        adjacentList = []
        if (col != 0):
            if (row != 0):
                adjacentList.append(input[row - 1][col - 1])
            if (row != len(input) - 1):
                adjacentList.append(input[row + 1][col - 1])
            adjacentList.append(input[row][col - 1])
        
        # check current col
        if (row != 0):
            adjacentList.append(input[row - 1][col])

        if (row != len(input) - 1):
            adjacentList.append(input[row + 1][col])

        # check outer col
        if (col != len(input[0]) - 1):
            if (row != 0):
                adjacentList.append(input[row - 1][col + 1])
            if (row != len(input) - 1):
                adjacentList.append(input[row + 1][col + 1])

            adjacentList.append(input[row][col + 1])

        adjacentList.sort()
        if (input[row][col] <= adjacentList[0]):
            lowPoints.append(input[row][col] + 1)

print(sum(lowPoints))

basins = []
basin  = []

# breadth first search of basins
# make all tiles tuples and mark unseen
# 9 cannot be basin so mark seen

for i in range(len(input)):
    input[i] = list(map(lambda ele: (ele, ele == 9), input[i]))

import os
clear = lambda: os.system('clear')

import time
def prettyPrint(input):
    for row in range(len(input)):
        for e in input[row]:
            if (e[1] is True and not e[0] is 9):
                print(('\x1b[6;30;42m' + str(e[0]) + '\x1b[0m').strip(), end='')
            else:
                print(e[0], end='')
        print('')
    time.sleep(0.01)
    clear()



def bfs(rowI, colI, input, basin):
    queue = [(rowI, colI)]

    while len(queue) != 0:
        row = queue[0][0]
        col = queue[0][1]

        # 9 aint no basin
        if (input[row][col][0] == 9):
            queue.pop(0)
            continue

        # check if seen
        if (input[row][col][1]):
            queue.pop(0)
            continue
        
        # have not seen and the value is not 9
        basin.append(input[row][col][0])
        input[row][col] = (input[row][col][0], True)
        prettyPrint(input)

        if (col != len(input[0]) - 1):
            queue.append((row, col + 1))
        if (row != len(input) - 1):
            queue.append((row + 1, col))
        if (row != 0):
            queue.append((row - 1, col))
        if (col != 0):
            queue.append((row, col - 1))
    
# look for unmarked location
# if found, begin dfs
basins = []
for row in range(len(input)):
    for col in range(len(input[0])):
        if (input[row][col][1] is False):
            basin = []
            bfs(row, col, input, basin)
            basins.append(basin)


basins.sort(key=lambda basin: len(basin))

basins = list(map(lambda basin: len(basin), basins))

# print(basins[-3:])

# import operator, functools

# print(functools.reduce(operator.mul, basins[-3:], 1))

