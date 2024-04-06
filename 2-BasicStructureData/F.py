import sys
n = int(input())
queue = {}
offset = 0
index = 0
for _ in range(n):
    event = list(map(int, sys.stdin.readline().split()))

    if event[0] == 1:
        queue[event[1]] = offset + index
        index += 1
    elif event[0] == 2:
        first_key = next(iter(queue))
        del queue[first_key]
        index -= 1
        offset += 1
    elif event[0] == 3:
        queue.popitem()
        index -= 1
    elif event[0] == 4:
        print(queue[event[1]] - offset)
    elif event[0] == 5:
        print(next(iter(queue)))