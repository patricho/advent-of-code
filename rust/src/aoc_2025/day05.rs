use crate::util::file::read_file_string;
use crate::util::misc::{assert_test, measure, show_results, split_lines_to_vec, to_int};

const FILE_TEST: &str = "../inputs/2025/05-test.txt";
const FILE_INPUT: &str = "../inputs/2025/05-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 3);
    assert_test(FILE_TEST, part2, 14);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let file = read_file_string(filename);
    let (ranges_string, ingredients_string) = file.split_once("\n\n").unwrap();

    let ranges = split_lines_to_vec(ranges_string)
        .iter()
        .map(|line| {
            let (a, b) = line.split_once("-").unwrap();
            (to_int(a), to_int(b))
        })
        .collect::<Vec<_>>();

    split_lines_to_vec(ingredients_string)
        .iter()
        .map(|n| to_int(n))
        .filter(|&i| ranges.iter().any(|&r| i >= r.0 && i <= r.1))
        .collect::<Vec<_>>()
        .len() as isize
}

fn part2(filename: &str) -> isize {
    let file = read_file_string(filename);
    let (ranges_string, _) = file.split_once("\n\n").unwrap();

    let mut ranges = split_lines_to_vec(ranges_string)
        .iter()
        .map(|line| {
            let (a, b) = line.split_once("-").unwrap();
            (to_int(a), to_int(b))
        })
        .collect::<Vec<_>>();

    ranges.sort_by(|a, b| a.0.cmp(&b.0));

    let mut result: isize = 0;
    let mut max_counted: isize = -1;

    for (start, end) in ranges {
        if max_counted >= end {
            // Range is already completely covered
            continue;
        }

        let adjusted_start = if max_counted >= start {
            max_counted + 1
        } else {
            start
        };

        let span = end - adjusted_start + 1;
        result += span;
        max_counted = end;
    }

    result
}
