class Solution:
    @staticmethod
    def minMaxDifference(num: int) -> int:
        num_str = str(num)
        mx, mn = num, num
        for ch in num_str:
            if ch != '9':
                mx = int(num_str.replace(ch, '9'))
                break
        mn = int(num_str.replace(num_str[0], '0'))
        return mx - mn
