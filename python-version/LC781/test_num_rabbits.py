from unittest import TestCase
from .num_rabbits import Solution


class TestSolution(TestCase):
    def test_num_rabbits(self):
        solution = Solution()
        self.assertEqual(5, solution.numRabbits(answers=[1, 1, 2]))
        self.assertEqual(11, solution.numRabbits(answers=[10, 10, 10]))
