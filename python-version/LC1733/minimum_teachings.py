from typing import List


class Solution:
    @staticmethod
    def minimumTeachings(n: int, languages: List[List[int]], friendships: List[List[int]]) -> int:
        not_communicate = set()
        for _, friendship in enumerate(friendships):
            st = set()
            flag = False
            for lan in languages[friendship[0] - 1]:
                st.add(lan)
            for lan in languages[friendship[1] - 1]:
                if lan in st:
                    flag = True
                    break
            if not flag:
                not_communicate.add(friendship[0] - 1)
                not_communicate.add(friendship[1] - 1)
        max_cnt = 0
        cnt = [0] * (n+1)
        for friend in not_communicate:
            for lan in languages[friend]:
                cnt[lan] += 1
                max_cnt = max(max_cnt, cnt[lan])
        return len(not_communicate) - max_cnt
