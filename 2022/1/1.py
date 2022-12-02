with open('input.txt') as input:
        lines = input.readlines()

calories = []
curCalorie = 0
for i, calorie in enumerate(lines):
    if calorie == "\n":
        calories.append(curCalorie)
        curCalorie = 0
    else:
        curCalorie += int(calorie)

print("part 1 answer", sorted(calories)[-1:][0])

print("part 2 answer", sum(sorted(calories)[-3:]))




