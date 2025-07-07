import heapq


class Solution:
    @staticmethod
    def maxEvents(events: list[list[int]]) -> int:
        max_day = 0
        for event in events:
            max_day = max(max_day, event[1])

        events.sort(key=lambda x: x[0])
        pq = []
        n = len(events)
        ans, j = 0, 0
        for i in range(1, max_day + 1):
            while j < n and events[j][0] <= i:
                heapq.heappush(pq, events[j][1])
                j += 1
            while pq and pq[0] < i:
                heapq.heappop(pq)
            if pq:
                heapq.heappop(pq)
                ans += 1
        return ans
