from unittest import TestCase
from unittest import main
from .count_symmetric_integers import Solution


class TestSolution(TestCase):
    def test_count_symmetric_integers(self):
        solution = Solution()
        self.assertEqual(9, solution.countSymmetricIntegers(1, 100))
        self.assertEqual(4, solution.countSymmetricIntegers(1200, 1230))


if __name__ == '__main__':
    main()
