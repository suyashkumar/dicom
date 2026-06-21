## 2024-05-24 - Buffered Reading in Hot Loops
**Learning:** Calling `io.ReadFull` (even on a buffered reader) millions of times in a loop (e.g. for every pixel) is extremely expensive due to function call overhead and internal checks.
**Action:** Always batch reads into chunks (e.g. 4KB) in hot loops, then process from the buffer. This yielded a 3x-6x speedup in DICOM parsing.
