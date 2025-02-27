> api: https://numpy.com.cn/doc/stable/reference/generated/numpy.argsort.html#numpy.argsort


### np.argwhere
```python
numpy.argwhere(condition)
```
- 用于返回满足特定条件的元素的索引。它通常用于查找数组中符合某个条件的元素的位置。

```python
import numpy as np

# 创建一个示例数组
arr = np.array([[1, 2, 3],
                 [4, 5, 6],
                 [7, 8, 9]])

# 查找大于 5 的元素的索引
indices = np.argwhere(arr > 5)

print("大于 5 的元素的索引:")
print(indices)
```
```python
大于 5 的元素的索引:
[[1 2]
 [2 0]
 [2 1]
 [2 2]]
```

### np.unique
```python
numpy.unique(arr, return_index=False, return_inverse=False, return_counts=False, axis=None)
```
- 用于找到数组中的唯一值，并返回这些唯一值的排序列表。这个函数在数据分析和处理时非常有用，尤其是在需要去重或统计元素时。
- 参数
  - arr：输入数组，可以是一维或多维数组。
  - return_index：如果为 True，则返回唯一值在原数组中的索引。
  - return_inverse：如果为 True，则返回原数组中每个元素在唯一值数组中的索引。
  - return_counts：如果为 True，则返回每个唯一值的出现次数。
  - axis：指定沿哪个轴查找唯一值（适用于多维数组）。

```python
import numpy as np

# 创建一个示例数组
arr = np.array([1, 2, 2, 3, 1, 4, 4, 4])

# 找到唯一值
unique_values = np.unique(arr)

print("唯一值:", unique_values)
```
```
唯一值: [1 2 3 4]
```

返回索引
```python
unique_values, indices = np.unique(arr, return_index=True)

print("唯一值:", unique_values)
print("索引:", indices)
```
```python
唯一值: [1 2 3 4]
索引: [0 1 3 5]
```

返回计数
```python
unique_values, counts = np.unique(arr, return_counts=True)

print("唯一值:", unique_values)
print("出现次数:", counts)
```
```python
唯一值: [1 2 3 4]
出现次数: [2 2 1 3]
```

### numpy.argsort
```python
numpy.argsort(a, axis=-1, kind=None, order=None, *, stable=None)[source]
```
- 返回能够对数组进行排序的索引。
- 使用 kind 关键字指定的算法，沿给定轴执行间接排序。它返回一个与 a 形状相同的索引数组，这些索引按排序顺序索引沿给定轴的数据。
- 参数：
  - a: array_like
    - 要排序的数组。
  - axis: int 或 None，可选
    - 要排序的轴。默认为 -1（最后一个轴）。如果为 None，则使用扁平化的数组。
  - kind: {'quicksort', 'mergesort', 'heapsort', 'stable'}，可选
    - 排序算法。默认为 'quicksort'。请注意，'stable' 和 'mergesort' 都使用 timsort 算法，并且通常，实际的实现会随着数据类型的变化而变化。保留 'mergesort' 选项是为了向后兼容。
  - order: str 或 str 列表，可选
    - 当 a 是一个具有定义字段的数组时，此参数指定首先比较哪些字段，其次比较哪些字段，依此类推。单个字段可以指定为字符串，并且不需要指定所有字段，但未指定的字段仍将按其在 dtype 中出现的顺序使用，以打破平局。
  - stable: bool，可选
    - 排序稳定性。如果为 True，则返回的数组将保持比较结果相等的 a 值的相对顺序。如果为 False 或 None，则不保证这一点。内部，此选项选择 kind='stable'。默认值：None。

版本 2.0.0 中的新功能。

- 返回：
  - index_array: ndarray, int
    - 沿指定 axis 排序 a 的索引数组。如果 a 是一维数组，则 `a[index_array]` 将产生排序后的 a。更一般地说，`np.take_along_axis(a, index_array, axis=axis)` 始终产生排序后的 a，无论其维度如何。

> 另请参见
> 
> sort: 描述使用的排序算法。
> 
> lexsort: 具有多个键的间接稳定排序。
> 
> ndarray.sort: 就地排序。
> 
> argpartition: 间接部分排序。
> 
> take_along_axis: 将 argsort 中的 index_array 应用于数组，如同调用 sort 一样。

备注
- 有关不同排序算法的说明，请参见 sort。
- 从 NumPy 1.4.0 开始，argsort 可用于包含 NaN 值的实数/复数数组。增强的排序顺序在 sort 中有说明。

示例

一维数组
```python
import numpy as np
x = np.array([3, 1, 2])
np.argsort(x)
array([1, 2, 0])
```
二维数组
```python
x = np.array([[0, 3], [2, 2]])
x
array([[0, 3],
       [2, 2]])
ind = np.argsort(x, axis=0)  # sorts along first axis (down)
ind
array([[0, 1],
       [1, 0]])
np.take_along_axis(x, ind, axis=0)  # same as np.sort(x, axis=0)
array([[0, 2],
       [2, 3]])
ind = np.argsort(x, axis=1)  # sorts along last axis (across)
ind
array([[0, 1],
       [0, 1]])
np.take_along_axis(x, ind, axis=1)  # same as np.sort(x, axis=1)
array([[0, 3],
       [2, 2]])
```
N 维数组的已排序元素的索引
```python
ind = np.unravel_index(np.argsort(x, axis=None), x.shape)
ind
(array([0, 1, 1, 0]), array([0, 0, 1, 1]))
x[ind]  # same as np.sort(x, axis=None)
array([0, 2, 2, 3])
```
使用键进行排序
```python
x = np.array([(1, 0), (0, 1)], dtype=[('x', '<i4'), ('y', '<i4')])
x
array([(1, 0), (0, 1)],
      dtype=[('x', '<i4'), ('y', '<i4')])
np.argsort(x, order=('x','y'))
array([1, 0])
np.argsort(x, order=('y','x'))
array([0, 1])
```