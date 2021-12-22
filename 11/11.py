with open('input.txt') as input:
    lines = input.readlines()

input = []

for line in lines:
    parsed = map(int, list(line.strip()))
    parsed = list(map(lambda x: (x, True), parsed))
    input.append(parsed)

#---------------------------------
print('part 2 ---------------')

def printPretty(input):
    for row in range(len(input)):
        print(" ".join(str(e[0]) for e in input[row]))

    print('')


# for each step, we iterate through entire input matrix
# increment each value by 1
#    if > 9 -> flash and increment all adjacent by 1
#    keep flashing and incrementing 
#    can flash only once per iteration

def flashInput(input, row1, col1, flashes):
    queue = [(row1, col1)]
    while(len(queue) > 0):
        row = queue[0][0]
        col = queue[0][1]
        queue.pop(0)

        
        # if it is True, there is no other work to do
        if (input[row][col][1]):
            continue

        input[row][col] = (input[row][col][0] + 1, input[row][col][1])

        if (input[row][col][0] > 9):
            input[row][col] = (0, True)
            flashes += 1

            # left col
            if (col != 0):
                if (row != 0):
                    queue.append((row - 1, col - 1))
                if (row != len(input) - 1):
                    queue.append((row + 1,col - 1))
                queue.append((row,col - 1))
            
            # middle col
            if (row != 0):
                queue.append((row - 1,col))

            if (row != len(input) - 1):
                queue.append((row + 1,col))

            # right col
            if (col != len(input[0]) - 1):
                if (row != 0):
                    queue.append((row - 1, col + 1))
                if (row != len(input) - 1):
                    queue.append((row + 1, col + 1))
                queue.append((row,col + 1))
        
    return flashes

flashes = 0

def setToFalse(input):
    for row in range(len(input)):    
        for col in range(len(input[0])):
            input[row][col] = (input[row][col][0], False)

def isWholeBoardZero(input):

    for row in range(len(input)):    
        for col in range(len(input[0])):
            if (input[row][col][0] is 0):
                continue
            else:
                return False

    return True

for step in range(10000):
    setToFalse(input)
    for row in range(len(input)):    
        for col in range(len(input[0])):
            if (input[row][col][1] is False):
                input[row][col] = (input[row][col][0] + 1, input[row][col][1])

            if (input[row][col][0] > 9):
                flashes = flashInput(input, row, col, flashes)
    if (isWholeBoardZero(input)):
            print(step)
            break
