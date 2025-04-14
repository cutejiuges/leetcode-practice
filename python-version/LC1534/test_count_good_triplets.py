from unittest import TestCase
from unittest import main
from .count_good_triplets import Solution


class TestSolution(TestCase):
    def test_count_good_triplets(self):
        solution = Solution()
        self.assertEqual(4, solution.countGoodTriplets(arr=[3, 0, 1, 1, 9, 7], a=7, b=2, c=3))
        self.assertEqual(0, solution.countGoodTriplets(arr=[1, 1, 2, 2, 3], a=0, b=0, c=1))


if __name__ == '__main__':
    main()
