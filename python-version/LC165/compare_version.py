class Solution:
    @staticmethod
    def compareVersion(version1: str, version2: str) -> int:
        v1, v2 = version1.split('.'), version2.split('.')
        len1, len2 = len(v1), len(v2)
        i = 0
        while i < len1 or i < len2:
            a = int(v1[i]) if i < len1 else 0
            b = int(v2[i]) if i < len2 else 0
            i += 1
            if a < b:
                return -1
            elif a > b:
                return 1
        return 0
