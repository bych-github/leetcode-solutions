## 题解 - 一次二分搜索 V.S. 两次二分搜索

- **一次二分搜索** - 由于矩阵按升序排列，因此可将二维矩阵转换为一维问题。对原始的二分搜索进行适当改变即可(求行和列)。时间复杂度为 `$$O(log(mn))=O(log(m)+log(n))$$`
- **两次二分搜索** - 先按行再按列进行搜索，即两次二分搜索。时间复杂度相同。

### 源码分析

仍然可以使用经典的二分搜索模板(lower bound)，注意下标的赋值即可。

1. 首先对输入做异常处理，不仅要考虑到matrix为null，还要考虑到matrix[0]的长度也为0。
2. 由于 lb 的变化处一定小于 target, 故在 else 中判断。