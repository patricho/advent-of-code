use crate::util::{
    file::read_file_lines,
    misc::{assert_test, measure, show_results, split_spaces_to_ints, to_int},
};

static FILE_TEST: &str = "data/2024/07_test.txt";
static FILE_INPUT: &str = "data/2024/07_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 3749);
    assert_test(FILE_TEST, part2, 11387);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> usize {
    solve(filename, 1)
}

fn part2(filename: &str) -> usize {
    solve(filename, 2)
}

fn solve(filename: &str, part: usize) -> usize {
    read_file_lines(filename)
        .into_iter()
        .map(|line| {
            let (target_str, nums_str) = line.split_once(": ").unwrap();
            let target = to_int(target_str);
            let nums = split_spaces_to_ints(nums_str);

            let valid = match part {
                1 => is_valid_part1(&target, &nums),
                _ => is_valid_part2(&target, &nums),
            };

            match valid {
                true => target as usize,
                false => 0,
            }
        })
        .sum::<usize>()
}

fn is_valid_part1(target: &isize, nums: &[isize]) -> bool {
    match nums.len() {
        0 => false,
        1 => nums[0] == *target,
        _ => {
            let (first, second, rest) = (nums[0], nums[1], nums[2..].to_vec());
            let (added, muld) = (first + second, first * second);

            let mut rest1 = rest.clone();
            rest1.insert(0, added);

            let mut rest2 = rest.clone();
            rest2.insert(0, muld);

            return is_valid_part1(&target, &rest1) || is_valid_part1(&target, &rest2);
        }
    }
}

fn is_valid_part2(target: &isize, nums: &Vec<isize>) -> bool {
    match nums.len() {
        0 => false,
        1 => nums[0] == *target,
        _ => {
            let (first, second, rest) = (nums[0], nums[1], nums[2..].to_vec());
            let (added, muld, concatd) = (
                first + second,
                first * second,
                format!("{first}{second}")
                    .parse::<isize>()
                    .unwrap_or_default(),
            );

            let mut rest1 = rest.clone();
            rest1.insert(0, added);

            let mut rest2 = rest.clone();
            rest2.insert(0, muld);

            let mut rest3 = rest.clone();
            rest3.insert(0, concatd);

            return is_valid_part2(&target, &rest1)
                || is_valid_part2(&target, &rest2)
                || is_valid_part2(&target, &rest3);
        }
    }
}
