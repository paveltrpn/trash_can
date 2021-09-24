
using LinearAlgebra

function vk_quest_solver()
    A::Matrix{Float64} = [1 5 2; 
                          8 1 7; 
                          1 2 1]

    B::Matrix{Float64} = [43 25 11]
    
    d = det(A)

    if d == 0
        println("determinant of A is zero! d = $d")
        return 
    end
    
    #X = B*(A^(-1))
    X = B*(inv(A))

    println("unknown element vector is: \n X1 = $(X[1]) \n X2 = $(X[2]) \n X3 = $(X[3])")
end

function main()
    vk_quest_solver()
end

#-----------main()----------
main()
#---------------------------