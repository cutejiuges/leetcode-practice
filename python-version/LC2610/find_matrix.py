class Solution:
    def findMatrix(self, nums: list[int]) -> list[list[int]]:
        cnt = {}
        for num in nums:
            cnt[num] = cnt.get(num, 0) + 1

        res = []
        while len(cnt) > 0:
            cur = list(cnt.keys())
            res.append(cur)
            for _, v in enumerate(cur):
                cnt[v] = cnt.get(v) - 1
                if cnt[v] == 0:
                    cnt.pop(v)

        return res
