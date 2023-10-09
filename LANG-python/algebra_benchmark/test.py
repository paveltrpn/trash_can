
from arr_alg import mtrx3Arr, mtrx3ArrFromEuler, CVec3Arr
from np_alg import mtrx3Np, mtrx3NpFromEuler, CVec3Np
from list_alg import mtrx3List, mtrx3ListFromEuler, CVec3List
from glm_alg import mtrx3GlmFromEuler, CVec3Glm

import timeit
import time

boxTrisArr = CVec3Arr([
    #  top
     1.0,  1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,
    -1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0, -1.0,
    #  bottom
     1.0, -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,
    -1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  front
    -1.0,  1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
     1.0,  1.0,  1.0,  1.0, -1.0,  1.0, -1.0, -1.0,  1.0,
    #  back
    -1.0,  1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0,
     1.0,  1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  left
    -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
    -1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0,  1.0, -1.0,
    #  right
     1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,  1.0,
     1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0
])

boxTrisNp = CVec3Np([
    #  top
     1.0,  1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,
    -1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0, -1.0,
    #  bottom
     1.0, -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,
    -1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  front
    -1.0,  1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
     1.0,  1.0,  1.0,  1.0, -1.0,  1.0, -1.0, -1.0,  1.0,
    #  back
    -1.0,  1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0,
     1.0,  1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  left
    -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
    -1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0,  1.0, -1.0,
    #  right
     1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,  1.0,
     1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0
])

boxTrisList = CVec3List([
    #  top
     1.0,  1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,
    -1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0, -1.0,
    #  bottom
     1.0, -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,
    -1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  front
    -1.0,  1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
     1.0,  1.0,  1.0,  1.0, -1.0,  1.0, -1.0, -1.0,  1.0,
    #  back
    -1.0,  1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0,
     1.0,  1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  left
    -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
    -1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0,  1.0, -1.0,
    #  right
     1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,  1.0,
     1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0
])

boxTrisGlm = CVec3Glm([
    #  top
     [1.0,  1.0,  1.0], [-1.0,  1.0,  1.0],  [1.0,  1.0, -1.0],
    [-1.0,  1.0,  1.0],  [1.0,  1.0, -1.0], [-1.0,  1.0, -1.0],
    #  bottom
    [ 1.0, -1.0,  1.0], [-1.0, -1.0,  1.0], [ 1.0, -1.0, -1.0],
    [-1.0, -1.0,  1.0], [ 1.0, -1.0, -1.0], [-1.0, -1.0, -1.0],
    #  front
    [-1.0,  1.0,  1.0],  [1.0,  1.0,  1.0], [-1.0, -1.0,  1.0],
    [ 1.0,  1.0,  1.0],  [1.0, -1.0,  1.0], [-1.0, -1.0,  1.0],
    #  back
    [-1.0,  1.0, -1.0],  [1.0,  1.0, -1.0], [-1.0, -1.0, -1.0],
    [ 1.0,  1.0, -1.0],  [1.0, -1.0, -1.0], [-1.0, -1.0, -1.0],
    #  left
    [-1.0,  1.0, -1.0], [-1.0,  1.0,  1.0], [-1.0, -1.0,  1.0],
    [-1.0, -1.0,  1.0], [-1.0, -1.0, -1.0], [-1.0,  1.0, -1.0],
    #  right
    [ 1.0,  1.0, -1.0], [ 1.0,  1.0,  1.0], [ 1.0, -1.0,  1.0],
    [ 1.0, -1.0,  1.0], [ 1.0, -1.0, -1.0], [ 1.0,  1.0, -1.0]
])

rtnMtrx3Arr     = mtrx3ArrFromEuler(45.0, 45.0, 45.0)
rtnMtrx3Np      = mtrx3NpFromEuler(45.0, 45.0, 45.0)
rtnMtrx3List    = mtrx3ListFromEuler(45.0, 45.0, 45.0)
rtnMtrx3Glm     = mtrx3GlmFromEuler(45.0, 45.0, 45.0)

ITERATIONS = 10000


def main() -> None:
    print("Apply {} times mtrx3Arr to CVec3Arr.".format(ITERATIONS))
    bench = timeit.Timer(lambda: boxTrisArr.applyMtrx3(rtnMtrx3Arr)).timeit(number=ITERATIONS)
    print("time is {} ms".format(int(round(bench*1000))))
    # for _ in range(ITERATIONS):
        # boxTrisArr.applyMtrx3(rtnMtrx3Arr)

    print("Apply {} times mtrx3Np to CVec3Np.".format(ITERATIONS))
    bench = timeit.Timer(lambda: boxTrisNp.applyMtrx3(rtnMtrx3Np)).timeit(number=ITERATIONS)
    print("time is {} ms".format(int(round(bench*1000))))
    # for _ in range(ITERATIONS):
        # boxTrisNp.applyMtrx3(rtnMtrx3Np)

    print("Apply {} times mtrx3List to CVec3List.".format(ITERATIONS))
    bench = timeit.Timer(lambda: boxTrisList.applyMtrx3(rtnMtrx3List)).timeit(number=ITERATIONS)
    print("time is {} ms".format(int(round(bench*1000))))
    # for _ in range(ITERATIONS):
        # boxTrisList.applyMtrx3(rtnMtrx3List)

    print("Apply {} times mtrx3Glm to CVec3Glm.".format(ITERATIONS))
    bench = timeit.Timer(lambda: boxTrisGlm.applyMtrx3(rtnMtrx3Glm)).timeit(number=ITERATIONS)
    print("time is {} ms".format(int(round(bench*1000))))
    # for _ in range(ITERATIONS):
        # boxTrisGlm.applyMtrx3(rtnMtrx3Glm)


if __name__ == "__main__":
    main()