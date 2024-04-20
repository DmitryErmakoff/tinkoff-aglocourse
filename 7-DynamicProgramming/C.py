class Queue:
    def __init__(self):
        self.q1 = []
        self.q2 = []
        self.qm1 = []
        self.qm2 = []

    def push(self, v, i):
        self.q1.append((v, i))
        if not self.qm1:
            self.qm1.append((v, i))
        else:
            if v > self.qm1[-1][0]:
                self.qm1.append((v, i))
            else:
                self.qm1.append((self.qm1[-1][0], self.qm1[-1][1]))

    def pop(self):
        if not self.q2:
            while self.q1:
                self.q2.append(self.q1.pop())
                if not self.qm2:
                    self.qm2.append((self.q2[-1][0], self.q2[-1][1]))
                else:
                    if self.qm2[-1][0] > self.q2[-1][0]:
                        self.qm2.append(self.qm2[-1])
                    else:
                        self.qm2.append((self.q2[-1][0], self.q2[-1][1]))
                self.qm1.pop()

        self.q2.pop()
        self.qm2.pop()

    def min(self):
        if not self.q1:
            return self.qm2[-1]
        if not self.q2:
            return self.qm1[-1]
        if self.qm1[-1][0] < self.qm2[-1][0]:
            return self.qm2[-1]
        else:
            return self.qm1[-1]


if __name__ == "__main__":
    n, k = map(int, input().split())
    a = [0] + list(map(int, input().split())) + [0]

    dp_otv = [None] * n
    dp = Queue()
    dp.push(0, 0)
    dp_otv[0] = (0, -1)

    for i in range(1, min(n, k)):
        dp_otv[i] = (dp.min()[0] + a[i], dp.min()[1])
        dp.push(dp_otv[i][0], i)

    for i in range(k, n):
        dp_otv[i] = (dp.min()[0] + a[i], dp.min()[1])
        dp.pop()
        dp.push(dp_otv[i][0], i)

    print(dp_otv[-1][0])

    i = n - 1
    otv = []
    while i != -1:
        otv.append(i + 1)
        i = dp_otv[i][1]

    print(len(otv) - 1)
    print(*reversed(otv))
