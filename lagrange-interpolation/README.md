# Lagrange Interpolation

This repository contains the implementation of the Lagrange Interpolation method in Go. It will be simpler and easier to understand for the sake of learning.

## Formula

The Lagrange Interpolation formula is given by:

$$
f(x) = \sum_{i=0}^{n} f(x_i) \cdot L_i(x)
$$

where:

- \( f(x) \) is the function that we want to interpolate
- \( x_i \) is the data point
- \( f(x_i) \) is the value of the function at \( x_i \)
- \( L_i(x) \) is the Lagrange basis polynomial
- \( x \) is the point at which we want to interpolate \( f(x) \)

We can also write the formula as:

$$
f(x) = \sum_{i=0}^{n} f(x_i) \cdot \prod_{\substack{0 \leq j \leq n \\ j \neq i}} \frac{x - x_j}{x_i - x_j}
$$

Where \( L_i(x) \) is the Lagrange basis polynomial given by:

$$
L_i(x) = \prod_{\substack{0 \leq j \leq n \\ j \neq i}} \frac{x - x_j}{x_i - x_j}
$$

## Reference

- [Lagrange Interpolation](https://en.wikipedia.org/wiki/Lagrange_interpolation)
