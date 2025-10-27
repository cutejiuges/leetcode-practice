from typing import List


class Solution:
    @staticmethod
    def numberOfBeams(bank: List[str]) -> int:
        ans, pre = 0, 0
        for line in bank:
            cnt = 0
            for ch in line:
                cnt += 1 if ch == '1' else 0
            if cnt > 0:
                ans += cnt * pre
                pre = cnt
        return ans
