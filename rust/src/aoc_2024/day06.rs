use std::collections::{HashMap, HashSet};

use crate::util::{
    file::read_file_lines,
    grid::{lines_to_char_grid, move_to_new_point, print_grid, Point, DOWN, LEFT, RIGHT, UP},
    misc::{assert_test, measure, show_results},
};

static FILE_TEST: &str = "data/2024/06_test.txt";
static FILE_INPUT: &str = "data/2024/06_input.txt";

pub fn main() {
    /* // https://togglebit.io/posts/rust-bitwise/
    // https://www.tutorialspoint.com/rust/rust_bitwise_operators.htm

    let x: u64 = 211;
    let y = 111;
    let d = 3;

    let xb = x << 16;
    let yb = y << 8;
    let comb = xb | yb | d;

    println!("x {x} in binary: : {x:b}");
    println!("y {y} in binary: : {y:b}");
    println!("d {d} in binary: : {d:b}");
    println!("x<<16 {xb} in binary: : {xb:b}");
    println!("y<<8 {yb} in binary: : {yb:b}");
    println!("d {d} in binary: : {d:b}");
    println!("comb {comb} in binary:   : {comb:b}");

    let xagain = comb >> 16;
    let yagain = (comb & (255 << 8)) >> 8;
    let dagain = comb & (255);

    println!("xagain {xagain} in binary: : {xagain:b}");
    println!("yagain {yagain} in binary: : {yagain:b}");
    println!("dagain {dagain} in binary: : {dagain:b}"); */

    assert_test(FILE_TEST, part1, 41);
    assert_test(FILE_TEST, part2, 6);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });

    // assert_test(FILE_INPUT, part1, 4988);
    // assert_test(FILE_INPUT, part2, 1697);
}

fn init_grid(
    filename: &str,
) -> (
    Vec<Vec<char>>,
    Point,
    HashSet<Point>,
    HashSet<(isize, isize)>,
) {
    // read input to grid
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);

    let mut start = Point { x: -1, y: -1 };
    let mut obstacles: HashSet<Point> = HashSet::new();
    let seen: HashSet<(isize, isize)> = HashSet::new();

    // find obstacles and starting position
    for row in 0..grid.len() {
        for col in 0..grid[row].len() {
            if grid[row][col] == '^' {
                start.x = col as isize;
                start.y = row as isize;
            } else if grid[row][col] == '#' {
                obstacles.insert(Point {
                    x: col as isize,
                    y: row as isize,
                });
            }
        }
    }

    return (grid, start, obstacles, seen);
}

const DIRS: [Point; 4] = [UP, RIGHT, DOWN, LEFT];

fn part1(filename: &str) -> usize {
    // directions to use, starting by moving up
    let mut dir_idx: usize = 0;

    let (grid, mut pos, obstacles, mut seen) = init_grid(filename);

    // record starting position
    seen.insert((pos.y, pos.x));

    loop {
        // try move
        let newpos = move_to_new_point(&pos, &DIRS[dir_idx]);

        if newpos.x < 0
            || newpos.x >= grid[0].len() as isize
            || newpos.y < 0
            || newpos.y >= grid.len() as isize
        {
            // out of bounds, now we're done
            break;
        }

        if obstacles.contains(&newpos) {
            // turn 90 degrees
            dir_idx = (dir_idx + 1) % 4;
        } else {
            // actually move
            pos = newpos;
        }

        // record seen
        seen.insert((pos.y, pos.x));
    }

    // return visited positions
    seen.len()
}

fn part2(filename: &str) -> usize {
    let (grid, start, obstacles, _) = init_grid(filename);

    let mut total_steps = 0;

    let mut seen = HashSet::new();

    let total_cycles = grid
        .iter()
        .enumerate()
        .map(|(row_no, row_chars)| {
            row_chars
                .iter()
                .enumerate()
                .filter(|(col_no, _)| {
                    let newobst = Point {
                        y: row_no as isize,
                        x: *col_no as isize,
                    };

                    let (cycle, steps) =
                        detect_cycle(&grid, &obstacles, newobst, &start, &mut seen);

                    total_steps += steps;

                    cycle
                })
                .count()
        })
        .sum::<usize>();

    println!("total_steps: {total_steps}");

    total_cycles
}

fn detect_cycle(
    grid: &Vec<Vec<char>>,
    obstacles: &HashSet<Point>,
    newobstacle: Point,
    start: &Point,
    seen: &mut HashSet<usize>,
) -> (bool, i32) {
    let mut dir_idx: usize = 0;

    let mut steps = 0;

    let mut pos = start.clone();

    seen.clear();
    let posx: usize = start.x as usize;
    let posy: usize = start.y as usize;
    let posd: usize = 0;
    seen.insert((posy << 16) | (posx << 8) | posd);

    loop {
        steps += 1;

        // try move
        let newpos = move_to_new_point(&pos, &DIRS[dir_idx]);

        if newpos.x < 0
            || newpos.x >= grid[0].len() as isize
            || newpos.y < 0
            || newpos.y >= grid.len() as isize
        {
            // out of bounds, no cycle
            return (false, steps);
        }

        if obstacles.contains(&newpos) || newobstacle == newpos {
            // turn 90 degrees
            dir_idx = (dir_idx + 1) % 4;
        } else {
            // actually move
            pos = newpos;
        }

        let pos_key: usize =
            ((pos.y as usize) << 16) | ((pos.x as usize) << 8) | (dir_idx as usize);

        if seen.contains(&pos_key) {
            return (true, steps);
        }

        // record seen
        seen.insert(pos_key);
    }
}

// mark_seen marks all visited positions on the grid, and prints the grid to stdout
/* fn mark_seen(grid: &mut Vec<Vec<char>>, seen: &HashMap<Point, Vec<usize>>) {
    seen.iter().for_each(|f| {
        grid[f.0.y as usize][f.0.x as usize] = 'X';
    });

    print_grid(&grid);
} */
