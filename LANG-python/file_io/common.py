
from io import BufferedReader
from typing import Any
import chardet  

def getFileSize(fname: str) -> int:
    file = open(fname, "rb")

    file.seek(0,2)  # move the cursor to the end of the file
                    # 0 - абсолютная позиция в файле
                    # 1 - позиция относительно нстоящего положения 
                    # 2 - позиция относительно конца файла
    size = file.tell()
    file.close()

    return size

# Берём размер файла из объекта Bffered reader
def getFileSizeBR(file: BufferedReader) -> int:
    pos = file.tell() # Запонили позицию в файле

    file.seek(0)

    file.seek(0,2)
    size = file.tell()

    file.seek(pos)

    return size

def checkEncoding(rawdata: bytes):
    result = chardet.detect(rawdata)
    return result['encoding']
