# Sudoku Solver

Attempt to learn Go by write a code to solve standard sudoku 9 by 9 matrix. No machine learning no AI.

## Features

Going to implement with 2 solvers.

1. WFC (Wave Function Collapse)
2. Human Intuition (Choice removal based on inference)

## Usage

Given input the sudoku matrix of.

```
+-----+-----+-----+
|     |    1|4   6|
|5   1|  2 9|     |
|  8  |    5|    2|
+-----+-----+-----+
|  3 2|     |5    |
|     |  2  |    2|
|  6  |  3  |     |
+-----+-----+-----+
|1    |    3|7    |
|8 2  |  9  |9    |
|    3|  4  |     |
+-----+-----+-----+
```

```bash
echo '     14 6' > input.txt
echo '5 1 29   ' >> input.txt
echo ' 8   5  2' >> input.txt
echo ' 32   5  ' >> input.txt
echo '      2  ' >> input.txt
echo '   2 6  3' >> input.txt
echo '    1    ' >> input.txt
echo '37  82  9' >> input.txt
echo '9     3 4' >> input.txt

go run main.go input.txt
```

