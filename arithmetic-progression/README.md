# Multiples of X,Y...N

This program was part of [Problem 1](https://projecteuler.net/problem=1) of Project Euler, where we need to find the sum of all the multiples of 3 or 5 below 1000.

## Usage

To use this code in its generalized form:

```bash
:!go run main.go 1000 3 5 -15

> [3 5 -15]
Multiples of: 3 [=166833.000000]
Multiples of: 5 [=99500.000000]                                                         
Multiples of: -15 [=-33165.000000]                                                      
Total multiples sum is: 233168
```

This CLI supports `negative` numbers as well. (as `-15`)

## Math Formulation

This problem is a classic arithmetic progression (AP) in which a sequence of numbers diverges with a constant ($d$) between consecutive terms.

The general form of an arithmetic progression can be written as:

$a, a + d, a + 2d, a + 3d, \ldots, a + (n-1)d$
where:
- $a$ is the first term,
- $d$ is the difference,
- $n$ is the number of terms between 0-$N$ where $N$ total size of the sum.

## Initial Term

This is the easiest part, and it's just the difference given by $k \times n$ where $k$ is the $difference$ we want to establish through terms.

## Last Term

The $n$-th term (also called $l$) of an AP can be calculated using the following formula:
$$
l = a + (n-1)d
$$

Furthermore, you need to find $n$ across `[0, N]`, so in order to get the latest possible term, you need to perform the following:

$$
\text{l} = 
\begin{cases} 
\left\lfloor \frac{N}{d} \right\rfloor \times d & \text{if } \left\lfloor \frac{N}{d} \right\rfloor \times d < 1000 \\
\left\lfloor \frac{N}{d} \right\rfloor \times d - d & \text{otherwise}
\end{cases}
$$

Additionally, after acquiring $l$, the rest of the formula is pretty straightforward. For example, after getting $l$ you can retrieve $n$ pretty quickly (the number of terms under `[0,N]` range) using the formula mentioned previously

## General Form

The final formula is dictated by:

$$
S_n = \frac{n}{2} (2a + (n-1)d)
$$
