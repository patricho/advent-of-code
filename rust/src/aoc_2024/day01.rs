use crate::util::file::read_file_lines;
use crate::util::misc::{count_values, get_hash_int, measure, split_spaces_to_ints};

const FILE_TEST: &str = "data/2024/01_test.txt";
const FILE_INPUT: &str = "data/2024/01_input.txt";

pub fn main() {
    assert_eq!(part1(FILE_TEST), 11);
    assert_eq!(part2(FILE_TEST), 31);

    measure(|| {
        part1(FILE_INPUT);
        part2(FILE_INPUT);
    });
}

fn part1(filename: &str) -> i32 {
    let (left, right) = get_sorted_lists(filename);

    let result = left.iter().zip(&right).map(|(l, r)| (l - r).abs()).sum();

    println!("file: {filename}, res: {result}");

    result
}

fn part2(filename: &str) -> i32 {
    let (left, right) = get_sorted_lists(filename);
    let counts = count_values(right);

    let result = left.iter().map(|l| l * get_hash_int(&counts, l)).sum();

    println!("file: {filename}, res: {result}");

    result
}

fn get_sorted_lists(filename: &str) -> (Vec<i32>, Vec<i32>) {
    let lines = read_file_lines(filename);

    let (mut left, mut right): (Vec<i32>, Vec<i32>) = lines
        .iter()
        .map(|line| {
            let ints = split_spaces_to_ints(line);
            (ints[0], ints[1])
        })
        .unzip();

    left.sort_unstable();
    right.sort_unstable();

    (left, right)
}
