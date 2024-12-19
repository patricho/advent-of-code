use crate::util::file::read_file_lines;
use crate::util::grid::{
    get_grid_point, lines_to_char_grid, move_point, Point, DIRECTIONS, DOWNLEFT, DOWNRIGHT, UPLEFT,
    UPRIGHT,
};
use crate::util::misc::{assert_test, measure, show_results};

static FILE_TEST: &str = "data/2024/04_test.txt";
static FILE_INPUT: &str = "data/2024/04_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 18);
    assert_test(FILE_TEST, part2, 9);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    })
}

fn part1(filename: &str) -> usize {
    let lines = read_file_lines(filename);
    let grid = lines_to_char_grid(&lines);

    grid.iter()
        .enumerate()
        .map(|line_idx| {
            let (y, line) = line_idx;
            return line
                .iter()
                .enumerate()
                .map(|char_idx| {
                    let (x, ch) = char_idx;

                    if *ch != 'X' {
                        return 0;
                    }

                    let pos = Point {
                        y: y as isize,
                        x: x as isize,
                    };

                    return DIRECTIONS
                        .iter()
                        .map(|dir| {
                            let ptx = move_point(&pos, &dir, 0);
                            let ptm = move_point(&pos, &dir, 1);
                            let pta = move_point(&pos, &dir, 2);
                            let pts = move_point(&pos, &dir, 3);

                            let chx = get_grid_point(&grid, &ptx);
                            let chm = get_grid_point(&grid, &ptm);
                            let cha = get_grid_point(&grid, &pta);
                            let chs = get_grid_point(&grid, &pts);

                            match (chx, chm, cha, chs) {
                                ('X', 'M', 'A', 'S') => 1,
                                _ => 0,
                            }
                        })
                        .sum::<usize>();
                })
                .sum::<usize>();
        })
        .sum()
}

fn part2(filename: &str) -> usize {
    let lines = read_file_lines(filename);
    let grid = lines_to_char_grid(&lines);

    grid.iter()
        .enumerate()
        .map(|line_idx| {
            let (y, line) = line_idx;
            return line
                .iter()
                .enumerate()
                .map(|char_idx| {
                    let (x, ch) = char_idx;

                    if *ch != 'A' {
                        return 0;
                    }

                    let pos = Point {
                        y: y as isize,
                        x: x as isize,
                    };

                    let p_ul = move_point(&pos, &UPLEFT, 1);
                    let p_ur = move_point(&pos, &UPRIGHT, 1);
                    let p_dl = move_point(&pos, &DOWNLEFT, 1);
                    let p_dr = move_point(&pos, &DOWNRIGHT, 1);

                    let ch_ul = get_grid_point(&grid, &p_ul);
                    let ch_ur = get_grid_point(&grid, &p_ur);
                    let ch_dl = get_grid_point(&grid, &p_dl);
                    let ch_dr = get_grid_point(&grid, &p_dr);

                    match (ch_ul, ch_ur, ch_dl, ch_dr) {
                        ('S', 'M', 'S', 'M') => 1,
                        ('M', 'S', 'M', 'S') => 1,
                        ('M', 'M', 'S', 'S') => 1,
                        ('S', 'S', 'M', 'M') => 1,
                        _ => 0,
                    }
                })
                .sum::<usize>();
        })
        .sum()
}
