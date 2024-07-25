# Multiples of X,Y...N

This program was part of [Problem 1](https://projecteuler.net/problem=1) of Project Euler. Where we need to find the sum of all the multiples of 3 or 5 below 1000.

Here is generalized so we can do:

```bash
:!go run main.go 1000 3 5 -15

[3 5 -15]
Multiples of: 3 [=166833.000000]                                                            Multiples of: 5 [=99500.000000]                                                         
Multiples of: -15 [=-33165.000000]                                                      
Total multiples sum is: 233168
```

This CLI supports `negative` numbers as well. (as `-15`)
