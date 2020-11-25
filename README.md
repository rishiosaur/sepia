![](assets/banner.svg)
# 🌇 Sepia
![Github Build status](https://img.shields.io/github/workflow/status/rishiosaur/sepia/build?event=push&style=flat-square)
![GitHub Workflow Status (event)](https://img.shields.io/github/workflow/status/rishiosaur/sepia/lint?event=push&label=lint&style=flat-square)
![GitHub](https://img.shields.io/github/license/rishiosaur/sepia?style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/rishiosaur/sepia?style=flat-square)
![GitHub contributors](https://img.shields.io/github/contributors/rishiosaur/sepia?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/rishiosaur/sepia?style=flat-square)

A minimal functional programming language.

## Installation

On macOS/Linux systems with `Homebrew` installed, you can add Sepia by running `brew install rishiosaur/taps/sepia`.

To test `sepia` out, run `sepia` in any folder!

### Running examples

Clone `rishiosaur/sepia`, then change directories into the cloned directory. Run `sepia examples/types.sp` as a first example.

### VS Code Extension

I've also designed a minimal VSC toolset around Sepia, which you can find [on the VS Code Marketplace](https://marketplace.visualstudio.com/items?itemName=rishiosaur.sepia).

---

## 🎨 Design Goals
Sepia started off as a project to get to learn interpreters in-depth in as short a timeframe as possible (I'll always remember those 20 hours).

However, it ended up being an idealized version of what I believe programming languages should be: it's a project that's aimed squarely at replacing some of the more well-known languages in my workflow.

### 📚 Readability.

I'm sick and tired of programming languages not being intuitive to write. Whether that be understanding the way that `this` works in JS or figuring out why type checking syntax is the way it is in Go, nearly every major programming language has some unintuitive piece of it.

Sepia aims to replace every bit of unclear syntax with something that makes sense (usually borrowed from a *mathematical* convention). Every piece of syntax and semantic should have some reason for its implementation.

### 🛠 Functionality (alternatively, the destruction of OOP).

I've always hated the notion of classes and methods 'on' types: it's obtuse, and doesn't lead to great generalizations.

That's why structures are exactly (and exclusively) what they sound like: a type that defines the structure of data. All functions that're supposed to work 'on' those types are just functions with an argument that requires that type. This not only means that parsing is easier, strong typing is available WHEN NEEDED, and programs execute faster, but it's also a *lot* easier to understand, and allows for incredible generalizations.

### 👁 Minimalism.

When possible, Sepia tries to reuse existing syntax or keywords, and places a focus on clever algorithmic work and efficiency (wrapping the interpreter in Go gives Sepia easy concurrency for asynchronous applications): it doesn't even have loops!

---