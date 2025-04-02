from unittest import TestCase
from unittest import main
from .maximum_triplet_value import Solution


class TestSolution(TestCase):
    def test_maximum_triplet_value(self):
        solution = Solution()
        self.assertEqual(77, solution.maximumTripletValue(nums=[12, 6, 1, 2, 7]))
        self.assertEqual(133, solution.maximumTripletValue(nums=[1, 10, 3, 4, 19]))
        self.assertEqual(0, solution.maximumTripletValue(nums=[1, 2, 3]))


if __name__ == '__main__':
    main()
