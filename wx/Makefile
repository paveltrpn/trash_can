CC		:= g++
CFLAGS	:= -Wall -O2 -std=c++17
CC_VERSION := $(shell $(CC) --version | grep g++)

DEVHOME	:= $(HOME)/code

PRJ_LIST = wx

$(info $(CC_VERSION))

.PHONY: $(PRJ_LIST) compile build clean all

all: compile build

compile: main.o 

build:
	$(CC) main.o `wx-config --libs` -o main

main.o: main.cpp
	$(CC) -c main.cpp `wx-config --cxxflags` -o main.o

clean:
	rm -rf *.o main