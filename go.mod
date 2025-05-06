module example/hello

go 1.24.2

replace example.com/greetings => ./greetings

replace example/values => ./values

replace example/loop => ./loop

replace example/array => ./array

replace example/conditions => ./conditions

replace example/pointers => ./pointers

replace example/struct => ./struct

replace example/slices => ./slices

replace example/maps => ./maps

require example/maps v0.0.0-00010101000000-000000000000
