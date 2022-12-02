with open('testInput.txt') as input:
    lines = input.readlines()

given = []
for line in lines:
    inpu = line.strip().split('|')
    sorte = inpu[0].strip().split(' ')
    sorte.sort(key = lambda x: len(x))
    given.append(sorte + inpu[1].strip().split(' '))

# 2 segments = 1
# 4 segments = 4
# 3 segments = 7
# 7 segments = 8
nums = {
        2: 1,
        4: 4,
        3: 7,
        7: 8
        }

count = 0
for signals in given:
    for output in signals[-4:]:
        if (nums.get(len(output), -1) != -1):
            count+=1
            
print(count)

print("part 2 ---------------")

#        0:      1:      2:      3:      4:
#         aaaa    ....    aaaa    aaaa    ....
#        b    c  .    c  .    c  .    c  b    c
#        b    c  .    c  .    c  .    c  b    c
#         ....    ....    dddd    dddd    dddd
#        e    f  .    f  e    .  .    f  .    f
#        e    f  .    f  e    .  .    f  .    f
#         gggg    ....    gggg    gggg    ....

#
#       5:      6:       7:      8:      9:
#         aaaa    aaaa    aaaa    aaaa    aaaa
#        b    .  b    .  .    c  b    c  b    c
#        b    .  b    .  .    c  b    c  b    c
#         dddd    dddd    ....    dddd    dddd
#        .    f  e    f  .    f  e    f  .    f
#        .    f  e    f  .    f  e    f  .    f
#         gggg    gggg    ....    gggg    gggg

# can only use set(7), set(1), set(4), set(8)
# a -> set(7) - set(1)
# g -> set(3) intersection set(2) intersection set(5) - a4 - a
# d -> set(3) intersection set(2) intersection set(5) - a - g
# b -> set(4) - set(1) - d 
# e -> set(8) - set(4) - a - g
# set(5) -> if b in number and e not in number and len(num) == 5
# f -> set(5) - b - a - g
# c -> set(1) - f

# 0 => len(6) and contains a, b, c, e, f, g
# 1 => already know
# 2 => len(5) and contains a, c, d, e, g
# 3 => len(5) and contains a, c, d, f, g
# 4 => already know
# 5 => len(5) and contains a, b, d, f, g
# 6 => len(6) and contains a, b, d, e, f, g
# 7 => already know
# 8 => already know
# 9 => len(6) and contains a, b, c, d, f, g

answers = []
for li in given:
    # numbers to use to decode
    nums = li[:10]
    nums.sort(key = lambda x: len(x))
    a1 = set(nums[0])
    a7 = set(nums[1])
    a4 = set(nums[2])
    a8 = set(nums[9])

    a = a7 - a1
    g = set(nums[3]).intersection(set(nums[4])).intersection(set(nums[5])) - a4 - a
    d = set(nums[3]).intersection(set(nums[4])).intersection(set(nums[5])) - a - g
    b = a4 - a1 - d
    e = a8 - a4 - a - g
    
    a5 = set()
    if (  'b' in nums[3]): a5 = set(nums[3])
    elif ('b' in nums[4]): a5 = set(nums[4])
    elif ('b' in nums[5]): a5 = set(nums[5])
    else: raise Exception('wtf why didnt this work')
    
    f = a5 - b - a - g - d
    c = a1 - f

    answer = []
    for output in li[-4:]:
        if (set(output) == a1):
            answer.append(1)
        elif(set(output) == a7):
            answer.append(7)
        elif(set(output) == a4):
            answer.append(4)
        elif(set(output) == a8):
            answer.append(8)
        elif(set(output) == a | b | c | e | f | g):
            answer.append(0)
        elif(set(output) == a | c | d | e | g):
            answer.append(2)
        elif(set(output) == a | c | d | f | g):
            answer.append(3)
        elif(set(output) == a5):
            answer.append(5)
        elif(set(output) == a | b | d | e | f | g):
            answer.append(6)
        elif(set(output) == a | b | c | d | f | g):
            answer.append(9)
        else:
            print('a', a, 'b', b, 'c', c, 'd', d, 'e', e, 'f', f, 'g', g)
            raise Exception(output, "cant parse this shit")

    print(li[-4:])
    answers.append(''.join(map(str, answer)))

answers = list(map(int, answers))
print(sum(answers))
