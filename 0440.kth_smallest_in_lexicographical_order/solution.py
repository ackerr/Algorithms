class Solution:
    def find_kth_number(self, n: int, k: int) -> int:
        cur = 1
        k -= 1

        while k > 0:
            step = self._step(n, cur, cur + 1)
            # 判断是否k值在子节点中
            if step <= k:
                k -= step
                cur += 1
            else:
                k -= 1
                cur *= 10
        return cur

    def _step(self, n, left, right):
        """ 两值直接的间距 """
        step = 0
        while left <= n:
            step += min(n + 1, right) - left
            left *= 10
            right *= 10
        return step
