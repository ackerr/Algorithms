from typing import List


class Solution:
    def three_sum(self, nums: List[int], target: int = 0) -> List[List[int]]:
        nums = sorted(nums)
        ans = []

        for k in range(len(nums) - 2):
            if k > 0 and nums[k] == nums[k - 1]:
                continue

            i, j = k + 1, len(nums) - 1

            while i < j:
                s = nums[k] + nums[i] + nums[j]

                if s < target:
                    i += 1

                    while i < j and nums[i] == nums[i - 1]:
                        i += 1
                elif s > target:
                    j -= 1

                    while i < j and nums[j] == nums[j + 1]:
                        j -= 1
                else:
                    ans.append([nums[k], nums[i], nums[j]])
                    i, j = i + 1, j - 1

                    while i < j and nums[i] == nums[i - 1]:
                        i += 1

                    while i < j and nums[j] == nums[j + 1]:
                        j -= 1

        return ans


if __name__ == "__main__":
    Solution().three_sum([-1, 0, -1, 1, 2, -1, 4])
