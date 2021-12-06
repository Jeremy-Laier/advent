with open('input.txt') as input:
    lines = input.readlines()

inpu = []
for line in lines:
    parsed = line.strip().replace(' -> ', ',').split(',')
    parsed = list(map(int, parsed))
    inpu.append(parsed)

def findLargestXY(coords):
    largestX = 0
    largestY = 0
    for coord in coords:
        largestX = max(coord[0], coord[2], largestX)
        largestY = max(coord[1], coord[3], largestY)
    # + 1 because coords are indexed starting from 0-> max inclusive
    return (largestX + 1, largestY + 1)


def printPretty(state):
    xy = findLargestXY(inpu)
    for y in range(xy[1]):
        for x in range(xy[0]):
            if (state[y][x] == 0):
                print('.', end='')
            else:
                print(str(state[y][x]), end='')
        print('')

def initialState():
    state = []
    xy = findLargestXY(inpu)
    for y in range(xy[1]):
        state.append([])
        for _ in range(xy[0]):
            state[y].append(0)
    return state

def drawHorti(state, x1, x2, y):
    for x in range(x1, x2+1):
        state[y][x] += 1
    return state

def drawVerty(state, y1, y2, x):
    for y in range(y1, y2 + 1):
        state[y][x] += 1

    return state


# coords => list of 4 numbers
def addLine(state, coords):
    x1n, y1n = coords[0], coords[1]
    x2n, y2n = coords[2], coords[3]

    x1, x2 = min(x1n, x2n), max(x1n, x2n)
    y1, y2 = min(y1n, y2n), max(y1n, y2n)
    
    if (x1 == x2):
        state = drawVerty(state, y1, y2, x1)
        return state
    elif(y1 == y2):
        state = drawHorti(state, x1, x2, y1)
        return state
    else: 
        return state


def solve(state, coords):
    for coord in coords:
        state = addLine(state, coord)

    dangers = 0
    for y in range(len(state)):
        for x in range(len(state[0])):
            if (state[y][x] > 1):
                dangers += 1
    print("dangers: ", dangers)
    return state


state = initialState()
state = solve(state, inpu)

print("part 2 ------------------")

def drawDiag(state, x1, x2, y1, y2):
    # make first point lower x coord
    if (x1 < x2):
        curPointer = (x1, y1)
        if (y1 < y2):
            accumulator = 1
        else:
            accumulator = -1
    else:
        curPointer = (x2, y2)
        if (y2 < y1):
            accumulator = 1
        else:
            accumulator = -1
    
    for x in range(min(x1, x2), max(x1, x2) + 1):
        state[curPointer[1]][x] += 1
        curPointer = (curPointer[0], curPointer[1] + accumulator)
    
    return state

def addLinePt2(state, coords):
    x1n, y1n = coords[0], coords[1]
    x2n, y2n = coords[2], coords[3]

    x1, x2 = min(x1n, x2n), max(x1n, x2n)
    y1, y2 = min(y1n, y2n), max(y1n, y2n)
    
    if (x1 == x2):
        state = drawVerty(state, y1, y2, x1)
        return state
    elif(y1 == y2):
        state = drawHorti(state, x1, x2, y1)
        return state
    else: 
        state = drawDiag(state, x1n, x2n, y1n, y2n)
        return state

def solvePt2(state, coords):
    for coord in coords:
        state = addLinePt2(state, coord)

    dangers = 0
    for y in range(len(state)):
        for x in range(len(state[0])):
            if (state[y][x] > 1):
                dangers += 1
    print("dangers: ", dangers)
    return state

state = initialState()
state = solvePt2(state, inpu)
