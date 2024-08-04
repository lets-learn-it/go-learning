# Strategy

- It lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

## Implementation

- In context class (this class has data), identify an algorithm thats prone to frequent changes. It may also be a massive conditional that selects and executes a variant of same algorithm at runtime.
- Declare strategy interface common to all variants of the algorithm.
- One by one, extract all algorithms into their own classes. They should all implement the strategy interface.
- In context class, add a field for storing a reference to a strategy object. The context may define an interface which lets the strategy access its data.
- Clients of the context must associate it with a suitable strategy that matches the way they expect the context to perform its primary job.
