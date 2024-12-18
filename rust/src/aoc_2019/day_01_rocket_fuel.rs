use crate::util::{measure, read_file_lines};

pub fn main() {
    let res = solve(1, "data/2019/01_test.txt");
    assert_eq!(res, 34241);

    measure(|| {
        solve(1, "data/2019/01_input.txt");
    });

    measure(|| {
        solve(2, "data/2019/01_input.txt");
    });

    // let temp = read_file_string("data/2019/01_test.txt").unwrap();
    // println!("Test: {temp}");
}

fn solve(part: i32, filename: &str) -> i32 {
    let lines = read_file_lines(filename);

    let sum = lines
        .iter()
        .map(|line| {
            let line_no = line.parse().unwrap();
            let mut n = calc(line_no);

            if part == 2 {
                let mut nn = n;
                while nn > 0 {
                    nn = calc(nn);
                    if nn > 0 {
                        n += nn
                    }
                }
            }

            return n;
        })
        .sum();

    println!("part: {part}, file: {filename}, sum: {sum}");

    return sum;
}

fn calc(n: i32) -> i32 {
    return ((n as f32 / 3.0).floor() - 2.0) as i32;
}
