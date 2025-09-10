from unittest import TestCase
from .minimum_teachings import Solution


class TestSolution(TestCase):
    def test_minimum_teachings(self):
        solution = Solution()
        self.assertEqual(1, solution.minimumTeachings(n=2, languages=[[1], [2], [1, 2]],
                                                      friendships=[[1, 2], [1, 3], [2, 3]]))
        self.assertEqual(2, solution.minimumTeachings(n=3, languages=[[2], [1, 3], [1, 2], [3]],
                                                      friendships=[[1, 4], [1, 2], [3, 4], [2, 3]]))
