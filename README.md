# gopurs-assert

## Local Go development

This checkout is part of the gopurs library family. Use the
[local Go development guide](../gopurs/README.md#develop-one-library-locally)
for toolchain setup, sibling dependencies, Spago configuration and Go commands.
The existing npm, Bower and Dhall commands below retain their JavaScript or
upstream roles.


[![Latest release](http://img.shields.io/github/release/purescript/purescript-assert.svg)](https://github.com/purescript/purescript-assert/releases)
[![Build status](https://github.com/purescript/purescript-assert/workflows/CI/badge.svg?branch=master)](https://github.com/purescript/purescript-assert/actions?query=workflow%3ACI+branch%3Amaster)
[![Pursuit](https://pursuit.purescript.org/packages/purescript-assert/badge)](https://pursuit.purescript.org/packages/purescript-assert)

Basic assertions library for low level testing. This is primarily for testing the core libraries that cannot use [`purescript-quickcheck`](https://github.com/purescript/purescript-quickcheck) without resulting in circular dependencies.

## Installation

```
spago install assert
```

## Documentation

Module documentation is [published on Pursuit](http://pursuit.purescript.org/packages/purescript-assert).
