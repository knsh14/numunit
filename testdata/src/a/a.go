package a

var (
    a = 1234 // want "use 1_234"
    b = 12_35 // want "use 1_235"
    c = 124 // OK
    d = 1_236 // OK
    e = int64(1_237) // OK
    f = 12_34567 // want "use 1_234_567"
)

