class Solution:
    @staticmethod
    def numberOfArrays(differences: list[int], lower: int, upper: int) -> int:
        base = 0
        mx, mn = 0, 0
        for diff in differences:
            base += diff
            mx = max(mx, base)
            mn = min(mn, base)
            if mx - mn > upper - lower:
                return 0
        return 1 + (upper - lower) - (mx - mn)
