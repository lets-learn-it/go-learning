# Visitor

- Lets you separate algorithms from objects on which they operate
- It suggests that you place the new behavior into a separate class called *visitor*, instead of trying to integrate it into existing classes.
- The original object that had to perform the behavior is now passed to one of the visitor’s methods as an argument, providing the method access to all necessary data contained within the object.