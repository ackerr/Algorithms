class Node:
    def __init__(self, val):
        self.val = val
        self.next = None


class MyLinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def get(self, index: int) -> int:
        if index > self.size or index < 0 or not self.head:
            return -1
        cur = self.head
        for _ in range(index):
            cur = cur.next
        return cur.val

    def addAtHead(self, val: int) -> None:
        n = Node(val)
        n.next = self.head
        self.head = n
        self.size += 1

    def addAtTail(self, val: int) -> None:
        n = Node(val)
        if self.head is None:
            self.head = n
        else:
            cur = self.head
            while cur.next:
                cur = cur.next
            cur.next = n
        self.size += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index > self.size or index < 0:
            return
        if index == 0:
            self.addAtHead(val)
        else:
            cur = self.head
            for _ in range(index - 1):
                cur = cur.next
            n = Node(val)
            n.next = cur.next
            cur.next = n
            self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        if index >= self.size or index < 0:
            return
        cur = self.head
        if index == 0:
            self.head = cur.next
        else:
            for _ in range(index - 1):
                cur = cur.next
            cur.next = cur.next.next
        self.size -= 1
