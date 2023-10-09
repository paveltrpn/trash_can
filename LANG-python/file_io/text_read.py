
from common import checkEncoding, getFileSizeBR
import chardet  

def textReadAsBytes(fname: str) -> str:
    """Читаем файл как бинарный и возвращаем строку,
    полученную из считанных байтов"""
    file = open(fname, "rb")
    size = getFileSizeBR(file)
    bytes = file.read(size)
    print(checkEncoding(bytes))
    return bytes.decode("utf-8") 

def textRead(fname: str) -> str:
    with open(fname, "r", encoding="utf8") as f:
            content = f.readlines()
            src = " ".join(content)

    return src