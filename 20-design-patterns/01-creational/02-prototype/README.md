# Prototype

- Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

## Problem

- You want to create copy of object. you can create new object using same class and copy fields but what about private fields?
- What if you are accepting interface and not object? You are not aware of object at all.

## Solution

- Let delegate cloning process to actual object itself.
- Pattern declares a common interface for all objects that support cloning. This interface lets you clone an object without coupling your code to the class of that object.
