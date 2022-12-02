import operator

# read input
with open('input.txt') as input:
    lines = input.readlines()

# read called numbers
numbers = []
for line in lines:
    numbers = line.strip().split(',')
    numbers = map(lambda ele: int(ele), numbers)
    break

# remove called numbers
lines = lines[1:]
boards = []

# clean boards input, turn rows into lists
lines = list(map(lambda ele: ele.strip().split(), lines))

# turn string elements into numbers
rows = []
for line in lines:
    rows.append(map(lambda ele: (int(ele), False), line))

# added cleaned lists into boards, 5x5
# list of 2d matrices
for line in range(0, len(rows), 5):
    boards.append(rows[line: line + 5])

def checkForVictory(board):
    # check rows
    for row in board:
        checked = []
        for num in row:
            if (num[1] is False):
                checked = []
                break
            else:
                checked.append(num)
        if (checked):
            return checked

    # check columns
    for col in range(5):
        columns = []
        for row in range(5):
            if (board[row][col][1] is True):
                columns.append(board[row][col])
            else:
                columns = []
                break
        if (len(columns) is 5):
            return columns

    # no solution yet
    return []

def updateBoard(board, number):
    newBoard = []
    for row in range(5):
        newRow = []
        for col in range(5):
            if (board[row][col][0] is number or board[row][col][1] is True):
                newRow.append((board[row][col][0], True))
            else:
                newRow.append((board[row][col][0], False))
        newBoard.append(newRow)

    return newBoard


def solve(boards, numbers):
    for number in numbers:
        bingo = []
        uppedBoards = boards
        for board in range(len(boards)):
            newBoard = updateBoard(uppedBoards[board], number)
            bingo = checkForVictory(newBoard)
            uppedBoards[board] = newBoard
            if (bingo):
                return (newBoard, number)

def calcAnswer(board, lastCalled):
    summ = 0
    for row in board:
        for ele in row:
            if (ele[1] is False):
                summ += ele[0]
    return summ * lastCalled

def printPretty(board):
    for row in range(5):
        print(" ".join("{:2d} {}".format(e[0], e[1]) for e in board[row]))

solvedTuple = solve(boards, numbers)

print(calcAnswer(solvedTuple[0], solvedTuple[1]))

print("part 2---------")

def solveLast(boards, numbers):
    bingoBoards = []
    for number in numbers:
        uppedBoards = boards
        for board in range(len(boards)):
            newBoard = updateBoard(uppedBoards[board], number)
            bingo = checkForVictory(newBoard)
            uppedBoards[board] = newBoard

            if (bingo):
                # check if all other boards have been solved
                li = list(range(len(boards)))
                li.remove(board)
                if (set(li) == set(bingoBoards)):
                    return (newBoard, number)
                bingoBoards.append(board)

solvedTuple = solveLast(boards, numbers)

print(calcAnswer(solvedTuple[0], solvedTuple[1]))

