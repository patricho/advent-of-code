use crate::util::file::read_file_lines;
use crate::util::misc::{assert_test, measure, show_results};

const FILE_TEST: &str = "../inputs/2025/06-test.txt";
const FILE_INPUT: &str = "../inputs/2025/06-input.txt";

struct Equation {
    numbers: Vec<isize>,
    sign: char,
}

pub fn main() {
    assert_test(FILE_TEST, part1, 4277556);
    assert_test(FILE_TEST, part2, 3263827);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> isize {
    let eqs = parse_input_part1(filename);
    sum_equations(&eqs)
}

fn part2(filename: &str) -> isize {
    let eqs = parse_input_part2(filename);
    sum_equations(&eqs)
}

fn parse_input_part1(filename: &str) -> Vec<Equation> {
    let lines = read_file_lines(filename);
    let lastline = &lines[lines.len() - 1];
    let columns: Vec<&str> = lastline.split_whitespace().collect();
    let mut eqs: Vec<Equation> = (0..columns.len())
        .map(|_| Equation {
            numbers: Vec::new(),
            sign: ' ',
        })
        .collect();

    for i in 0..lines.len() - 1 {
        let parts: Vec<&str> = lines[i].split_whitespace().collect();
        for (j, part) in parts.iter().enumerate() {
            let num: isize = part.parse().unwrap_or(0);
            eqs[j].numbers.push(num);
        }
    }

    let lastline_parts: Vec<&str> = lastline.split_whitespace().collect();
    for (j, part) in lastline_parts.iter().enumerate() {
        eqs[j].sign = part.chars().next().unwrap_or(' ');
    }

    eqs
}

fn parse_input_part2(filename: &str) -> Vec<Equation> {
    let lines = read_file_lines(filename);
    let columns = lines[0].len();
    let mut eqs: Vec<Equation> = Vec::new();
    let mut pending = Equation {
        numbers: Vec::new(),
        sign: ' ',
    };

    for x in (0..columns).rev() {
        let mut empty = true;
        let mut num_string = String::new();

        for line in &lines {
            let chars: Vec<char> = line.chars().collect();
            if x >= chars.len() {
                continue;
            }

            let n = chars[x];
            if n == ' ' {
                continue;
            }

            if n == '*' || n == '+' {
                pending.sign = n;
            } else {
                num_string.push(n);
            }

            empty = false;
        }

        if empty {
            eqs.push(pending);
            pending = Equation {
                numbers: Vec::new(),
                sign: ' ',
            };
        } else if !num_string.is_empty() {
            pending.numbers.push(num_string.parse().unwrap_or(0));
        }
    }

    if !pending.numbers.is_empty() {
        eqs.push(pending);
    }

    eqs
}

fn sum_equations(eqs: &[Equation]) -> isize {
    let mut result: isize = 0;

    for eq in eqs {
        let mut res = eq.numbers[0];
        match eq.sign {
            '*' => {
                for i in 1..eq.numbers.len() {
                    res *= eq.numbers[i];
                }
            }
            '+' => {
                for i in 1..eq.numbers.len() {
                    res += eq.numbers[i];
                }
            }
            _ => {}
        }
        result += res;
    }

    result
}
