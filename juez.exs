[expected, reported] = System.argv
a = File.stream!(expected, [], 1024)
b = File.stream!(reported, [], 1024)
ok = Stream.zip(a, b) |> Stream.map( fn({x, y}) -> x == y end ) |> Enum.all?
if !ok do
    exit(1)
end