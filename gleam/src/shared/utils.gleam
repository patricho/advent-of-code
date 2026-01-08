import gleam/int
import gleam/io
import gleam/list.{filter, first, last, reduce, sort}
import gleam/result.{unwrap}
import gleam/string.{contains, split, split_once}
import simplifile.{read}

pub fn file_data(filename) {
  case read(filename) {
    Ok(data) -> data
    Error(_) -> ""
  }
}

pub fn file_lines(filename) {
  split(file_data(filename), "\n")
}

pub fn dbg(o) {
  let _ = io.debug(o)
  Nil
}

pub fn split_chars(line) {
  split(line, "")
}

pub fn split_string_once(line, separator) {
  unwrap(split_once(line, separator), #("", ""))
}

pub fn keep_digits(chars) {
  filter(chars, contains("0123456789", _))
}

pub fn get_first(chars) {
  case first(chars) {
    Ok(c) -> c
    Error(_) -> "0"
  }
}

pub fn get_last(chars) {
  case last(chars) {
    Ok(c) -> c
    Error(_) -> "0"
  }
}

pub fn to_int(str) {
  case int.parse(str) {
    Ok(n) -> n
    Error(_) -> 0
  }
}

pub fn sum_ints(list) {
  reduce(list, fn(acc, x) { acc + x })
}

pub fn sort_ints(l) {
  sort(l, int.compare)
}
