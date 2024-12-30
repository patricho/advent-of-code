use std::ops::{Add, AddAssign, Sub, SubAssign};

#[derive(Copy, Debug, PartialEq, Eq, Clone, Hash)]
pub struct Point {
    pub y: isize,
    pub x: isize,
}

impl Point {
    pub const fn new(y: isize, x: isize) -> Self {
        Point { y, x }
    }
}

impl Add for Point {
    type Output = Self;
    fn add(self, other: Self) -> Self::Output {
        Point::new(self.y + other.y, self.x + other.x)
    }
}

impl AddAssign for Point {
    fn add_assign(&mut self, other: Self) {
        self.y += other.y;
        self.x += other.x;
    }
}

impl Sub for Point {
    type Output = Self;
    fn sub(self, other: Self) -> Self::Output {
        Point::new(self.y - other.y, self.x - other.x)
    }
}

impl SubAssign for Point {
    fn sub_assign(&mut self, other: Self) {
        self.y -= other.y;
        self.x -= other.x;
    }
}

pub const UP: Point = Point::new(-1, 0);
pub const DOWN: Point = Point::new(1, 0);
pub const LEFT: Point = Point::new(0, -1);
pub const RIGHT: Point = Point::new(0, 1);
pub const UPLEFT: Point = Point::new(-1, -1);
pub const UPRIGHT: Point = Point::new(-1, 1);
pub const DOWNLEFT: Point = Point::new(1, -1);
pub const DOWNRIGHT: Point = Point::new(1, 1);
pub const DIRECTIONS: [Point; 8] = [LEFT, RIGHT, UP, DOWN, UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT];

pub fn grid_contains(grid: &Vec<Vec<char>>, p: &Point) -> bool {
    p.y >= 0 && (p.y as usize) < grid.len() && p.x >= 0 && (p.x as usize) < grid[0].len()
}

pub fn get_grid_point(grid: &Vec<Vec<char>>, p: &Point) -> char {
    match p {
        Point { y, .. } if *y < 0 || *y >= grid.len() as isize => char::default(),
        Point { x, .. } if *x < 0 || *x >= grid[0].len() as isize => char::default(),
        Point { y, x } => grid[*y as usize][*x as usize],
    }
}

pub fn set_grid_point(grid: &mut Vec<Vec<char>>, p: &Point, c: char) {
    grid[p.y as usize][p.x as usize] = c
}

pub fn move_point_steps(p: &Point, dir: &Point, step: isize) -> Point {
    Point { y: p.y + (dir.y * step), x: p.x + (dir.x * step) }
}

pub fn move_to_new_point(p: &Point, delta: &Point) -> Point {
    Point { y: p.y + delta.y, x: p.x + delta.x }
}

pub fn move_point(p: &mut Point, delta: &Point) {
    (*p).x += delta.x;
    (*p).y += delta.y;
}

pub fn lines_to_char_grid(lines: &Vec<String>) -> Vec<Vec<char>> {
    lines
        .iter()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<_>>()
}

pub fn print_grid(grid: &Vec<Vec<char>>) {
    grid.iter().for_each(|line| {
        println!("{}", line.iter().collect::<String>());
    })
}
