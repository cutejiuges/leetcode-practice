from unittest import TestCase
from .find_even_numbers import Solution


class TestSolution(TestCase):
    def test_find_even_numbers(self):
        solution = Solution()
        self.assertEqual([102, 120, 130, 132, 210, 230, 302, 310, 312, 320],
                         solution.findEvenNumbers(digits=[2, 1, 3, 0]))
        self.assertEqual([222, 228, 282, 288, 822, 828, 882], solution.findEvenNumbers(digits=[2, 2, 8, 8, 2]))
        self.assertEqual([], solution.findEvenNumbers(digits=[3, 7, 5]))
