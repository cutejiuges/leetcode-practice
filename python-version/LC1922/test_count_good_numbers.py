from unittest import TestCase
from unittest import main
from .count_good_numbers import Solution


class TestSolution(TestCase):
    def test_count_good_numbers(self):
        solution = Solution()
        self.assertEqual(5, solution.countGoodNumbers(n=1))
        self.assertEqual(400, solution.countGoodNumbers(n=4))
        self.assertEqual(564908303, solution.countGoodNumbers(n=50))


if __name__ == '__main__':
    main()
