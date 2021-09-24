
@echo OFF

echo TEMPERATURE CONVERTER GCC

set GCCDIR="E:\TDM-GCC-64\bin"
set LIBDIR="E:\TDM-GCC-64\Lib"
set INCDIR="E:\TDM-GCC-64\Include"
set WORKDIR=%~dp0

set SOURCES=main.c
set OFILES=main.o
set OUTPUT=main.exe
set LIBS=
set CFLAGS=-std=c11 -c
set LFLAGS=

@echo ON
%GCCDIR%\gcc %CFLAGS% %SOURCES%
%GCCDIR%\gcc %OFILES% %LIBS% %LFLAGS% -o %OUTPUT%
