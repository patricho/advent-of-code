import gleam/list.{all, filter, filter_map, index_map, map}
import gleam/string.{split}
import gleeunit/should
import shared/utils.{dbg, file_lines, to_int}

const file_path = "./src/aoc_2024/data/"

pub fn main() {
  part1("02_test.txt") |> should.equal(2)
  part2("02_test.txt") |> should.equal(4)

  dbg(part1("02_input.txt"))
  dbg(part2("02_input.txt"))
}

pub fn part1(file) {
  get_todays_numbers(file)
  |> filter(fn(line) { all_diffs_ok(get_diffs(line)) })
  |> list.length
}

pub fn part2(file) {
  get_todays_numbers(file)
  |> filter(fn(line) {
    all_diffs_ok(get_diffs(line)) || any_combination_ok(line)
  })
  |> list.length
}

fn get_todays_numbers(file) {
  file_lines(file_path <> file)
  |> map(fn(line) { split(line, " ") |> map(to_int) })
}

fn get_diffs(line) {
  line
  |> list.window_by_2
  |> map(fn(tup) {
    let #(prev, cur) = tup
    cur - prev
  })
}

fn all_diffs_ok(diffs) {
  let all_dec =
    diffs
    |> all(fn(diff) { diff <= -1 && diff >= -3 })

  let all_inc =
    diffs
    |> all(fn(diff) { diff >= 1 && diff <= 3 })

  { all_dec || all_inc } && !{ all_dec && all_inc }
}

fn any_combination_ok(line) {
  remove_one_at_a_time(line)
  |> list.any(fn(line) { get_diffs(line) |> all_diffs_ok })
}

fn remove_one_at_a_time(line) {
  line
  |> index_map(fn(_, idx) {
    line
    |> index_map(fn(item, i) { #(i != idx, item) })
    |> filter_map(fn(pair) {
      case pair {
        #(True, item) -> Ok(item)
        #(False, _) -> Error(Nil)
      }
    })
  })
}
