import unittest
from browser_history import BrowserHistory

class BrowserHistoryTest(unittest.TestCase):
    def test_browserHistory(self):
        bh = BrowserHistory("leetcode.com")
        bh.visit("google.com")
        bh.visit("facebook.com")
        bh.visit("youtube.com")
        print(bh.back(1))
        print(bh.back(1))
        print(bh.forward(1))
        bh.visit("linkedin.com")
        print(bh.forward(2))
        print(bh.back(2))
        print(bh.back(7))

if __name__ == '__main__':
    unittest.main()