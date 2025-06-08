class Solution:
    @staticmethod
    def clearStars(s: str) -> str:
        ss = list(s)
        lists = [[] for _ in range(26)]
        for i, ch in enumerate(s):
            if ch != '*':
                lists[ord(ch) - ord('a')].append(i)
            else:
                for j, li in enumerate(lists):
                    if li:
                        ss[lists[j].pop()] = '*'
                        break

        ans = ""
        for _, ch in enumerate(ss):
            ans += ch if ch != '*' else ''
        return ans
