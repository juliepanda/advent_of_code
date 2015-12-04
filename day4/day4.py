import hashlib

# part 1
# for i in xrange(0, 1000000):
#     m = hashlib.md5()
#     hashstr = "ckczppom" + str(i)
#     m.update(hashstr)
#     if m.hexdigest()[0:5] == "00000":
#         print m.hexdigest()
#         print hashstr
#         break


for i in xrange(0, 10000000):
    m = hashlib.md5()
    hashstr = "ckczppom" + str(i)
    m.update(hashstr)
    if m.hexdigest()[0:6] == "000000":
        print m.hexdigest()
        print hashstr
        break
