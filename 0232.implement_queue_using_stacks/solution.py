class MyQueue:
    def __init__(self):
        self.queue = []

    def push(self, x):
        self.queue.append(x)

    def pop(self):
        if not self.empty():
            return self.queue.pop(0)
        return -1

    def peek(self):
        if not self.empty():
            return self.queue[0]
        return -1

    def empty(self):
        return not bool(len(self.queue))
