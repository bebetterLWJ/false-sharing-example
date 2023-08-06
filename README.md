false-sharing example code
execute  go test --bench=. -benchtime=1000x to get performance difference between pad and nopad

tip: Carefully delete the fields you think are useless, maybe it is used to avoid false sharing