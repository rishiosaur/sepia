
value while = f(whileTrue, doFunc) ->
   if (whileTrue()) ->
       doFunc()
       while(whileTrue, doFunc)
    end

   else ->
       true
       end
end

value i = 0

value isTrue = f() ->
    i < 15
end

while(isTrue, f() ->
    update i = i+1

    print(string(i))
end)
