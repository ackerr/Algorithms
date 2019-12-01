from typing import List


class Solution:
    def four_sum(self, nums: List[int], target: int = 0) -> List[List[int]]:

        nums = sorted(nums)
        res = []

        for i in range(len(nums) - 3):

            if i > 0 and nums[i] == nums[i - 1]:  # 移除首位相同的项
                continue

            if sum(nums[:3], nums[i]) > target:
                break

            if sum(nums[-3:], nums[i]) < target:
                continue

            for j in range(i + 1, len(nums) - 2):

                if j - i > 1 and nums[j] == nums[j - 1]:  # 移除第二位相同的项
                    continue

                if sum(nums[j : j + 3], nums[i]) > target:
                    break

                if sum([nums[i], nums[j], nums[-1], nums[-2]]) < target:
                    continue

                left, right = j + 1, len(nums) - 1

                while left < right:
                    s = nums[i] + nums[j] + nums[left] + nums[right]

                    if s < target:
                        left += 1
                    elif s > target:
                        right -= 1
                    else:
                        res.append([nums[i], nums[j], nums[left], nums[right]])
                        left += 1
                        right -= 1

                        while left < right and nums[left] == nums[left - 1]:
                            left += 1
                        while left < right and nums[right] == nums[right + 1]:
                            right -= 1
        return res


if __name__ == "__main__":
    assert Solution().four_sum([0, 0, 0, 0]) == [[0, 0, 0, 0]]
