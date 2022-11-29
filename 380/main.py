import random


class RandomizedSet:

    def __init__(self):
        self.map = {}
        self.array = []

    def insert(self, val: int) -> bool:
        if val not in self.map:
            self.map[val] = len(self.array)
            self.array.append(val)
            return True
        return False

    def remove(self, val: int) -> bool:
        if val in self.map:
            self.map[self.array[-1]] = self.map[val]
            self.array[self.map[val]] = self.array[-1]
            self.array.pop()
            self.map.pop(val)
            return True
        return False

    def getRandom(self) -> int:
        return random.choice(self.array)


# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
