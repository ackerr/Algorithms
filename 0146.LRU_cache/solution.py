class DoublyLinkedList:
    """ 双向链表 """

    def __init__(self, key=0, val=0):
        self.pre = None
        self.next = None
        self.key = key
        self.val = val


class LRUCache:
    """ LRU (双向链表结合哈希表)"""

    def __init__(self, capacity: int):
        self.cap = capacity
        self.cache = {}
        self.head = DoublyLinkedList()
        self.tail = DoublyLinkedList()
        self.head.next = self.tail
        self.tail.pre = self.head
        self.size = 0

    def get(self, key: int) -> int:
        if key in self.cache:
            node = self.cache[key]
            self.remove_node(node)
            self.move_to_head(node)
            return node.val
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            node = self.cache[key]
            node.val = value
            self.remove_node(node)
            self.move_to_head(node)
        else:
            node = DoublyLinkedList(key, value)
            self.cache[key] = node
            self.move_to_head(node)
            self.size += 1
            if self.size > self.cap:
                node = self.tail.pre
                self.remove_node(node)
                self.cache.pop(node.key)
                self.size -= 1

    def remove_node(self, node):
        node.pre.next = node.next
        node.next.pre = node.pre

    def move_to_head(self, node):
        node.pre = self.head
        node.next = self.head.next
        self.head.next.pre = node
        self.head.next = node
