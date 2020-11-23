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

value fn = f() ->
    val i = i+1
end

value isTrue = f() ->
    i < 2
end

print(string(while(isTrue, fn)))