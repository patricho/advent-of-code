#[derive(Debug)]
pub struct Point {
    pub x: isize,
    pub y: isize,
}

pub const UP: Point = Point { y: -1, x: 0 };
pub const DOWN: Point = Point { y: 1, x: 0 };
pub const LEFT: Point = Point { y: 0, x: -1 };
pub const RIGHT: Point = Point { y: 0, x: 1 };
pub const UPLEFT: Point = Point { y: -1, x: -1 };
pub const UPRIGHT: Point = Point { y: -1, x: 1 };
pub const DOWNLEFT: Point = Point { y: 1, x: -1 };
pub const DOWNRIGHT: Point = Point { y: 1, x: 1 };
pub const DIRECTIONS: [Point; 8] = [LEFT, RIGHT, UP, DOWN, UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT];

pub fn get_grid_point(grid: &Vec<Vec<char>>, p: &Point) -> char {
    match p {
        Point { y, .. } if *y < 0 || *y >= grid.len() as isize => char::default(),
        Point { x, .. } if *x < 0 || *x >= grid[0].len() as isize => char::default(),
        Point { y, x } => grid[*y as usize][*x as usize],
    }
}

pub fn move_point(p: &Point, dir: &Point, step: isize) -> Point {
    Point {
        y: p.y + (dir.y * step),
        x: p.x + (dir.x * step),
    }
}

pub fn lines_to_char_grid(lines: &Vec<String>) -> Vec<Vec<char>> {
    lines
        .iter()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<_>>()
}
