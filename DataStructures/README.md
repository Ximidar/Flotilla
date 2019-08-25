# Data Structures

This is for storing common data structured between Flotilla Programs. This houses all the protobuffer code.

Generate JS proto:

```
protoc --js_out=import_style=commonjs,binary:js_proto/ FileStructures.proto
```