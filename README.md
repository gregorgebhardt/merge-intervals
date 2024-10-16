# Merge Intervals

The merge interval app takes a space-separated list of intervals and merges overlapping intervals. The resulting list of intervals is written to standard output.
The app is written in Go and uses the module [redblack](https://github.com/gregorgebhardt/redblack) to store the intervals in a red-black tree.

## Requirements to run the code

- [Go 1.23](go.dev/dl) or later

## How to run the code

To build the app, run the following command in the root directory of the project:

```bash
go build -o merge
```

Alternatively, you can run the app directly with the following command:

```bash
go run .
```

## Usage

The app can read the intervals either from standard input or from a file. The intervals are expected to be in the format `[start, end]` where `start` and `end` are integers.

```bash
./merge -f intervals.txt
```
where `intervals.txt` contains the intervals, e.g.,
```txt
[1, 4] [-5, 0] [3, 7] [15, 32] [10, 15] [3, 8]
```

Or pipe the intervals to the app:

```bash
echo "[1, 3] [2, 4] [5, 7] [6, 8]" | ./merge
```

The app will output the merged intervals to standard output:

```txt
[1, 4] [5, 8]
```

## Testing

To run the tests, execute the following command in the root directory of the project:

```bash
go test ./...
```
