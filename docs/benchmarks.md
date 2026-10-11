# Benchmarks

Performance and allocation baselines for `go-sentinel-helper`.

The library's core promise is that the sentinel/companion pattern costs
essentially nothing: value-type sentinels stay on the stack with no
allocation, and only the *inherently* heap-shaped operations (the protobuf
`*wrapperspb` pointer wrappers, and `String()` which calls `fmt`) pay.

These numbers exist to keep that promise honest. They are not a leaderboard —
they are a regression tripwire.

## Running

```sh
make bench            # all packages, count=3, benchtime=1s, with -benchmem
make bench-allocs     # quick single-pass run
make bench-record     # writes bench/baseline.txt
make bench-compare    # benchstat diff against the recorded baseline
```

Knobs (override on the command line):

| Variable     | Default          | Meaning                                   |
|--------------|------------------|-------------------------------------------|
| `BENCHTIME`  | `1s`             | `-benchtime` per benchmark                |
| `BENCHCOUNT` | `3`              | `-count`; >1 lets benchstat compute spread |
| `BENCH`      | `.`              | `-bench` regex, e.g. `BENCH='BenchmarkBytes'` |
| `BENCHDIR`   | `bench`          | where recorded output lives               |
| `BENCH_FILE` | `bench/baseline.txt` | recorded baseline consumed by `bench-compare` |

`make bench-compare` needs benchstat:

```sh
go install golang.org/x/perf/cmd/benchstat@latest
```

## Environment

Recorded with `-benchtime=200ms -count=3`, `go1.25.6`, `GOARCH=amd64`.

| Field     | Value                                                  |
|-----------|--------------------------------------------------------|
| CPU       | Intel(R) Core(TM) Ultra 9 185H (22 logical threads seen) |
| OS/Arch   | linux/amd64                                            |
| Benchtime | 200ms                                                  |
| Count     | 3                                                      |

Numbers below are the **median of three runs**, rounded. Absolute ns/op will
shift on other hardware; the *allocs/op* and *relative* comparisons are the
durable signal.

## Headline result: the pattern is allocation-free where it should be

| Package         | Value representation            | Non-formatting ops allocs |
|-----------------|----------------------------------|---------------------------|
| `intutils`      | `int` alias, `math.MinInt` sentinel | **0** |
| `floatutils`    | `float32`/`float64`, NaN sentinel   | **0** |
| `stringutils`   | `string` alias, `\x00unspecified`   | **0** |
| `boolutils`     | tri-state struct (Pattern 2-D)      | **0** |
| `protobufwrapper` | `*wrapperspb.*` pointer           | 0 for predicates, **1** for merge/copy |

Every value-type operation — `IsSpecified`, `TakeOrElse`, `Merge`, `Same`,
`SemanticEqual`, `Equal`, `Copy` — reports `0 B/op, 0 allocs/op` across all
three runs. The `boolutils` tri-state struct in particular confirms Pattern 2-D:
even the struct-based representation stays register-resident.

## `protobufwrapper` — pointer wrappers

Predicates are free; anything that produces a new value allocates the wrapper.

| Benchmark                            | ns/op | B/op | allocs/op |
|--------------------------------------|------:|-----:|----------:|
| `Int32Value/IsSpecified`             | 4.29  | 0    | 0         |
| `Int32Value/IsSpecified/Unspecified` | 4.45  | 0    | 0         |
| `Int32Value/TakeOrElse/Hit`          | 3.64  | 0    | 0         |
| `Int32Value/TakeOrElse/Miss`         | 3.81  | 0    | 0         |
| `Int32Value/Same/Pointer`            | 4.08  | 0    | 0         |
| `Int32Value/SemanticEqual`           | 4.86  | 0    | 0         |
| `Int32Value/Coalesce/Nil`            | 4.27  | 0    | 0         |
| `Int32Value/Merge/BothSpecified`     | 35.3  | 48   | 1         |
| `Int32Value/Copy`                    | 35.6  | 48   | 1         |
| `Int32Value/String/Specified`        | 115   | 16   | 1         |

The same shape holds for `Int64Value`, `UInt32Value`, `UInt64Value`,
`FloatValue`, `DoubleValue`, `BoolValue` and `StringValue`: predicate/identity
ops are 0 allocs, `Merge`/`Copy` cost exactly one allocation (the fresh
wrapper), and `String` costs one small allocation for the formatted text.

### `BytesValue` — linear in payload

`BytesValue` is the only wrapper that copies a payload slice, so `Merge` and
`Copy` are O(n). Throughput stays flat (~2.1–3.0 GB/s), which is what you want —
it means the copy is memory-bandwidth bound and nothing pathological is
happening per byte.

| Payload | Op              | ns/op   | throughput | B/op      | allocs/op |
|---------|-----------------|--------:|-----------:|----------:|----------:|
| 16 B    | `Merge`         | 80.3    | 199 MB/s   | 80        | 2         |
| 16 B    | `Copy`          | 70.7    | 226 MB/s   | 80        | 2         |
| 1 KiB   | `Merge`         | 484.5   | 2114 MB/s  | 1088      | 2         |
| 1 KiB   | `Copy`          | 462.0   | 2217 MB/s  | 1088      | 2         |
| 1 MiB   | `Merge`         | 379570  | 2763 MB/s  | 1048644   | 2         |
| 1 MiB   | `Copy`          | 498629  | 2103 MB/s  | 1048645   | 2         |

`B/op` tracks `len(payload) + header`, and `allocs/op` stays at 2 regardless of
size (the new `[]byte` plus the wrapper). Callers merging large byte payloads in
a hot loop should expect one heap-sized copy per merge — that is the documented
semantics, not a leak.

## The cost of the companion indirection

`companions.NewScalarCompanion` stores each operation in a `func` field, so a
call through the companion is an indirect call through a struct field and
**cannot be inlined**. The direct unexported helper does inline.

| Path                          | ns/op  | allocs/op |
|-------------------------------|-------:|----------:|
| `CompanionDispatch/DirectHelper` | 0.25 | 0         |
| `CompanionDispatch/ViaCompanion` | 4.30 | 0         |

That is roughly a **17× penalty** on trivial predicates. It is the deliberate
trade-off for a single, uniform public API surface (`XCompanion`) and for the
contract that only the companion is exported. The absolute cost is ~4 ns per
call, which is noise for typical singular-value handling and only becomes
visible if a companion op is called in a tight inner loop over a large
collection. If that ever matters, call the package helper directly inside the
package — do not rewrite the companion pattern.

## Formatting is the expensive path

`String()` routes through `fmt.Sprintf`, which dominates any benchmark that
touches it:

| Benchmark                              | ns/op | B/op | allocs/op |
|----------------------------------------|------:|-----:|----------:|
| `IntValue/String/Specified`            | 118.7 | 16   | 1         |
| `IntValueFormatting/CompanionString`   | 113.2 | 16   | 1         |
| `IntValueFormatting/SprintfBaseline`   | 106.6 | 16   | 1         |
| `Float64Companion/String/Specified`    | 253.3 | 32   | 3         |
| `Float64Companion/String/Unspecified`  | 151.7 | 32   | 2         |
| `FloatFormatting/SprintfBaseline`      | 188.5 | 16   | 1         |
| `BooleanValueFormatting/SprintfV`      | 225.7 | 40   | 2         |

The float `String()` allocates more than the raw `Sprintf` baseline (3 vs 1)
because it formats a `%T` type name in addition to the value, and because the
`any` boxing in the generic path escapes. This is cosmetic/debug output; avoid
it in hot paths.

Note that `boolutils` is the exception that proves the rule — its `String()`
hand-rolls the switch and needs **0 allocs** for both `true` and `Unspecified`,
because it never calls `fmt`. If formatting ever does show up in a profile,
that is the pattern to copy.

## Interpreting the numbers

- **0 allocs/op** is the invariant to protect. A change that introduces an
  allocation into a value-type op is a real regression, not noise.
- **B/op = 48 for wrapper `Merge`/`Copy`** is the `wrapperspb.*` struct plus
  header; it should stay constant.
- **ns/op** moves with hardware, GC state and CPU frequency. Trust
  `bench-compare` deltas over absolute values, and re-record the baseline only
  when the recorded numbers are deliberately superseded.
- The 22 suffix (`-22`) is `GOMAXPROCS`; benchmarks are single-goroutine but
  the suffix is part of the name, so keep it in mind when grepping.

## Recorded baseline

The raw, unedited output backing every table above lives in
[`bench/baseline.txt`](../bench/baseline.txt) (392 lines, `-count=3`). Regenerate
with `make bench-record`.
