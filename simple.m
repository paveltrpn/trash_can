#!/usr/bin/octave -qf

function foo();
	x = linspace(0, 2*pi, 100);
	y = sin(x);
	plot(x,y);
endfunction

foo();
pause();
