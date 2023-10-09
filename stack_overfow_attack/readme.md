Пример атаки через переполнение буфера на стеке.    
Взято из https://www.youtube.com/watch?v=BSSuxvY8tLA.   

main.cpp - уязвимый код.    
exploit.asm - полезная нагрузка.    
patch.sh - скрипт для дополнителльного патча скомпилированного эксплойта.

Сборка и запуск:    
$ make - сборка main и exploit  
$ sh patch.sh   
$ ./main "$(< exploit)" 

Опция "-z execstack" - сделать стек исполняемым.