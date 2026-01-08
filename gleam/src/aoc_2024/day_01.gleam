import gleam/dict.{get}
import gleam/int.{absolute_value as abs}
import gleam/list.{length, map, unzip, zip}
import gleam/result.{unwrap}
import gleeunit/should
import shared/utils.{
  dbg, file_lines, sort_ints, split_string_once as split_string, sum_ints,
  to_int,
}

const file_path = "./src/aoc_2024/data/"

pub fn main() {
  part1("01_test.txt") |> should.equal(Ok(11))
  part2("01_test.txt") |> should.equal(Ok(31))

  dbg(part1("01_input.txt"))
  dbg(part2("01_input.txt"))
}

pub fn part1(file) {
  let #(left, right) = get_todays_lists(file_path, file)

  let left = sort_ints(left)
  let right = sort_ints(right)

  // wrong day :D
  // |> list.filter(fn(diff) { diff >= 1 && diff <= 3 })

  zip(left, right)
  |> map(fn(tup) { abs(tup.1 - tup.0) })
  |> sum_ints
}

pub fn part2(file) {
  let #(left, right) = get_todays_lists(file_path, file)

  let right_counts =
    right
    |> list.group(fn(r) { r })

  left
  |> map(fn(l) {
    let rc = length(unwrap(get(right_counts, l), []))
    l * rc
  })
  |> sum_ints
}

fn get_todays_lists(file_path, file) {
  file_lines(file_path <> file)
  |> map(fn(line) {
    let #(l, r) = split_string(line, "   ")
    #(to_int(l), to_int(r))
  })
  |> unzip
}
