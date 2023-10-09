
import glm
from math import cos, sin
import numpy as np

from common import IdRw

def mtrx3GlmFromEuler(yaw: float, pitch: float, roll: float) -> glm.mat3:
    rt = glm.mat3()

    cosy = cos(yaw)
    siny = sin(yaw)
    cosp = cos(pitch)
    sinp = sin(pitch)
    cosr = cos(roll)
    sinr = sin(roll)

    rt[0][0] = cosy*cosr - siny*cosp*sinr
    rt[0][1] = -cosy*sinr - siny*cosp*cosr
    rt[0][2] = siny * sinp

    rt[1][0] = siny*cosr + cosy*cosp*sinr
    rt[1][1] = -siny*sinr + cosy*cosp*cosr
    rt[1][2] = -cosy * sinp

    rt[2][0] = sinp * sinr
    rt[2][1] = sinp * cosr
    rt[2][2] = cosp

    return rt

class CVec3Glm(object):
    __slots__ = ["data", "capacity", "dataSize"]
    
    def __init__(self, arr: list) -> None:
        if not (len(arr)%3 == 0):
            print("CVec3Arrayf(): WARNING! Initial array lenght is not in rate of 3!")

        self.data = glm.array(np.array(arr, dtype=np.float32))
        self.capacity = len(arr)
        self.dataSize = len(arr)*4 # size of array.array("f", [...]) element is 4 bytes 
    
    def applyMtrx3(self, A: glm.mat3) -> None:
        for vec in self.data:
            vec = vec*A
