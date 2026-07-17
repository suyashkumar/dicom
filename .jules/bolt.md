# Bolt's Journal

## 2024-05-22 - [Buffer Allocation in Hot Loops]
**Learning:** In Go, repeatedly allocating buffers inside a loop (especially in I/O intensive code) generates significant GC pressure.
**Action:** Use `sync.Pool` or pre-allocate buffers outside the loop when possible.

## 2024-05-22 - [Reflection in Binary Reading]
**Learning:** `binary.Read` uses reflection which is slow for simple types in hot paths.
**Action:** Use `io.ReadFull` with a byte slice and `binary.LittleEndian.UintXX` (or BigEndian) for primitive types in performance-critical sections.
