import unittest
from ordered_stream import OrderedStream

class OrderedStreamTest(unittest.TestCase):
    def test_orderedStream(self):
        stream = OrderedStream(5)
        res = []
        res.append(stream.insert(3, "ccccc"))
        res.append(stream.insert(1, "aaaaa"))
        res.append(stream.insert(2, "bbbbb"))
        res.append(stream.insert(5, "eeeee"))
        res.append(stream.insert(4, "ddddd"))

if __name__ == '__main__':
    unittest.main()