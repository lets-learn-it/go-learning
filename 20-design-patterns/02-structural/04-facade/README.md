# Facade

- It is structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.

## Problem

- Your code work with broad set of objects that belong to a library or framework. You need to initialize all of those objects, keep track of dependencies, execute methods in the correct order and so on.
- **Business login of your classes would become tightly coupled to the implementation details of 3rd-party library.**

## Solution

- Add class between your business login and 3rd-party library which simplifies interface for your business logic. This class is called facade.
- **facade might provide limited functionality in comparison to working with the subsystem directly**.
- You can swap library very easily because of facade class.
- Cons: facade can become **god object** coupled to all classes of an app.
