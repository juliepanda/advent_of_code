with open('input1.txt', 'r') as f:
    words = f.read()

count = 0
if len(words) > 0:
    for i, c in enumerate(words):
        # if count == -1:
        #     # part 2
        #     print i
        #     break
        if c == '(':
            count += 1
        elif c == ')':
            count -= 1
        else:
            break
    print count
