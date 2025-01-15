import unittest
import heapq

class Solution:
    def minOperations(self, nums: list[int], k: int) -> int:
        ans = 0
        heap = nums[:]
        heapq.heapify(heap)
        while heap[0] < k:
            x = heapq.heappop(heap)
            y = heapq.heappop(heap)
            heapq.heappush(heap, 2*x+y)
            ans += 1
        return ans
    
class SolutionTest(unittest.TestCase):
    def TestMinOperations(self):
        solution = Solution()
        self.assertEqual(2, solution.minOperations([2,11,10,1,3], 10))
        self.assertEqual(4, solution.minOperations([1,1,2,4,9], 20))


if __name__ == '__main__':
    unittest.main()