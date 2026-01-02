use crate::util::file::read_file_lines;
use crate::util::grid::{grid_contains, lines_to_char_grid, Point};
use crate::util::misc::{assert_test, debug_print, measure, show_results};

const FILE_TEST: &str = "../inputs/2025/07-test.txt";
const FILE_INPUT: &str = "../inputs/2025/07-input.txt";

struct Equation {
    numbers: Vec<isize>,
    sign: char,
}

pub fn main() {
    assert_test(FILE_TEST, part1, 21);
    assert_test(FILE_TEST, part2, 40);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let (splits, _) = traverse(filename);
    splits
}

fn part2(filename: &str) -> isize {
    let (_, combinations) = traverse(filename);
    combinations
}

fn parse_grid(filename: &str) -> Vec<Vec<char>> {
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);
    grid
}

fn traverse(filename: &str) -> (isize, isize) {
    let grid = parse_grid(filename);

    let mut splits = 0;

    let mut beams = vec![0; grid[0].len()];

    let start_idx = grid[0].iter().position(|&r| r == 'S').unwrap();

    beams[start_idx] = 1;

    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] != '^' || beams[x] == 0 {
                continue;
            }

            splits += 1;

            let pl = Point::new(y as isize, (x - 1) as isize);
            let pr = Point::new(y as isize, (x + 1) as isize);

            if grid_contains(&grid, &pl) {
                beams[x - 1] += beams[x];
            }

            if grid_contains(&grid, &pr) {
                beams[x + 1] += beams[x];
            }

            beams[x] = 0
        }
    }

    let combinations = beams.iter().sum();

    (splits, combinations)
}
