from typing import List


def binary_search(find: int, mas: List[int]) -> int | None:
    low = 0
    high = len(mas) - 1
    while low <= high:
        mid = (low + high) // 2
        val = mas[int(mid)]
        if val == find:
            return val
        if val < find:
            low = mid + 1
        else:
            high = mid - 1
    return None


find = 10
mas = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
res = binary_search(find, mas)
print(res)
