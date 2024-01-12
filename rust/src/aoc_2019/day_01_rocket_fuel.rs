use crate::util;

pub fn main() {
    solve(1, "data/2019/01_test.txt");
    // assert_eq!(res, 34241);

    solve(1, "data/2019/01_input.txt");
    // assert_eq!(res, 3317970);

    solve(2, "data/2019/01_input.txt");
    // assert_eq!(res, 4974073);
}

fn solve(part: i32, filename: &str) -> i32 {
    let lines = util::read_file(filename).unwrap();

    let sum: i32 = lines
        .iter()
        .map(|line| {
            let lineno = line.parse::<i32>().unwrap();
            let mut n = calc(lineno);

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

    println!("part {} file {} sum: {}", part, filename, sum);

    return sum;
}

fn calc(n: i32) -> i32 {
    return ((n as f32 / 3_f32).floor() - 2_f32) as i32;
}
