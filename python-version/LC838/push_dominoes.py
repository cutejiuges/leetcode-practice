class Solution:
    @staticmethod
    def pushDominoes(dominoes: str) -> str:
        """
        1. 在原始数组左侧拼一个L，右侧拼一个R方便处理且不影响结果
        2. 原始的L、R不会发生变化
        3. 两个L之间的全部变成L，两个R之间的全部变成R
        4. 左L和右R之间的全部不变，左R和右L之间的向中间靠
        :param dominoes: 原始摆放的骨牌
        :return: 经过推倒变换最终的结果
        """

        dominoes = "L" + dominoes + "R"
        left, right = 0, 1
        ans = ""
        while right < len(dominoes):
            while right < len(dominoes) and dominoes[right] == '.':
                right += 1
            left_char, right_char = dominoes[left], dominoes[right]
            if left_char == right_char:
                for i in range(left+1, right):
                    ans += left_char
            elif left_char == 'R':
                half_len = (right - left - 1) >> 1
                for i in range(half_len):
                    ans += 'R'
                if (right - left) & 1 == 0:
                    ans += '.'
                for i in range(half_len):
                    ans += 'L'
            else:
                for i in range(left+1, right):
                    ans += '.'
            ans += right_char
            left = right
            right += 1
        return ans[:-1]
