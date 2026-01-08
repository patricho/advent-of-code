import gleam/io
import gleam/list.{map}
import shared/utils.{
  file_lines, get_first, get_last, keep_digits, split_chars, sum_ints, to_int,
}

pub fn main() {
  file_lines("./src/aoc_2023/data/01_input.txt")
  |> map(fn(line) { split_chars(line) |> keep_digits })
  |> map(fn(chars) { to_int(get_first(chars) <> get_last(chars)) })
  |> sum_ints
  |> io.debug
}
