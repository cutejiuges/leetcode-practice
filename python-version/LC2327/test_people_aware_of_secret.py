from unittest import TestCase
from .people_aware_of_secret import Solution


class TestSolution(TestCase):
    def test_people_aware_of_secret(self):
        solution = Solution()
        self.assertEqual(5, solution.peopleAwareOfSecret(n=6, delay=2, forget=4))
        self.assertEqual(6, solution.peopleAwareOfSecret(n=4, delay=1, forget=3))
