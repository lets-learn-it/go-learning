# Observer

- Also known as Listener
- It lets you define subscription mechanism to notify multiple objects about any events that happen to the object they are observing.



- All objects which want to observe particular object has to satify predefined interface.
- Then those objects call subscribe method on publisher object. Publisher object keeps list of all subscribers and when event happens, it calls predifined method which is implemented by all subscribers.
