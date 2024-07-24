# Lagrange Interpolation

This repository contains the implementation of the Lagrange Interpolation method in Go. It will be simpler and easier to understand for the sake of learning.

## Formula

The Lagrange Interpolation formula is given by:

```
f(x) = Σ(yi * Li(x))
```

Where `Li(x)` is the Lagrange basis polynomial given by:

```
Li(x) = Π((x - xj) / (xi - xj))
```

where:

## Reference

- [Lagrange Interpolation](https://en.wikipedia.org/wiki/Lagrange_polynomial)
