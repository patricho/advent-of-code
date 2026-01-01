use crate::util::file::{read_file_lines, read_file_string};
use crate::util::grid::{
    get_grid_point, grid_contains, lines_to_char_grid, move_to_new_point, Point, DIRECTIONS,
};
use crate::util::misc::{assert_test, measure, show_results, split_lines_to_vec, to_int};

const FILE_TEST: &str = "../inputs/2025/04-test.txt";
const FILE_INPUT: &str = "../inputs/2025/04-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 13);
    assert_test(FILE_TEST, part2, 43);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let mut grid = parse_grid(filename);
    count_neighbors(&mut grid)
}

fn part2(filename: &str) -> isize {
    let mut grid = parse_grid(filename);

    let mut result = 0;

    loop {
        let neighbors = count_neighbors(&mut grid);

        result += neighbors;

        if neighbors == 0 {
            break;
        }

        remove_neighbors(&mut grid);
    }

    result
}

fn remove_neighbors(grid: &mut Vec<Vec<char>>) {
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] != 'X' {
                continue;
            }

            grid[y][x] = '.';
        }
    }
}

fn count_neighbors(grid: &mut Vec<Vec<char>>) -> isize {
    let mut result = 0;

    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == '.' {
                continue;
            }

            let point = Point::new(y as isize, x as isize);

            let mut neighbors = 0;

            DIRECTIONS.iter().for_each(|dir| {
                let new_point = move_to_new_point(&point, &dir);

                if get_grid_point(&grid, &new_point) == '@'
                    || get_grid_point(&grid, &new_point) == 'X'
                {
                    neighbors += 1;
                }
            });

            if neighbors < 4 {
                grid[y][x] = 'X';
                result += 1;
            }
        }
    }

    result
}

fn parse_grid(filename: &str) -> Vec<Vec<char>> {
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);
    grid
}
