with open('input.txt') as input:
    lines = input.readlines()

chunks = []
for line in lines:
    chunks.append(list(line.strip()))

validItems = {
        '[': True,
        '<': True,
        '{': True,
        '(': True
        }

wrong = []
proper = []
wut = []
for i  in range(len(chunks)):
    stack = []
    for j,val in enumerate(chunks[i]):

        if (len(stack) > 0):
            if (stack[-1] is '{' and val is '}'):
                stack.pop()
                if (j is len(chunks[i]) - 1):
                    proper.append(chunks[i])
                    wut.append(stack)
                continue

            if (stack[-1] is '(' and val is ')'):
                stack.pop()
                if (j is len(chunks[i]) - 1):
                    proper.append(chunks[i])
                    wut.append(stack)
                continue

            if (stack[-1] is '<' and val is '>'):
                stack.pop()
                if (j is len(chunks[i]) - 1):
                    proper.append(chunks[i])
                    wut.append(stack)
                continue

            if (stack[-1] is '[' and val is ']'):
                if (j is len(chunks[i]) - 1):
                    proper.append(chunks[i])
                    wut.append(stack)
                stack.pop()
                continue

        if (validItems.get(val)):
            stack.append(val)
            if (j is len(chunks[i]) - 1):
                proper.append(chunks[i])
                wut.append(stack)
            continue
        else:
            wrong.append(val)
            break


scores = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137
        }

score = map(lambda x: scores.get(x), wrong)
print(sum(score))

print('part 2-----------')
fixing = []
fixed = {
        '{': 3,
        '(': 1,
        '[':2,
        '<':4
        }

wut.reverse()
for i in wut:
    p = []
    for j in i:
        p.append(fixed.get(j))
    p.reverse()
    fixing.append(p)


points = []
for fixed in fixing:
    pts = 0
    for i, val in enumerate(fixed):
        pts *= 5
        pts += val
    points.append(pts)

points.sort()

print(points)
print(points[int(len(points) / 2)])
