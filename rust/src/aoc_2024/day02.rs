use crate::util::file::read_file_lines;
use crate::util::misc::{assert_test, measure, show_results, split_spaces_to_ints};

static FILE_TEST: &str = "data/2024/02_test.txt";
static FILE_INPUT: &str = "data/2024/02_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 2);
    assert_test(FILE_TEST, part2, 4);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    })
}

fn part1(filename: &str) -> usize {
    let lines = read_file_lines(filename);
    let diffs = lines_to_diffs(lines);

    let ok_rows = diffs.iter().filter(|ld| linediffs_ok(*ld)).count();

    return ok_rows;
}

fn part2(filename: &str) -> usize {
    let lines = read_file_lines(filename);

    let ok_rows = lines
        .iter()
        .map(|line| split_spaces_to_ints(line))
        .filter(|line| linediffs_ok(&line_diff(line)) || any_line_variant_ok(line))
        .count();

    return ok_rows;
}

fn lines_to_diffs(lines: Vec<String>) -> Vec<Vec<isize>> {
    lines
        .iter()
        .map(|line| split_spaces_to_ints(line))
        .map(|line| line_diff(&line))
        .collect()
}

fn line_diff(lineints: &Vec<isize>) -> Vec<isize> {
    let mut diffs: Vec<isize> = Vec::new();

    for i in 1..lineints.len() {
        let diff = lineints[i] - lineints[i - 1];
        diffs.push(diff);
    }

    return diffs.clone();
}

fn linediffs_ok(line: &Vec<isize>) -> bool {
    let mut inc = false;
    let mut dec = false;
    let mut diffok = true;

    line.iter().for_each(|diff| {
        if diff > &0 {
            inc = true;
        } else if diff < &0 {
            dec = true;
        }

        if diff.abs() < 1 || diff.abs() > 3 {
            diffok = false;
        }
    });

    return (inc || dec) && !(inc && dec) && diffok;
}

fn any_line_variant_ok(line: &Vec<isize>) -> bool {
    // Instead, we can use concat() to concatenate two vector slices into a new vector. This
    // approach allocates memory for the new vector in a single operation, minimizing reallocations
    // and processing each item only once. As a result, it has a time complexity of O(m + n), where
    // m and n are the number of items in each slice, respectively.  Note that concat() creates a
    // new vector from the slices without taking ownership of the original items. Therefore, we need
    // to provide slice references when calling concat().
    // https://github.com/KyraZzz/AoC24/blob/main/day2/day2.md

    (0..line.len()).any(|i| {
        let partial_line = [&line[0..i], &line[i + 1..]].concat();
        let diffs = line_diff(&partial_line);
        linediffs_ok(&diffs)
    })
}
