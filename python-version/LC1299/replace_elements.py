class Solution:
    def replaceElements(self, arr: list[int]) -> list[int]:
        n = len(arr)
        ans = [0] * n
        mx = -1
        for i in range(n-1, -1, -1):
            ans[i] = mx
            mx = max(mx, arr[i])
        return ans