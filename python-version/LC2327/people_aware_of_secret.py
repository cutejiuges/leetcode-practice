from collections import deque


class Solution:
    @staticmethod
    def peopleAwareOfSecret(n: int, delay: int, forget: int) -> int:
        # 定义两个双端队列，表示仅知道秘密的和能传播秘密的
        know, share = deque([(1, 1)]), deque([])
        # 定义两个整数，表示仅知道秘密的人数和传播秘密的人数
        know_cnt, share_cnt = 1, 0
        # 从第2天开始模拟
        for i in range(2, n+1):
            # 到了能传播秘密的哪一天
            if know and know[0][0] == i - delay:
                know_cnt -= know[0][1]
                share_cnt += know[0][1]
                share.append(know[0])
                know.popleft()
            # 到了忘记秘密的那天
            if share and share[0][0] == i - forget:
                share_cnt -= share[0][1]
                share.popleft()
            # 传播秘密
            if share:
                know_cnt += share_cnt
                know.append((i, share_cnt))
        return (share_cnt + know_cnt) % (10 ** 9 + 7)
