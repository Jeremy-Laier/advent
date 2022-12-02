with open('input.txt') as input:
    lines = input.readlines()

# graph 
# start -> A, B
# A -> c, b, end
# etc
caves = {}

for line in lines:
    pat = line.split('-')
    fro= pat[0].strip()
    to = pat[1].strip()
    caves.setdefault(fro, []).append(to)
    caves.setdefault(to, []).append(fro)

# goal: find number of paths from start to end that visit small caves at most once

# sorting is probably over complicating it
# for cave in caves:
#    caves[cave].sort()

from copy import deepcopy
def areTwoSmallsVisited(curPath):
    path = list(filter(lambda x: x != 'start', curPath))
    path = list(filter(lambda x: curPath.count(x) > 1, path))
    path = list(filter(lambda x: not x.isupper(), path))
    return len(set(path)) > 0

def dfs(caves, currentCave, path, allPaths):
    curPath = deepcopy(path)

    if (currentCave == 'end'):
        if (currentCave not in curPath):
            curPath.append(currentCave)

        allPaths.append(curPath)
        return

    # if cave is in the path and 
    if (currentCave in curPath and not currentCave.isupper() and areTwoSmallsVisited(curPath)):
        return
    else:
        curPath.append(currentCave)

        for cave in caves[currentCave]:
            if (cave == 'start'):
                continue
            else:
                dfs(caves, cave, curPath, allPaths)
            
path = []
paths = []
dfs(caves, 'start', path, paths)

def printPretty(paths):
    for path in paths:
        print('')
        print(', '.join(path))

print(len(paths))
