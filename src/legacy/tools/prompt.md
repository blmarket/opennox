Create a go program which takes a go code and find the following patterns:

```
*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
```

1. find unsafe.Add with first operand being unsafe.Pointer type, and second operand can be resolved to a constant integer. (Just integer or multiply of integers)
2. it should be cast to a pointer of primitive types. Allowed primitive types are: `uint32`, `int32`.

Some example replacements:

```
*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
=> *&v1.field_12
(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
=> &v1.field_12
(*uint32)(unsafe.Add(unsafe.Pointer(v1), 48))
=> &v1.field_12
(*uint32)(unsafe.Add(unsafe.Pointer(v1), 36))
=> &v1.field_9
```

Create the program under convert_pointer_arithmetic/ and it can be run with following command:

```
go run ./convert_pointer_arithmetic/ <file.go>
```

It should print all list of places, and update the file in-place.
