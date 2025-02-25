import unittest
from allocator import Allocator

class AllocatorTest(unittest.TestCase):
    def test_allocator(self):
        allocator = Allocator(10)
        print(allocator.allocate(1, 1))
        print(allocator.allocate(1, 2))
        print(allocator.allocate(1, 3))
        print(allocator.freeMemory(2))
        print(allocator.allocate(3, 4))
        print(allocator.allocate(1, 1))
        print(allocator.allocate(1, 1))
        print(allocator.freeMemory(1))
        print(allocator.allocate(10, 2))
        print(allocator.freeMemory(7))

if __name__ == '__main__':
    unittest.main()