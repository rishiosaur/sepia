# A standard factorial implementation.

val factorial = f(n) ->
    if (n < 0) ->
        return -1;
    end
    
    else ->
        if (n == 0) ->
            return 1;
        end

        else ->
            return (n * factorial(n - 1));
        end
    end
end

val n = 16

factorial(n)