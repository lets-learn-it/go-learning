# Mediator

- It lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.
- It makes easy to modify, extend and reuse individual components because they're no longer dependent on the dozens of other classes.
- It is similar to `facade` as it abstracts functionality of existing classes.
  **Differences**
  - Mediator abstracts/centralizes arbitrary communication between colleague objects, it routinely "adds value" , and it is known/referenced by colleague objects. (it defines a multidirectional protocol)
  - Facade defines simpler interface to a subsystem, it doesn't add new functionality, and it is not known by subsystem classes. (it defines unidirectional protocol)
