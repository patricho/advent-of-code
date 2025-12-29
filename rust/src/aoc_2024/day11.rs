use std::collections::HashMap;

use crate::util::{
    file::read_file_string,
    misc::{assert_test, digit_count, measure, show_results, split_digits, split_spaces_to_ints},
};

static FILE_TEST: &str = "data/2024/11_test.txt";
static FILE_INPUT: &str = "data/2024/11_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 55312);
    assert_test(FILE_TEST, part2, 65601038650482);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });

    assert_test(FILE_INPUT, part1, 213625);
    assert_test(FILE_INPUT, part2, 252442982856820);
}

fn part1(filename: &str) -> usize {
    solve(filename, 25)
}

fn part2(filename: &str) -> usize {
    solve(filename, 75)
}

fn solve(filename: &str, blinks: isize) -> usize {
    let input = read_file_string(filename);
    let inputs: Vec<isize> = split_spaces_to_ints(&input);

    let mut stones: HashMap<usize, usize> = HashMap::new();

    inputs.iter().for_each(|i| {
        *stones.entry(*i as usize).or_default() += 1;
    });

    for _ in 0..blinks {
        stones = blink(&stones);
    }

    stones.iter().map(|(_, count)| count).sum()
}

fn blink(stones: &HashMap<usize, usize>) -> HashMap<usize, usize> {
    let mut new_stones: HashMap<usize, usize> = HashMap::new();

    stones.iter().for_each(|(&stone, &count)| {
        if stone == 0 {
            *new_stones.entry(1).or_default() += count;
        } else if digit_count(stone) % 2 == 0 {
            let (l, r) = split_digits(stone as usize);
            *new_stones.entry(l).or_default() += count;
            *new_stones.entry(r).or_default() += count;
        } else {
            *new_stones.entry(stone * 2024).or_default() += count;
        }
    });

    new_stones
}
