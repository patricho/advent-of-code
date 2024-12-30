use std::collections::{HashMap, HashSet};

use crate::util::{
    file::read_file_lines,
    grid::{get_grid_point, grid_contains, lines_to_char_grid, Point},
    misc::{assert_test, measure, show_results},
};

static FILE_TEST: &str = "data/2024/08_test.txt";
static FILE_INPUT: &str = "data/2024/08_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 14);
    assert_test(FILE_TEST, part2, 34);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> usize {
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);
    let frequencies = find_frequencies(&grid);
    let mut antinodes: HashSet<Point> = HashSet::new();

    frequencies
        .into_iter()
        .for_each(|(_, points)| {
            for p1 in &points {
                for p2 in &points {
                    if p1 == p2 {
                        continue;
                    }

                    let delta = *p1 - *p2;
                    let anti = *p1 + delta;

                    if !grid_contains(&grid, &anti) {
                        continue;
                    }

                    antinodes.insert(anti);
                }
            }
        });

    antinodes.len()
}

fn part2(filename: &str) -> usize {
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);
    let frequencies = find_frequencies(&grid);
    let mut antinodes: HashSet<Point> = HashSet::new();

    frequencies
        .into_iter()
        .for_each(|(_, points)| {
            for p1 in &points {
                for p2 in &points {
                    if p1 == p2 {
                        continue;
                    }

                    // Antennas also count as antinodes now
                    antinodes.insert(*p1);
                    antinodes.insert(*p2);

                    let mut p1 = p1.clone();
                    let mut p2 = p2.clone();

                    loop {
                        let delta = p1 - p2;
                        let anti = p1 + delta;

                        if !grid_contains(&grid, &anti) {
                            break;
                        }

                        antinodes.insert(anti);

                        p2 = p1.clone();
                        p1 = anti.clone();
                    }
                }
            }
        });

    antinodes.len()
}

/* fn print_antinodes_grid(grid: &Vec<Vec<char>>, antinodes: &HashSet<Point>) {
    grid.iter()
        .enumerate()
        .for_each(|(y, line)| {
            line.iter()
                .enumerate()
                .for_each(|(x, c)| {
                    if antinodes.contains(&Point::new(y as isize, x as isize)) {
                        print!("{}", '#');
                    } else {
                        print!("{}", c);
                    }
                });
            println!("");
        });
} */

fn find_frequencies(grid: &Vec<Vec<char>>) -> HashMap<char, Vec<Point>> {
    let mut frequencies: HashMap<char, Vec<Point>> = HashMap::new();

    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            let p = Point::new(y as isize, x as isize);
            let c = get_grid_point(&grid, &p);
            if c == '.' {
                continue;
            }
            frequencies
                .entry(c)
                .or_insert_with(Vec::new)
                .push(p);
        }
    }

    frequencies
}
