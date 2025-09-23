from unittest import TestCase
from .compare_version import Solution


class TestSolution(TestCase):
    def test_compare_version(self):
        solution = Solution()
        self.assertEqual(-1, solution.compareVersion(version1="1.2", version2="1.10"))
        self.assertEqual(0, solution.compareVersion(version1="1.01", version2="1.001"))
        self.assertEqual(0, solution.compareVersion(version1="1.0", version2="1.0.0.0"))
