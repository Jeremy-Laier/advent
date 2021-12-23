with open('testInput.txt') as input:
    lines = input.readlines()

input = []
# 2d array of nums
for line in lines:
    parsed = list(line.strip())
    input.append(parsed)

print(input)

def lookupTable(char, stack, corruption):
    if (char is ']'):
        if (stack[-1] is '['):
            stack.pop()
            return
        else:
            corruption.append(stack[-1])
            return

    if (char is ']'):
        if (stack[-1] is '['):
            stack.pop()
            return

    if (stack[-1] is '['):
        stack.pop()
        return
    else:
        corruption.append(stack[-1])
        return



for line in input:
    stack = []

