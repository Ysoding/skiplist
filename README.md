## SkipList


### Example

```go
  sl := NewSkipList[int, string]()
  sl.Insert(10, "test1")
  sl.Insert(30, "test1")
  sl.Insert(20, "test2")
  sl.Insert(40, "test2")

  v, exist := sl.Search(40)
  
  sl.Delete(40)
```
