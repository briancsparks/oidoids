# Oidoids





## Notes

A central idea is that each oid is a running goroutine (like the Boids example.) This allows
each oid to keep its state internal without locks using the Actor model (basically.) But
using actors does not give us thread-safe access to the environment.

### One

The first thought was that the state of the environment is pushed into the oid's goroutine
over channels (that's what they're for), but who would be pushing that data? How would that
entity know what data my oid is interested in?

### Two

This is the new idea.

When each oid is deciding what to do, they can all access a read-only view of the environment.
Thus, when most processing is occurring (when oids are deciding what to do), all the access
to the environment is to read the data, and hence just a single lock for the entire
environment. This should be easy using immutable techniques.

Then, each oid's manipulation of the environment is done in a second "offscreen" environment.
The offscreen is managed as an Actor and each oid pushes its desired changes over a channel.
From an oid's point-of-view, my only way to manipulate the environment is to move myself
to another location, so I send a message to the offscreen that I want to move from point
A to B, and the offscreen will do it, and maybe do other side-effects.



