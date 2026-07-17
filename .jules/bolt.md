
## 2025-12-17 - Replacing binary.Read with direct ByteOrder calls
**Learning:** `binary.Read` uses reflection which is expensive for frequent small reads. Direct usage of `io.ReadFull` and `binary.ByteOrder` interface provides significant speedup (~20% in this case) while maintaining code clarity.
**Action:** When reading primitive types in performance-critical paths, prefer manual decoding using `binary.ByteOrder` or `math` over `binary.Read`.
