class Solution:
    def three_sum_closest(self, nums, target):
        nums = sorted(nums)
        ans = sum(nums[:3])

        for i in range(len(nums) - 2):
            j, k = i + 1, len(nums) - 1

            while j < k:
                s = nums[i] + nums[j] + nums[k]

                if abs(ans - target) > abs(s - target):
                    ans = s

                if s < target:
                    j += 1
                elif s > target:
                    k -= 1
                else:
                    return s

        return ans


if __name__ == "__main__":
    nums = [1, 0, -1, 3, -2, 1]
    assert Solution().three_sum_closest(nums, 2) == 2
