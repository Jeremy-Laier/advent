#!/usr/bin/python
# -*- coding: utf-8 -*-
 # read file

with open('input.txt') as input:
    lines = input.readlines()

matrix = []
for line in lines:
    row = list(line.strip())
    row = map(lambda ele: int(ele), row)
    matrix.append(row)

gamma = ''
epsilon = ''

for column in range(len(matrix[0])):
    zeroes = 0
    ones = 0
    for row in range(len(matrix)):
        if matrix[row][column] == 0:
            zeroes += 1
        else:
            ones += 1
    if zeroes > ones:
        epsilon += '1'
        gamma += '0'
    else:
        epsilon += '0'
        gamma += '1'

print ('gamma ', gamma, 'epsilon ', epsilon)
print int(gamma, 2) * int(epsilon, 2)

print '----------------------------- part 2 ------------------------------'


# return most common bit in column

def commonBit(matrix, position):
    zeroes = 0
    ones = 0
    for row in range(len(matrix)):
        if matrix[row][position] == 0:
            zeroes += 1
        else:
            ones += 1

    if zeroes > ones:
        return 0
    else:
        return 1


def leastCommonBit(matrix, position):
    zeroes = 0
    ones = 0
    for row in range(len(matrix)):
        if matrix[row][position] == 0:
            zeroes += 1
        else:
            ones += 1

    if zeroes <= ones:
        return 0
    else:
        return 1


def removeFromMatrix(matrix, commonBit, position):
    newMatrix = []
    for row in range(len(matrix)):
        if matrix[row][position] == commonBit:
            newMatrix.append(matrix[row])
    return newMatrix


def rating(matrix, rate):
    answer = matrix
    for col in range(len(matrix[0])):
        if len(answer) == 1:
            break

        bit = 0
        if rate == 1:
            bit = commonBit(answer, col)
        else:
            bit = leastCommonBit(answer, col)

        answer = removeFromMatrix(answer, bit, col)

    return answer


oxy = rating(matrix, 1)[0]
c02 = rating(matrix, 0)[0]

oxy = ''.join(str(e) for e in oxy)
c02 = ''.join(str(e) for e in c02)

print int(oxy, 2) * int(c02, 2)
