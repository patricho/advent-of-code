use crate::util::file::read_file_lines;
use crate::util::misc::{assert_test, measure, show_results, to_int};

const FILE_TEST: &str = "../inputs/2025/01-test.txt";
const FILE_INPUT: &str = "../inputs/2025/01-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 3);
    assert_test(FILE_TEST, part2, 6);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let mut position = 50;

    get_steps(filename)
        .iter()
        .map(|step| {
            position += step;
            match position % 100 == 0 {
                true => 1,
                false => 0,
            }
        })
        .sum()
}

fn part2(filename: &str) -> isize {
    let mut position = 50;

    get_steps(filename)
        .iter()
        .map(|&step| {
            let prev_position = position;

            position += step;

            let prev_pos_div = ((prev_position as f64) / 100_f64).floor() as isize;
            let pos_div = ((position as f64) / 100_f64).floor() as isize;

            let mut diff = (pos_div - prev_pos_div).abs();

            if step < 0 && position % 100 == 0 {
                diff += 1;
            } else if step < 0 && prev_position % 100 == 0 {
                diff -= 1;
            }

            diff
        })
        .sum()
}

fn get_steps(filename: &str) -> Vec<isize> {
    read_file_lines(filename)
        .iter()
        .map(|line| {
            let mut step_str = line.replace("L", "-");
            step_str = step_str.replace("R", "");
            to_int(&step_str)
        })
        .collect()
}
