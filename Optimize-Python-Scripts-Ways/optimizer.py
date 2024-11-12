"""
for optimizing the python script which is made for complex numerical computations do the following:
1. Memory Management
    Use __slots__: If your script defines classes, using __slots__ can reduce memory overhead by preventing the creation of a default __dict__ for each instance, leading to significant memory savings when many instances are created 1.
    Efficient Data Structures: Utilize built-in data structures like sets and dictionaries for faster lookups and memory efficiency compared to lists

2. Caching Techniques:
    For calling the same function multiple times with same parameters use the caching mechanism

3. Parallelization Techniques
Use multiprocessing: For CPU-bound tasks, leverage the multiprocessing module to run computations in parallel across multiple CPU cores. This can significantly reduce execution time
Use MultiThreading: For a I/O bound task, leverage the multi threading module to run computations in concurrent manner, and it uses the threading module. Also to take things furthermore you can use the asyncio module with aihttp module when you need to send the larger pool of requests and when you need to send the smaller pool of requests threading is good for those workloads.
"""
import functools
import time
class Point:
    __slots__ = ['x', 'y']

    def __init__(self, x, y):
        self.x = x
        self.y = y

p = Point(1, 2)
# p.z = 3

print(p.x)
# print(p.z)

# @functools.lru_cache(maxsize=10)
def compute_heavy_function(n):
    start_time = time.time()
    result = complex_function(n)
    end_time = time.time()
    print(result)
    print(f"Total Time taken: {end_time - start_time:.6f}")

def factorial(n):
    if n == 0 or n == 1:
        return 1
    else:
        return n * factorial(n - 1)

def complex_function(n):
    fact = factorial(n)
    return fact

compute_heavy_function(100)