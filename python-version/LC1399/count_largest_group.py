class Solution:
    def countLargestGroup(self, n: int) -> int:
        mp = dict()
        max_size, ans = 0, 0
        for i in range(1, n+1):
            sm = self.__get_digit_sum(i)
            cur_list = mp.get(sm, [])
            cur_list.append(i)
            mp[sm] = cur_list
            if len(cur_list) > max_size:
                max_size = len(cur_list)
                ans = 1
            elif len(cur_list) == max_size:
                ans += 1
        return ans

    @staticmethod
    def __get_digit_sum(x: int) -> int:
        ans, t = 0, x
        while t > 0:
            ans += t % 10
            t //= 10
        return ans
