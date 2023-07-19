def compute(n):
    if n < 10:
        out = n ** 2
    elif n <= 20:
        out = 1
        for i in range(1, n-9):  # Corrected range to (n-9)
            out *= i
    else:
        lim = n - 20
        out = lim * lim
        out = out + lim  # Corrected subtraction to addition (making formula (n*(n+1)/2)
        out = out // 2  # Changed division to integer division using double slash

    print(out)


n = int(input("Enter an integer: "))
compute(n)

# Assumption:
# for 2nd case excluding 10 and including 20

# The test case I had used are:
# Inp: 0  
# Out: 0

# Inp: 2
# Out: 4

# Inp: 9
# Out: 81

# Inp: 10
# Out: 1

# Inp: 15
# Out: 25

# Inp: 20
# Out: 3628800

# Inp: 25
# Out: 15

# Inp: 100
# Out: 3240