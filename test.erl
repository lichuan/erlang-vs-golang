-module(test).
-export([start/0, work1/1, work2/1, work_busy/0]).

fibo(0) -> 0;
fibo(1) -> 1;
fibo(N) when N > 1 -> fibo(N - 1) + fibo(N - 2).

work1(Cnt) ->
    fibo(20),
    sha256_busy(10000),
    io:format("erlang................work1 cnt: ~p~n", [Cnt]),
    work1(Cnt + 1).

work2(Cnt) ->
    fibo(20),
    sha256_busy(10000),
    io:format("erlang................work2 cnt: ~p~n", [Cnt]),
    work2(Cnt + 1).

sha256_busy(0) -> ok;
sha256_busy(N) ->
    _Bin = crypto:hash(sha256, "This is a short sha256 test string"
      ++ integer_to_list(N)),
    sha256_busy(N - 1).

work_busy() ->
    sha256_busy(10000),
    work_busy().

start() ->
    io:format("start to run~n"),
    spawn(test, work1, [1]),
    spawn(test, work2, [1]),
    [spawn(test, work_busy, []) || _ <- lists:seq(1, 10)],
    timer:sleep(infinity).



