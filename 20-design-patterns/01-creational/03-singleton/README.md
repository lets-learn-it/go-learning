# Singleton

- This design pattern ensure that a class has only one instance, while providing a global access point to this instance.

## Problem

- How to controll access to some shared resource like file or database? May be using single object shared between all & that object locks resource internally.
- How to handle global variables or data? It should not happen that some code is overwriting contents of those variables and crashing app.

## Solution

- Don't allow direct creation of object. (Make default constructor private)
- Create a method (static) which acts as a constructor. Under the hood, method calls private constructor to create an object and saves it. All following calls return cached object.
