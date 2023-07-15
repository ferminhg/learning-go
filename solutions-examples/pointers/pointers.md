## The basics
To understand why pointers are even needed in go in the first place, we need to establish that go is a pass by value language.
That is, when passing a variable to a function, it will create a copy:

```go
type Person struct {
    name string
    age  int
}

func anonymize(p Person) {
    p.name = "Unknown" // p is a copy
}

func main() {
    person := Person{name: "John", age: 33}
    anonymize(person)
    fmt.Println(person.name) // Still prints out "John"
}
```

In other languages such as Java, the person's name would have changed. In these languages only primitive types such as integers are passed by value.

### How can we make the anonymize function change person?
For the sake of simplicity, let's assume that the memory address of the "person" instance is 27 (in reality it will be something like 0x140000a0030).

Let's pass in the address rather than the value as we initially did.

```go
anonymize(27)
```

The function receives a copy of 27. But this time around we're unconcerned that this is a copy because we can lookup the original object that resides at
the memory address of 27 and make changes to it:

```go
func anonymize(p MemoryAddressOfPerson) {
    // retrieve person object from the memory address 27
    // update name of the person now affects original
}
```

Let's implement what we just discussed.

In order to get the memory address of the person variable which we said it was 27 in our example, we can use `&`:

```go
anonymize(&person)
```

The `anonymize` function expects a Person, whereas we want to send it the memory address pointing to one. To do so we can use `*`:

```go
func anonymize(p *Person) {
    ...
}
```

So now, `p` is no longer an instance to a Person, but a pointer to that person.

In order to get the person instance from the pointer we need to use `*` to de-reference it:

```go
p  // a pointer to an instance of Person
*p // an instance of Person
```

Therefore the following will now work.
```go
func anonymize(p *Person) {
    (*p).name = "Unknown"
}
```

Changing the value of what the pointer points to, is quite common, and in go the following will work just as well.

```go
func anonymize(p *Person) {
    p.name = "Unknown"
}
```

Calling `anonymize(&person)` will now change the person.

### To recap, this is the syntax we learned about:
```
pointer of a type: *Type
pointer to a variable: &someVariable
value from a pointer variable: *somePointerVariable
```

## When to use pointers

### Mutability
Just like we seen in our previous example, this allows changing objects directly, and it's the main reason you should be using pointers.

### Optional values
Unlike regular types, pointers can be null.

These can be especially helpful when returning an error.
```go
func newPerson(name string) (*Person, error) {
    if len(name) == 0 {
        return nil, errors.New("Blank name provided")
    }

    return &Person{name: name}, nil
}
```

### Performance
Copying large structs and other big objects can be expensive. A pointer bypasses this issue.

On the other hand, passing smaller values, even though it implies a copy can be faster than passing a pointer.

In other words - are you passing an array with thousands of elements? Use a pointer. If not, decide based on other factors (like the two already mentioned).

Even so, this might not be needed if for instance you're using something like a slice. This is because internally it's made of a pointer to an array.
