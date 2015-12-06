# vowel = {
#     "a": True,
#     "e": True,
#     "i": True,
#     "o": True,
#     "u": True
# }

# bad_strings = {
#     "ab": True,
#     "cd": True,
#     "pq": True,
#     "xy": True
# }


nice_words = 0
with open('input5.txt', 'r') as f:
    words = f.read().split('\n')
    for word in words:
        # vowel_count = 0
        prev = ''
        prev2 = ''
        # consec = False
        # is_bad = False
        match3 = False
        seen = {}
        seen_char = False
        for i, c in enumerate(word):
            # if bad_strings.get(prev+c, False):
            #     is_bad = True
            #     break
            if len(prev) > 0:
                a = seen.get(prev+c, -1)
                if a != i-1:
                    if a != -1:
                        seen_char = True
                    else:
                        seen[prev + c] = i
            # if prev == c:
            #     consec = True
            # if vowel.get(c, False):
            #     vowel_count += 1
            if prev2 == c:
                match3 = True
            prev2 = prev
            prev = c
            # if vowel_count >= 3 and not is_bad and consec and match3 and seen_char:
            if seen_char and match3:
                nice_words += 1
                break
print nice_words
