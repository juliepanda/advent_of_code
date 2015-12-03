def bubble_sort(arr):
    for i in xrange(len(arr)):
        for j in xrange(i+1, len(arr)):
            if arr[i] > arr[j]:
                arr[i], arr[j] = arr[j], arr[i]
    return arr

with open('input2.txt', 'r') as f:
    words = [bubble_sort([int(x) for x in line.split('x') if len(x) > 0]) for line in f.read().split('\n')]
    total_ribbon_needed, total_wrap_needed = 0, 0
    for word in words:
        if len(word) < 3:
            break
        s, m, l = word[0], word[1], word[2]
        total_ribbon_needed += 2*s + 2*m + s*m*l
        total_wrap_needed += (2*s*m + 2*m*l + 2*l*s) + s
    print total_wrap_needed, total_ribbon_needed
