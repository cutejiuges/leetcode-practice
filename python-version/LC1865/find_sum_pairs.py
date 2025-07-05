class FindSumPairs:
    def __init__(self, nums1: list[int], nums2: list[int]):
        self.__nums1 = nums1
        self.__nums2 = nums2
        self.__mp = dict()
        for num in nums2:
            self.__mp[num] = self.__mp.get(num, 0) + 1

    def add(self, index: int, val: int) -> None:
        old = self.__nums2[index]
        self.__mp[old] = self.__mp.get(old, 0) - 1
        self.__nums2[index] += val
        self.__mp[old + val] = self.__mp.get(old + val, 0) + 1

    def count(self, tot: int) -> int:
        ans = 0
        for num in self.__nums1:
            ans += self.__mp.get(tot - num, 0)
        return ans
