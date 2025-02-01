from typing import List

# https://leetcode.com/problems/array-nesting/description/

# Дан массив целых чисел nums длиной n, где nums является перестановкой чисел в диапазоне [0, n - 1].

# Вы должны построить множество s[k] = {nums[k], nums[nums[k]], nums[nums[nums[k]]], ...} при соблюдении следующего правила:

# Первый элемент в s[k] начинается с выбора элемента nums[k] с индексом k.
# Следующий элемент в s[k] должен быть nums[nums[k]], затем nums[nums[nums[k]]], и так далее.
# Мы прекращаем добавлять элементы непосредственно перед тем, как в s[k] появится дубликат.

# Верните длину самого длинного множества s[k].

# Input: nums = [5,4,0,3,1,6,2]
# Output: 4
# Explanation:
# nums[0] = 5, nums[1] = 4, nums[2] = 0, nums[3] = 3, nums[4] = 1, nums[5] = 6, nums[6] = 2.
# One of the longest sets s[k]:
# s[0] = {nums[0], nums[5], nums[6], nums[2]} = {5, 6, 2, 0}


def arrayNesting(nums: List[int]) -> int:
    visited = [False] * len(nums)
    max_length = 0
    for i in range(len(nums)):
        if not visited[i]:
            start = i
            count = 0
            while not visited[start]:
                visited[start] = True
                start = nums[start]
                count += 1
            max_length = max(max_length, count)
    return max_length


print(arrayNesting([4, 5, 0, 3, 1, 6, 2]))
