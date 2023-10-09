################################
# Windows-специфичный Makefile #
################################
CC       := clang++
CFLAGS   := -Wall -std=c++17 -target x86_64-pc-windows-gnu 
LFLAGS 	 := -lstdc++ -lgdi32 -lUser32

SRCS := src/main.cpp

OBJS := main.o 

#OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: $(OBJS) main

main.o: src/main.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/main.cpp

main: $(OBJS)
	g++ $(OBJS) $(LIBDIR) -o main $(LFLAGS)

clean:
	rm -f *.o main
