# Wrapper for func

Is a func that take a function as an argument and return a new function that does something before or after calling the original function.

It also called middleware

Use Cases
Wrappers and decorators find applications in various scenarios, such as:

- Logging: Wrapping an object to automatically log method calls, parameters, and return values.
- Caching: Decorating an object to cache results of expensive operations, improving performance by avoiding repeated computations.
- Authentication/Authorization: Adding a security layer to enforce access control rules on methods or resources.
- Rate Limiting: Throttling the usage of a resource or API by wrapping it with a decorator that enforces limits on the number of calls or frequency of usage.
- Monitoring: Instrumenting an object to collect metrics, track resource usage, or trigger alerts when certain conditions are met.
