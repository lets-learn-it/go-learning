# Command

- It turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delays or queue a request's execution, and support undoable operations.

## Implementation

- Declare command interface with single execution method
  ```go
  type Command interface {
    execute()
  }
  ```
- Start extracting requests into concrete command classes that implement command interface. Each class must have a set of fields for storing the request arguments along with a reference to the actual receiver object.
  ```go
  // reciever
  type Device interface {
    on()
    off()
  }

  type OnCommand struct {
    device Device    // receiver object
  }

  func (c *OnCommand) execute() {
    c.device.on()
  }

  type OffCommand struct {
    device Device    // receiver object
  }

  func (c *OffCommand) execute() {
    c.device.off()
  }
  ```
- Identify classes that will act as senders. Add fields for storing commands into these classes. Senders should communicate with their commands only via command interface. Senders usually dont create command objects on their own, but rather get them from client code.
  ```go
  type Button struct {
    command Command
  }

  func (b *Button) press() {
    b.command.execute()
  }
  ```
-