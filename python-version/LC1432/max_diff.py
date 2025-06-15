class Solution:
    @staticmethod
    def maxDiff(num: int) -> int:
        num_str = str(num)
        mx, mn = num, num

        for ch in num_str:
            if ch != '9':
                mx = int(num_str.replace(ch, '9'))
                break

        if num_str[0] != '1':
            mn = int(num_str.replace(num_str[0], '1'))
        else:
            for ch in num_str[1:]:
                if ch >= '2':
                    mn = int(num_str.replace(ch, '0'))
                    break

        return mx - mn
