
function is_prime(n::Int64)::Bool
    sqrt_n::Int64 = ceil(Int64,sqrt(n))

    if n == 1 
        return false
    end

    for i = 2:sqrt_n
        if (n%i) == 0
            return false
        end
    end

    return true
end

function main()
    println("is_prime - ", is_prime(3))
end

main()