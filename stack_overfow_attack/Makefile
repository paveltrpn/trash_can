CC = g++
CFLAGS = 

.phony: all clean

all: main.o exploit
	$(CC) main.o -o main -z execstack

main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp

exploit: exploit.asm
	nasm exploit.asm -o exploit

clean:
	rm -rf *.o main exploit