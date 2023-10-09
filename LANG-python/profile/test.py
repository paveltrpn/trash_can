
import mtrx4 as mtrx4
import vec3 as vec3
import numpy as np
import array as ar
from collections import namedtuple
#pip import pympler
import pympler as pm

from pympler import muppy
from pympler import summary
from pympler import asizeof
import glm as glm

class npVec3(object):
    __slots__ = ["data"]
    
    def __init__(self, x: float=0.0, y: float=0.0, z: float=0.0) -> None:
        self.data = np.array([x, y, z], dtype=np.float64)

class tpVec3(object):
    #  __slots__ allow us to explicitly declare data members (like properties) and deny the creation of __dict__ and __weakref__.
    # The space saved over using __dict__ can be significant".
    __slots__ = ["data"]

    def __init__(self, x: float=0.0, y: float=0.0, z: float=0.0) -> None:
        self.data = tuple([x, y, z])

vec3_t = ar.array("f", [0.0, 0.0, 0.0])

def foofoo(size: int):
    fooAr = ar.array("d",[val for val in range(size)])
    print("fooAr = {}".format(fooAr))
    print("fooAr size = {}".format(asizeof.asizeof(fooAr)))

def main() -> None:

    vFoo = vec3_t
    vFoo[1] = 100
    print(vFoo)

    allObj = muppy.get_objects()
    # foo = mtrx4.mtrx4()
    # bar = vec3.vec3()
    foo2 = npVec3()
    foo3 = tpVec3()
    fooAr = ar.array("d", [0.0, 1.0, 1.0, 1.0])
    vec3Lst = [vec3.vec3(), vec3.vec3(),vec3.vec3(), vec3.vec3()]
    fooNpVec3 = np.array([0.0, 0.0, 0.0], dtype=np.float64)
    foofoo(20)
    allObj2 = muppy.get_objects()

    sum1 = summary.summarize(allObj)
    sum2 = summary.summarize(allObj2)

    summary.print_(summary.get_diff(sum1, sum2))
 
    print("asizeof(mtrx4()) - {}".format(asizeof.asizeof(mtrx4.mtrx4())))
    print("asizeof(vec3()) - {}".format(asizeof.asizeof(vec3.vec3())))
    print("asizeof(npVec3()) - {}".format(asizeof.asizeof(npVec3())))
    print("asizeof(tpVec3()) - {}".format(asizeof.asizeof(tpVec3())))
    print("asizeof(fooAr) - {}".format(asizeof.asizeof(fooAr)))
    print("asizeof(vec3Lst) - {}".format(asizeof.asizeof(vec3Lst)))
    print("asizeof(vFoo) - {}".format(asizeof.asizeof(vFoo)))
    print("asizeof(glm.veс3()) - {}".format(asizeof.asizeof(glm.vec3(0.0))))
    print("asizeof(fooNpVec3) - {}".format(asizeof.asizeof(fooNpVec3)))

if __name__ == "__main__":
    main()