from concurrent import futures
from concurrent.futures.thread import _worker
import math
import time
import sys

def calc(val: float) -> float:
    time.sleep(1)
    result= math.sqrt(val)
    return result

def use_threads(num: int, values) -> float:
    t1 = time.time()
    with futures.ThreadPoolExecutor(num) as tex:
        results = tex.map(calc, values)
    t2 = time.time()
    return t2 - t1

def use_processes(num: int, values) -> float:
    t1 = time.time()
    with futures.ProcessPoolExecutor(num) as pex:
        reuslts = pex.map(calc, values)
    t2 = time.time()
    return t2 - t1

def main(workers, values):
    print(f"Using {workers} workers for {len(values)} values")
    t_sec = use_threads(workers, values)
    print(f"Threads took {t_sec:.4f} seconds")
    p_sec = use_processes(workers, values)
    print(f"Processes took {p_sec:.4f} seconds")

if __name__ == "__main__":
    workers = int(sys.argv[1])
    values = list(range(1,6))
    main(workers, values)