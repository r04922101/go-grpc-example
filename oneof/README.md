# Oneof

## Proto

```protobuf
message HelloRequest {
  oneof name {
    string first_name = 1;
    string last_name = 2;
  }

  oneof gender {
    bool male = 3;
    bool female = 4;
  }

  oneof default_values {
    int32 default_int = 5;
    string default_string = 6;
  }
}
```

I defined 3 `oneof` values, and I will use each for different purposes.

- `name`: set a non-default value. 
- `gender`: not set.
- `default_values`: set a default value.

## Outcome

- In generated Go code, each field is compiled into a `struct` which implements a generated `interface` and contains its real value.
- We can use type assertion to determine which `oneof` field has been set.
- On the other hand, if a `oneof` field is set to default value, the value will be serialized to the wire, and it is diffrent from not being set, which is `nil` struct.