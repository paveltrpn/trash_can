
function print_var(z::Int64)
    println("print_var = ", z)
end

function fib(n::Int64)
    prv::Int64, next::Int64, now::Int64 = 0, 1, 1

    if n < 1 
        println("n must be above 1!")
        return
    end

    println("number 0 = ", prv)
    
    for i = 1:n
        println("number $i = $now")
        now = prv + next
        prv = next
        next = now
    end 
end

local y::Int64 = 0

println("Hello World")

fib(14)