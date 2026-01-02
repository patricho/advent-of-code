use crate::util::file::read_file_lines;
use crate::util::misc::{assert_test, measure, show_results};

const FILE_TEST: &str = "../inputs/2025/03-test.txt";
const FILE_INPUT: &str = "../inputs/2025/03-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 357);
    assert_test(FILE_TEST, part2, 3121910778619);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    check_file(filename, 2)
}

fn part2(filename: &str) -> isize {
    check_file(filename, 12)
}

fn check_file(filename: &str, target_len: usize) -> isize {
    read_file_lines(filename)
        .iter()
        .map(|line| check_range(&line, target_len))
        .sum::<isize>()
}

fn check_range(row: &str, target_len: usize) -> isize {
    let chars: Vec<char> = row.chars().collect();
    let mut start_idx: usize = 0;
    let mut result = String::new();
    let mut max_idx = chars.len() - target_len;

    loop {
        let mut pick = '0';

        for idx in start_idx..=max_idx {
            let candidate = chars[idx];

            if candidate > pick {
                pick = candidate;
                start_idx = idx + 1;
            }
        }

        result.push(pick);

        if result.len() >= target_len {
            break;
        }

        max_idx += 1;
    }

    result.parse().unwrap_or(0)
}
