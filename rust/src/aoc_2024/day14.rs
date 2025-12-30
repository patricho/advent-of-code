use std::{collections::HashSet, ops::Range};

use crate::util::{
    file::read_file_lines,
    grid::Point,
    misc::{assert_test, measure, show_results, to_int},
};

const FILE_TEST: &str = "../inputs/2024/14-test.txt";
const FILE_INPUT: &str = "../inputs/2024/14-input.txt";

pub fn main() {
    assert_test(FILE_TEST, |filename| part1(filename, 11, 7), 12);
    assert_test(FILE_TEST, |filename| part2(filename, 11, 7), 1);

    measure(|| {
        show_results(FILE_INPUT, 1, |filename| part1(filename, 101, 103));
        show_results(FILE_INPUT, 2, |filename| part2(filename, 101, 103));
    });

    assert_test(FILE_INPUT, |filename| part1(filename, 101, 103), 230435667);
    assert_test(FILE_INPUT, |filename| part2(filename, 101, 103), 7709);
}

fn part2(filename: &str, width: isize, height: isize) -> usize {
    let robots = parse_robots(filename);

    let mut new_robots = Vec::clone(&robots);

    let robots_count = new_robots.len();

    let mut steps = 0;

    loop {
        new_robots = take_robot_steps(&new_robots, 1, width, height);

        steps += 1;

        let mut seen: HashSet<(isize, isize)> = HashSet::new();

        new_robots.iter().for_each(|r| {
            if seen.contains(&(r.p.y, r.p.x)) {
                // No need to check further
                return;
            }
            seen.insert((r.p.y, r.p.x));
        });

        // If no robots are overlapping, perhaps that's it?
        if seen.len() == robots_count {
            break;
        }
    }

    // Print grid, showing the christmas tree
    /*(0..height).clone().for_each(|y| {
        (0..width).clone().for_each(|x| {
            let robots_count = new_robots
                .iter()
                .filter(|&r| r.p.x == x && r.p.y == y)
                .count();

            if robots_count == 0 {
                print!(".");
            } else {
                print!("{robots_count}");
            }
        });
        print!("\n");
    });
    println!("");*/

    steps
}

fn part1(filename: &str, width: isize, height: isize) -> usize {
    let robots = parse_robots(filename);

    let new_robots = take_robot_steps(&robots, 100, width, height);

    let half_width = width / 2;
    let half_height = height / 2;

    // Halves
    let top_half = 0..half_height;
    let bottom_half = half_height + 1..height;
    let left_half = 0..half_width;
    let right_half = half_width + 1..width;

    // Quadrants
    let tl = count_quadrant(&top_half, &left_half, &new_robots);
    let tr = count_quadrant(&top_half, &right_half, &new_robots);
    let bl = count_quadrant(&bottom_half, &left_half, &new_robots);
    let br = count_quadrant(&bottom_half, &right_half, &new_robots);

    tl * tr * bl * br
}

#[derive(Clone, Debug)]
struct Robot {
    p: Point,
    v: Point,
}

fn parse_robots(filename: &str) -> Vec<Robot> {
    let lines = read_file_lines(filename);

    lines
        .iter()
        .map(|line| {
            let parts = line
                .split_whitespace()
                .map(|s| {
                    s.to_string()[2..]
                        .split(",")
                        .map(|s| to_int(&s))
                        .collect::<Vec<_>>()
                })
                .collect::<Vec<_>>();

            Robot {
                p: Point {
                    x: parts[0][0],
                    y: parts[0][1],
                },
                v: Point {
                    x: parts[1][0],
                    y: parts[1][1],
                },
            }
        })
        .collect::<Vec<_>>()
}

fn take_robot_steps(robots: &Vec<Robot>, steps: isize, width: isize, height: isize) -> Vec<Robot> {
    let mut new_robots: Vec<Robot> = Vec::clone(&robots);

    robots.iter().enumerate().for_each(|(idx, r)| {
        let mut new_x = (r.p.x + r.v.x * steps) % width;
        let mut new_y = (r.p.y + r.v.y * steps) % height;

        if new_x < 0 {
            new_x = width + new_x;
        }

        if new_y < 0 {
            new_y = height + new_y;
        }

        new_robots[idx] = Robot {
            p: Point { y: new_y, x: new_x },
            v: r.v,
        };
    });

    new_robots
}

fn count_quadrant(y_range: &Range<isize>, x_range: &Range<isize>, robots: &Vec<Robot>) -> usize {
    let mut sum = 0_usize;

    // TODO: Do a single count for the quadrant instead of stepping through coordinates one by one

    y_range.clone().for_each(|y| {
        x_range.clone().for_each(|x| {
            let robots_count = robots.iter().filter(|&r| r.p.x == x && r.p.y == y).count();
            sum += robots_count;
        });
    });

    sum
}
