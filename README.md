# QATrader_Go
golang version  QATRADER


The D family of types is used to concisely build BSON objects using native Go types. This can be particularly useful for constructing commands passed to MongoDB. The D family consists of four types:

D: A BSON document. This type should be used in situations where order matters, such as MongoDB commands.

M: An unordered map. It is the same as D, except it does not preserve order.

A: A BSON array.

E: A single element inside a D.
