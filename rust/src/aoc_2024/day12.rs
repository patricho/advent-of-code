use std::collections::{HashMap, HashSet};

use crate::util::{
    file::read_file_lines,
    grid::{
        get_grid_point, lines_to_char_grid, move_to_new_point, Point, DOWN, DOWNLEFT, DOWNRIGHT,
        LEFT, ORTHOGOMALS, RIGHT, UP, UPLEFT, UPRIGHT,
    },
    misc::{assert_test, measure, show_results},
};

const FILE_TEST: &str = "../inputs/2024/12-test.txt";
const FILE_TEST_2: &str = "../inputs/2024/12-test2.txt";
const FILE_INPUT: &str = "../inputs/2024/12-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 1930);
    assert_test(FILE_TEST, part2, 1206);
    assert_test(FILE_TEST_2, part1, 140);
    assert_test(FILE_TEST_2, part2, 80);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });

    assert_test(FILE_INPUT, part1, 1456082);
    assert_test(FILE_INPUT, part2, 872382);
}

fn part1(filename: &str) -> usize {
    let (area_perimeter, _) = solve(filename);
    area_perimeter
}

fn part2(filename: &str) -> usize {
    let (_, area_sides) = solve(filename);
    area_sides
}

fn solve(filename: &str) -> (usize, usize) {
    let input = read_file_lines(filename);
    let grid = lines_to_char_grid(&input);

    let mut visited: HashSet<(isize, isize)> = HashSet::new();
    let mut groups: HashMap<(char, isize), Vec<Point>> = HashMap::new();
    let mut group_key = 0_isize;

    // Group adjacent position into sections
    for row in 0..grid.len() {
        for col in 0..grid[row].len() {
            group_key += 1;
            let pos = Point {
                x: col as isize,
                y: row as isize,
            };
            visit_point(pos, group_key, &grid, &mut visited, &mut groups);
        }
    }

    let mut area_perim = 0;
    let mut area_sides = 0;

    // Count area and perimeter for each group
    groups.iter().for_each(|(_, group_points)| {
        let mut group_perim = 0;
        let mut group_sides = 0;

        // Count neighbors and sides for each point in group
        group_points.iter().for_each(|point| {
            let mut neighbors = 0;
            let mut corners = 0;

            ORTHOGOMALS.iter().for_each(|dir| {
                let new_point = move_to_new_point(&point, &dir);
                if group_points.contains(&new_point) {
                    neighbors += 1;
                }
            });

            // For perimiter we can notice that each space adds 4 - number of neighbors space to the perimiter
            let point_perim = 4 - neighbors;

            // For sides in part 2 we can count the corners instead
            let point_left = move_to_new_point(&point, &LEFT);
            let point_right = move_to_new_point(&point, &RIGHT);
            let point_up = move_to_new_point(&point, &UP);
            let point_down = move_to_new_point(&point, &DOWN);
            let point_upleft = move_to_new_point(&point, &UPLEFT);
            let point_upright = move_to_new_point(&point, &UPRIGHT);
            let point_downleft = move_to_new_point(&point, &DOWNLEFT);
            let point_downright = move_to_new_point(&point, &DOWNRIGHT);

            // Outer corners
            if !group_points.contains(&point_left) && !group_points.contains(&point_up) {
                corners += 1;
            }
            if !group_points.contains(&point_right) && !group_points.contains(&point_up) {
                corners += 1;
            }
            if !group_points.contains(&point_left) && !group_points.contains(&point_down) {
                corners += 1;
            }
            if !group_points.contains(&point_right) && !group_points.contains(&point_down) {
                corners += 1;
            }

            // Inner corners
            if group_points.contains(&point_left)
                && group_points.contains(&point_up)
                && !group_points.contains(&point_upleft)
            {
                corners += 1;
            }
            if group_points.contains(&point_right)
                && group_points.contains(&point_up)
                && !group_points.contains(&point_upright)
            {
                corners += 1;
            }
            if group_points.contains(&point_left)
                && group_points.contains(&point_down)
                && !group_points.contains(&point_downleft)
            {
                corners += 1;
            }
            if group_points.contains(&point_right)
                && group_points.contains(&point_down)
                && !group_points.contains(&point_downright)
            {
                corners += 1;
            }

            group_perim += point_perim;
            group_sides += corners;
        });

        area_perim += group_points.len() * group_perim;
        area_sides += group_points.len() * group_sides;
    });

    (area_perim, area_sides)
}

fn visit_point(
    pos: Point,
    group_idx: isize,
    grid: &Vec<Vec<char>>,
    visited: &mut HashSet<(isize, isize)>,
    groups: &mut HashMap<(char, isize), Vec<Point>>,
) {
    if visited.contains(&(pos.y, pos.x)) {
        // Already assigned to a group
        return;
    }

    let group_key = get_grid_point(&grid, &pos);

    // Record position
    visited.insert((pos.y, pos.x));
    groups
        .entry((group_key, group_idx))
        .or_insert_with(Vec::new)
        .push(pos);

    // Try to visit neighbors (if in same set)
    ORTHOGOMALS.iter().for_each(|d| {
        let new_point = move_to_new_point(&pos, &d);
        let new_key = get_grid_point(&grid, &new_point);

        if new_key != group_key {
            return;
        }

        visit_point(new_point, group_idx, grid, visited, groups);
    });
}
