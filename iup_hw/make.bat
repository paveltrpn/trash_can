
@echo OFF

echo IUP GUI HELLO WORLD GCC

set GCCDIR="E:\TDM-GCC-64\bin"
set LIBDIR="E:\TDM-GCC-64\Lib"
set INCDIR="E:\TDM-GCC-64\Include"
set IUPDIR=E:\code\iup-3.20_Win64
set WORKDIR=%~dp0

set SOURCES=main.c
set OFILES=main.o
set OUTPUT=main.exe
set LIBS=
set CFLAGS=-std=c11 -c
set LFLAGS=-s -liupimglib -liup

@echo ON

%GCCDIR%\gcc --version
%GCCDIR%\gcc -I%IUPDIR%\include %CFLAGS% %SOURCES% 
%GCCDIR%\gcc -L%IUPDIR% %OFILES% %LIBS% %LFLAGS% -o %OUTPUT%
