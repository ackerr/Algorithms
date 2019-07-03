class MyStack(object):
    def __init__(self):
        self.ans = []

    def push(self, x):
        self.ans.append(x)

    def pop(self):
        return self.ans.pop()

    def top(self):
        return self.ans[-1]

    def empty(self):
        return self.ans == []
