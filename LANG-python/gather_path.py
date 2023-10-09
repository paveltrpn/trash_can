from ast import And
import contextlib
import os
from typing import Any


# FILTERED = [".jpg", ".py", ".png", ".css"]

# Этот класс работает так же как и функция с декоратором менеджера контекста
# ниже - chdir(), только использут методы __enter__() и __exit__() чтобы 
# сделать класс с менеджером контекста
class chdir_c:
    def __init__(self, wd: str) -> None:
        self.workDir = wd

    def __enter__(self) -> None:
        self.this_dir = os.getcwd()
        os.chdir(self.workDir)

    def __exit__(self, exc_type, exc_value, traceback) -> None:
        os.chdir(self.this_dir)


# этот декоратор позволяет превратить функцию-генератор в простой диспетчер контекста
@contextlib.contextmanager
def chdir(path: str) -> Any:
    """
    Сначала переходим по заданному пути.
    В конце возвращаемся в исходную папку.
    """
    this_dir = os.getcwd()
    os.chdir(path)
    try:
        yield
    finally:
        os.chdir(this_dir)


def gather_paths() -> None:
    for root, _, files in os.walk('.'):
        for fname in files:
            # Удаляет из вывода файлы с соответствующим рвсширением
            # if os.path.splitext(fname)[1] in FILTERED:
                # continue
            path = os.path.join(root, fname)
            if path.startswith('.'):
                path = path[1:]
                print(path)


if __name__ == '__main__':
    # with chdir("/home/paveltrpn/code/sketches/LANG-python"):
        # gather_paths()

    with chdir_c("/home/paveltrpn/code/sketches/LANG-python") as _:
        gather_paths()

    # input('Press return to continue.')
