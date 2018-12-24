#!/usr/bin/octave -qf

function Mendeleev_Klapeiron()
	R = 8.314;

	xT = linspace(237,500,50);
	yV = linspace(1,10,50);
	
	[T, V] = meshgrid(xT, yV);
	P = (T .* R) ./ V;
			
	surf(T, V, P);
endfunction

function VanDerVaals()
	R = 8.314;
	a = 3.64; 
	b = 42.67;

	xT = linspace(243,343,100);
	yV = linspace(0.01,0.1,100);
	
	[T, V] = meshgrid(xT, yV);
	P = ((T .* R) ./ (V .- b)) - (a ./ (V .^ 2));

	surf(T, V, P);
endfunction

%Mendeleev_Klapeiron();
VanDerVaals();

pause();
