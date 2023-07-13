Учитывая целочисленный массив nums, отсортированный в неубывающем порядке , 
удалите дубликаты на месте, чтобы каждый уникальный элемент появлялся только один раз. 
Относительный порядок элементов должен быть сохранен . Затем верните количество уникальных элементов в nums.

Учитывайте количество уникальных элементов numsto be k, чтобы вас приняли, вам нужно сделать следующее:

Измените массив numsтаким образом, чтобы первые kэлементы numsсодержали уникальные элементы в том порядке, в котором они присутствовали numsизначально. Остальные элементы numsне так важны, как и размер nums.
Вернуться k.

Пример 1:
Ввод: nums = [1,1,2]
Вывод: 2, nums = [1,2,_]
Объяснение: Ваша функция должна возвращать k = 2, причем первые два элемента nums равны 1 и 2 соответственно.
Неважно, что вы оставляете за возвращенным k (следовательно, это символы подчеркивания).

Пример 2:
Ввод: числа = [0,0,1,1,1,2,2,3,3,4]
Выход: 5, числа = [0,1,2,3,4,_,_,_,_, _]
Объяснение: Ваша функция должна возвращать k = 5, причем первые пять элементов nums равны 0, 1, 2, 3 и 4 соответственно.
Неважно, что вы оставляете за возвращенным k (следовательно, это символы подчеркивания).