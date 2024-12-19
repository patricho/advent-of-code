use crate::util::file::read_file_string;
use crate::util::misc::{assert_test, measure, show_results, to_int};
use regex::Regex;

static FILE_TEST: &str = "data/2024/03_test.txt";
static FILE_INPUT: &str = "data/2024/03_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 161);
    assert_test(FILE_TEST, part2, 48);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    })
}

fn part1(filename: &str) -> usize {
    let file = read_file_string(filename);
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();

    re.captures_iter(&file)
        .map(|m| (to_int(&m[1]) * to_int(&m[2])) as usize)
        .sum()
}

fn part2(filename: &str) -> usize {
    let file = read_file_string(filename);
    let re = Regex::new(r"(mul\((\d+),(\d+)\))|do\(\)|don't\(\)").unwrap();
    let mut active = true;

    re.captures_iter(&file)
        .map(|instr| match instr[0].to_string() {
            ms if ms.starts_with("mul") && active => {
                (to_int(&instr[2]) * to_int(&instr[3])) as usize
            }
            ms if ms.eq("do()") => {
                active = true;
                0
            }
            ms if ms.eq("don't()") => {
                active = false;
                0
            }
            _ => 0,
        })
        .sum()
}
