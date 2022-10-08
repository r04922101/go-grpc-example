# Update Field

## What is Updated

1. Rename field 1 from `name` to `message`.
2. Reuse/repurpose field 2, change it from `int32` type to `bool` type, rename it from `number` to `is_new`.
3. Deprecate field 3.
4. Introduce a new field 4 `timestamp`.

## New Server + Old Client

### Server

The server receives a request and print each field.

### Client

The client sends a request:

```go
{
    "Name": "Tony Huang",
    "Number": 101,
    "Sport": pb.HelloRequest_SPORT_BASEBALL
}
```

### Outcome

The server prints

```bash
2022/10/08 15:01:21 received message from client "Tony Huang", is new: true, timestamp: <nil> 
```

The server output that:
1. Renaming a field without changing its type is fine, but need to take care of its underlying meaning.
2. An integer with value `101` is deserialized into a boolean value `true`.
3. Reserved field is ignored.
4. New field (timestamp) is in its default value, which is `nil`.

## Old Server + New Client

### Server

The server receives a request and print each field.

### Client

The client sends a request:

```go
{
		Message:   "Hello!",
		IsNew:     true,
		Timestamp: timestamppb.New(time.Now()),
}
```

### Outcome

The server prints

```bash
2022/10/08 15:11:56 received message from client "Hello!", number: 1, sport: SPORT_UNKNOWN
```

The server output shows that:
1. Renaming a field without changing its type is fine, but need to take care of its underlying meaning.
2. An boolean with value `true` is deserialized into an integer value `1`.
3. Sport field (deprecated field in new proto) is in its default value.
4. New field (timestamp) is ignored.