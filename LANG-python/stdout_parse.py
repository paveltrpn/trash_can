
import os
import subprocess

def main() -> None:
    cmd = ["git", "status"]

    # Вызываем комманду, перехватываем stdout, декодируем в utf-8 (обязательно!)
    # и удаляем символы табуляции
    gitStatusMsg = subprocess.run(cmd, capture_output=True).stdout.decode("utf-8").replace("\t", "")

    # Разбиваем строку на список строк по переносу строки
    msgStrList = gitStatusMsg.split("\n")

    for line in msgStrList:
        print(line)

    # Ищем подстроку в списке строк
    sub = "modified:"
    fName = filter(lambda x: sub in x, msgStrList)

    print(list(fName))

if __name__ == "__main__":
    main()