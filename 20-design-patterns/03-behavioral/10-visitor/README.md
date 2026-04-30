# Visitor

- Allows adding extra behaviors to entire hierarchies of types
- Lets you separate algorithms from objects on which they operate
- A pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single `Accept(v *Visitor)` method throughout the entire hierarchy.
- Create `Visitor` with `VisitFoo(f Foo)`, `VisitBar(v Bar)` etc for each element in the hierarchy.
- Each `Accept(v *Visitor)` simply calls `Visitor.VisitXxx(self)`
- It suggests that you place the new behavior into a separate class called *visitor*, instead of trying to integrate it into existing classes.
- The original object that had to perform the behavior is now passed to one of the visitor’s methods as an argument, providing the method access to all necessary data contained within the object.

## Dispatch

- It answers question, *Which function to call?*
- **Single Dispatch**: depends on name of request and type of receiver.
- **Double Dispatch**: depends on name of request and type of two receivers (type of visitor and type of element being visited.)
- [Visitor and Double Dispatch - refactoring.guru/](https://refactoring.guru/design-patterns/visitor-double-dispatch)

## Example
