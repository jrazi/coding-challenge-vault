# Problem Statement

**Time limit:** 2.5 seconds  
**Memory limit:** 256 MB  

In a city management game, you have to manage 3 buildings in such a way that their electricity consumption reaches the most optimal possible state. To understand this, a referee in the game at each stage performs tasks to count the number of negative points. These tasks are one of the following two types:

1. **Judge of unit electricity consumption:**  
   Change the electricity consumption of unit `k` from building `a` to `x`. This is expressed as `a_k = x`.

2. **Arbitration by giving the amount `r` followed by three units `i`, `j`, and `k` of three buildings `a`, `b`, and `c`:**  
   The goal is to find the following pattern and gives the player one point for each occurrence of this pattern:  
   $$1 ≤ i < j < k ≤ r$$
   $$b_{a_i} = a_j = c_{a_k}$$ 

   Here, `a_i` is the unit power consumption of the building `a` at position `i`.

### Input

The first line of the input contains two integers `n` and `m` where:

- `n` is the number of units in the three buildings.
- `m` is the number of tasks that the referee performs.
  
`1 ≤ n ≤ 200,000`  
`1 ≤ m ≤ 500,000`

Next, there are three lines each containing `n` integers representing the electricity consumption of the building units `a`, `b`, and `c` respectively. 

`1 ≤ a_i, b_i, c_i ≤ n`

Following these, there are `m` lines of tasks in the following formats:

- **CHANGE(k, x):** Change the electricity consumption of unit `k` in the first sequence to `x`.
- **PRINT(r):** Output the number of points the player gets by finding the pattern within the first `r`.

### Output

For each task of the second type (`PRINT`), print the numerical value of the output.

### Examples

**Example of Input**: 

```plaintext
5 4
1 2 3 4 5
2 3 4 5 1
5 1 2 3 4
PRINT(5)
CHANGE(2,3)
PRINT(4)
PRINT(5)
```

**Expected Output**:

```plaintext
3
0
2
```