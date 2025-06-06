class Solution:
    @staticmethod
    def robotWithString(s: str) -> str:
        n = len(s)
        suf_min = ['z'] * (n+1)
        for i in range(n-1, -1, -1):
            suf_min[i] = min(suf_min[i+1], s[i])

        ans = ''
        st = suf_min[:0]
        for i, ch in enumerate(s):
            st.append(ch)
            while st and st[-1] <= suf_min[i+1]:
                ans += st.pop()
        return ans
