class Allocator:

    def __init__(self, n: int):
        self.__n = n
        self.__memory = [0]*n
        

    def allocate(self, size: int, mID: int) -> int:
        cnt = 0
        for i in range(self.__n):
            if self.__memory[i] != 0:
                cnt = 0
            else:
                cnt += 1
                if cnt == size:
                    for j in range(i-cnt+1, i+1):
                        self.__memory[j] = mID
                    return i-cnt+1
        return -1
        

    def freeMemory(self, mID: int) -> int:
        cnt = 0
        for i in range(self.__n):
            if self.__memory[i] == mID:
                self.__memory[i] = 0
                cnt += 1
        return cnt