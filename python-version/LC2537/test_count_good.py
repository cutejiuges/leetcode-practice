from unittest import TestCase
from unittest import main
from .count_good import Solution


class TestSolution(TestCase):
    def test_count_good(self):
        solution = Solution()
        self.assertEqual(1, solution.countGood(nums=[1, 1, 1, 1, 1], k=10))
        self.assertEqual(4, solution.countGood(nums=[3, 1, 4, 3, 2, 2, 4], k=2))


if __name__ == '__main__':
    main()
