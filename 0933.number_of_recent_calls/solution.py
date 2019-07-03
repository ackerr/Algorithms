class RecentCounter:
    def __init__(self):
        self.nums = []

    def ping(self, t: int) -> int:
        while self.nums and t - 3000 > self.nums[0]:
            self.nums.pop(0)
        self.nums.append(t)
        return len(self.nums)
