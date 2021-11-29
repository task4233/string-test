import csv

runtimes = []
rb = []
mems = []
mb = []
allocs = []
ab = []
with open('log.raw', 'r') as f:
    lines = f.readlines()
    # print(lines)

for idx in range(len(lines)):
    if idx % 10 < 5 or 8 < idx % 10:
        continue

    ss = lines[idx].split()
    if idx % 10 < 8:
        rb.append(ss[2])
        mb.append(ss[4])
        ab.append(ss[6])
        continue
    
    runtimes.append(rb)
    mems.append(mb)
    allocs.append(ab)
    rb = []
    mb = []
    ab = []    

with open('log_runtime.csv', 'w') as f:
    writer = csv.writer(f)
    writer.writerow(['withOp', 'withJoin', 'withBytes', 'withBytesSetCap'])
    for data in runtimes:
        writer.writerow(data)

with open('log_memories.csv', 'w') as f:
    writer = csv.writer(f)
    writer.writerow(['withOp', 'withJoin', 'withBytes', 'withBytesSetCap'])
    for data in mems:
        writer.writerow(data)

with open('log_allocates.csv', 'w') as f:
    writer = csv.writer(f)
    writer.writerow(['withOp', 'withJoin', 'withBytes', 'withBytesSetCap'])
    for data in allocs:
        writer.writerow(data)
