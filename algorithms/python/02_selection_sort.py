# Функция поиска наименьшего значения в массиве
def find_smallest(arr):
    value = arr[0]
    index = 0
    for i in range(1, len(arr)):
        if arr[i] < value:
            value = arr[i]
            index = i
    return index


def selectionSort(arr):
    newArr = []
    for i in range(len(arr)):
        # Находим индекс наименьшего элемента
        smallest = find_smallest(arr)
        # Удаляем его из исходного массива и добавляем в новый
        newArr.append(arr.pop(smallest))
    return newArr


# Пример использования
print(selectionSort([5, 3, 6, 2, 10]))
