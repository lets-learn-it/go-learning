# Proxy

- A type that functions as an interface to a particular resource. That resource may be remote, expensive to construct or may require logging or some other added functionality.
- Proxy lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.
- **Proxy keeps same interface as of original service**
- **Decorator and Proxy have similar structures**, but very different intents. Both patterns are built on the composition principle, where one object is supposed to delegate some of the work to another. The difference is that a Proxy usually manages the life cycle of its service object on its own, whereas the composition of Decorators is always controlled by the client. [[https://doeken.org/blog/decorator-vs-proxy-pattern]](https://doeken.org/blog/decorator-vs-proxy-pattern)

## Summary

- Proxy has the same interface as the underlying object
- To create a proxy, simply replicate the existing interface of an object
- Add relevant functionality to the redefined methods.
- Different proxies like communication, logging, caching etc
