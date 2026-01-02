use regress::Regex;

use crate::util::file::read_file_string;
use crate::util::misc::{assert_test, measure, show_results, to_int};

const FILE_TEST: &str = "../inputs/2025/02-test.txt";
const FILE_INPUT: &str = "../inputs/2025/02-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 1227775554);
    assert_test(FILE_TEST, part2, 4174379265);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let twice_regex = Regex::new(r"^(\d+)\1$").unwrap();
    check(filename, twice_regex)
}

fn part2(filename: &str) -> isize {
    let multiple_regex = Regex::new(r"^(\d+)\1+$").unwrap();
    check(filename, multiple_regex)
}

fn check(filename: &str, re: Regex) -> isize {
    let input = read_file_string(filename);
    input
        .split(",")
        .map(|s| check_range(&s, &re))
        .sum::<isize>()
}

fn check_range(range_string: &str, re: &Regex) -> isize {
    let (start_string, end_string) = range_string.split_once("-").unwrap();
    let start = to_int(start_string);
    let end = to_int(end_string);

    (start..=end)
        .map(|n| (n, n.to_string()))
        .filter(|(_, n_str)| re.find(&n_str).is_some())
        .map(|(n, _)| n)
        .sum::<isize>()
}
