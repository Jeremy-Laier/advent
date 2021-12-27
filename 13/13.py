with open('input.txt') as input:
    lines = input.readlines()

dots = []
folds = []
for line in lines:
    parsed = str(line.strip())
    # dont parse empty line
    if ('' == parsed):
        continue

    # fold goes into fold list
    if ('fold' in parsed):
        parsed = parsed.split('=')
        if ('x' in parsed[0]):
            folds.append(('x', parsed[1]))
        else:
            folds.append(('y', parsed[1]))

    # dots go into dots list
    else:
        parsed = parsed.split(',')
        parsed = list(map(int, parsed))
        parsed.reverse()
        dots.append(parsed)

def buildPaper(dots):
    maxX, maxY, paper = 0, 0, []

    for dot in dots:
        if (dot[0] > maxX):
            maxX = dot[0]

        if (dot[1] > maxY):
            maxY = dot[1]

    for row in range(maxX + 1):
        paper.append([])
        for _ in range(maxY + 1):
            paper[row].append(0)
        
    return paper

def insertDots(paper, dots):
    for dot in dots:
        paper[dot[0]][dot[1]] = 1
            

def printPretty(paper):
    for row in range(len(paper)):
        for col in range(len(paper[0])):
            if (paper[row][col] == 0):
                print('.', end = '')
            else:
                print('#', end = '')
        print('')

overlappingDots = 0

def foldUp(paper, y):
    global overlappingDots
    for row in range(y + 1, len(paper)):
        for col in range(len(paper[0])):
            if (paper[row][col] == 1):
                if (paper[len(paper) - row - 1][col] == 1):
                    overlappingDots += 1

                paper[len(paper) - row - 1][col] = 1

    paper  = paper[:y]
    return paper

def foldLeft(paper, x):
    global overlappingDots
    for row in range(len(paper)):
        for col in range(x + 1, len(paper[0])):
            if (paper[row][col] == 1):
                if (paper[row][len(paper[0]) - col - 1] == 1):
                    overlappingDots ++ 1

                paper[row][len(paper[0]) - col - 1] = 1
    for row in range(len(paper)):
        paper[row] = paper[row][:x]

    return paper

def visibleDots(paper):
    sum = 0
    for row in range(len(paper)):
        for col in range(len(paper[0])):
            if (paper[row][col] == 1):
                sum += 1
    return sum

def foldPaper(paper, folds):
    for fold in folds:
        # printPretty(paper)
        if (fold[0] == 'x'):
            paper = foldLeft(paper, int(fold[1]))
        else:
            paper = foldUp(paper, int(fold[1]))
            
    # printPretty(paper)
    return paper

paper = buildPaper(dots)
insertDots(paper, dots)
paper = foldPaper(paper, folds)

print(visibleDots(paper) + overlappingDots)
printPretty(paper)
