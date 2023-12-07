

def partOne(filename):
    results = dict()
    with open(filename) as f:
        for line in f:
            p = 0
            cardnum, numbers = line.split(": ")
            winningstr, ownstr = numbers.split(" | ")
            winningset = set(winningstr.split(" "))
            while "" in winningset:
                winningset.remove("")
            ownset = set(ownstr.strip("\n").split(" "))
            while "" in ownset:
                ownset.remove("")
            for n in ownset:
                if n in winningset:
                    if p == 0:
                        p = 1
                    else:
                        p *= 2
            results[cardnum] = p
    return sum(results.values())

def partTwo(filename):
    results = list()
    total = 0
    with open(filename) as f:
        for line in f:
            p = 0
            _, numbers = line.split(": ")
            winningstr, ownstr = numbers.split(" | ")
            winningset = set(winningstr.split(" "))
            while "" in winningset:
                winningset.remove("")
            ownset = set(ownstr.strip("\n").split(" "))
            while "" in ownset:
                ownset.remove("")
            for n in ownset:
                if n in winningset:
                    p+=1
            results.append(p)
    multi = list()
    for _ in range(0, len(results)):
        multi.append(1)
    for i in range(0, len(results)):
        wins = results[i]
        for j in range(i+1, i+1+wins):
            if j >= len(multi):
                continue
            multi[j] += multi[i]

    return sum(multi)


print("Part One Test: " + str(partOne("test.txt")))
print("Part One Ans: " + str(partOne("data.txt")))

print("Part Two Test: " + str(partTwo("test.txt")))
print("Part Two Ans: " + str(partTwo("data.txt")))