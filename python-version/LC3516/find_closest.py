class Solution:
    @staticmethod
    def findClosest(x: int, y: int, z: int) -> int:
        a, b = abs(x - z), abs(y - z)
        if a == b:
            return 0
        return 1 if a < b else 2
