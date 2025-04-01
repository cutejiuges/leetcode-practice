from unittest import TestCase
from unittest import main
from .most_points import Solution


class TestSolution(TestCase):
    def test_most_points(self):
        solution = Solution()
        self.assertEqual(5, solution.mostPoints(questions = [[3,2],[4,3],[4,4],[2,5]]))
        self.assertEqual(7, solution.mostPoints(questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]))


if __name__ == '__main__':
    main()
