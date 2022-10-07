GoF structural patterns

## Adapter : provides a wrapper with an interface to make incompatible type compatible.enables collaboration between objects with incompatible interfaces

- comprises of
  - target: interface the client calls and invoke method on adapter and adaptee
  - adapter: converts incompatible interface of the adaptee to interface the client wants
  - apatee
  - client: wants incompatible interface implemented by the adapter

## Bridge : allows seperation of abstraction from it's implementation. allows implementation detail to change at runtime. it decouple implementation from abstraction 🤯

composition over inheritance

-abstraction: interface implmented as abstract class
-refined abstraction
-implementer
-concrete implementer

## composite: allows grouping of similar objects in a single object in a tree like form.

comprises of:

- component interface: defines default behaviour of all objects and behaviour for accessing the component of the composite
- composite and component classes: implement the compoonent interface
- client: interacts with the component interface to invoke the composite

# decomposer

# facade

# flyweight

# private class
