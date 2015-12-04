def switch(c, x, y, visited, count):
    if c == '<':
        x[0] -= 1
    elif c == '>':
        x[0] += 1
    elif c == '^':
        y[0] += 1
    elif c == 'v':
        y[0] -= 1
    else:
        return
    key = str(x[0]) + '-' + str(y[0])
    if key not in visited:
        visited[key] = True
        count[0] += 1


with open('input3.txt', 'r') as f:
    words = f.read()
    x, y, rx, ry = [0], [0], [0], [0]
    visited = {}
    visited[str(x[0]) + '-' + str(y[0])] = True
    count = [1]
    for i, c in enumerate(words):
        if i % 2 == 0:
            switch(c, rx, ry, visited, count)
        else:
            switch(c, x, y, visited, count)
    print count[0]
