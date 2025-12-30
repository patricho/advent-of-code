use regex::Regex;

use crate::util::{
    file::read_file_string,
    misc::{assert_test, measure, show_results, to_int},
};

const FILE_TEST: &str = "../inputs/2024/13-test.txt";
const FILE_INPUT: &str = "../inputs/2024/13-input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 480);
    assert_test(FILE_TEST, part2, 875318608908);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });

    assert_test(FILE_INPUT, part1, 28753);
    assert_test(FILE_INPUT, part2, 102718967795500);
}

fn part1(filename: &str) -> usize {
    let machines = parse_machines(filename);
    solve_machines(&machines, 0.0)
}

fn part2(filename: &str) -> usize {
    let machines = parse_machines(filename);
    solve_machines(&machines, 10000000000000.0)
}

struct Machine {
    button_ax: f64,
    button_ay: f64,
    button_bx: f64,
    button_by: f64,
    prize_x: f64,
    prize_y: f64,
}

fn parse_machines(filename: &str) -> Vec<Machine> {
    let input = read_file_string(filename);

    // Split on empty lines
    let machine_strings = input
        .split("\n\n")
        .map(|s| s.to_string())
        .collect::<Vec<_>>();

    let btn_regex = Regex::new(r"Button .: X\+([0-9]+), Y\+([0-9]+)").unwrap();
    let prize_regex = Regex::new(r"Prize: X=([0-9]+), Y=([0-9]+)").unwrap();

    // Parse machines (button A, button B, prize)
    let machines = machine_strings
        .iter()
        .map(|ms| {
            let machine_lines = ms.split("\n").map(|s| s.to_string()).collect::<Vec<_>>();

            let a = btn_regex.captures(&machine_lines[0]).unwrap();
            let btn_ax = to_int(&a.get(1).unwrap().as_str()) as f64;
            let btn_ay = to_int(&a.get(2).unwrap().as_str()) as f64;

            let b = btn_regex.captures(&machine_lines[1]).unwrap();
            let btn_bx = to_int(&b.get(1).unwrap().as_str()) as f64;
            let btn_by = to_int(&b.get(2).unwrap().as_str()) as f64;

            let p = prize_regex.captures(&machine_lines[2]).unwrap();
            let prize_x = to_int(&p.get(1).unwrap().as_str()) as f64;
            let prize_y = to_int(&p.get(2).unwrap().as_str()) as f64;

            Machine {
                button_ax: btn_ax,
                button_ay: btn_ay,
                button_bx: btn_bx,
                button_by: btn_by,
                prize_x,
                prize_y,
            }
        })
        .collect::<Vec<_>>();

    machines
}

fn solve_machines(machines: &Vec<Machine>, offset: f64) -> usize {
    // Formulas from: https://github.com/rjwut/advent/blob/main/src/solutions/2024/day-13.md

    machines
        .iter()
        .map(|m| {
            let prize_x = m.prize_x + offset;
            let prize_y = m.prize_y + offset;

            let presses_b = ((m.button_ax * prize_y) - (prize_x * m.button_ay))
                / ((m.button_ax * m.button_by) - (m.button_bx * m.button_ay));

            let presses_a = (prize_x - (presses_b * m.button_bx)) / m.button_ax;

            if presses_a.fract() > 0.0 || presses_b.fract() > 0.0 {
                return 0.0;
            }

            presses_a * 3.0 + presses_b
        })
        .sum::<f64>() as usize
}
