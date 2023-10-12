import numpy as np

N = 10
TOP = 0
LEFT = 1
DOWN = 2
RIGHT = 3

def main():
    m = np.zeros((N, N), dtype=int)
    i, j = N // 2, N // 2
    iMin, iMax = i, i
    jMin, jMax = j, j
    d = TOP
    for step in range(1, N*N+1):
        m[i][j] = step
        if d == TOP:
            i -= 1
            if i < iMin:
                d = LEFT
                iMin = i
        elif d == LEFT:
            j -= 1
            if j < jMin:
                d = DOWN
                jMin = j
        elif d == DOWN:
            i += 1
            if i > iMax:
                d = RIGHT
                iMax = i
        elif d == RIGHT:
            j += 1
            if j > jMax:
                d = TOP
                jMax = j
    printSpiral(m)

def printSpiral(spiral):
    print(f"matrix: {N}x{N}")
    print("===========")
    for i in range(N):
        for j in range(N):
            print(f"{spiral[i][j]}\t", end="")
        print()

main()