
using Base 
using Core
using Random

const f_dom = 10 #in julia 1.0 throw - "type declaretion global vars is not yet supported"

#makes a array with y = tgAf*x + b whith small deviation in (-1,1)
function make_rnd_points(n::Int)::Array{Float32,1}
    tgAf::Float32 = 0.0
    b::Float32    = 5.0
    ret::Array{Float32,1} = zeros(n)

    for x = 1:n
        ret[x] = (tgAf*x + b) + ((rand(Float32)*2-1)/10)
        #println(ret[x]) #test output
    end
    
    return ret
end

function main()
    f_val = make_rnd_points(f_dom)
end

#-----------main()----------
main()
#---------------------------