use core::hash::Hash;
use std::collections::HashMap;
use std::fmt::{Debug, Display};
use std::time::Instant;

pub fn assert_test<T: Eq + Debug>(file: &str, func: fn(&str) -> T, want: T) {
    assert_eq!(func(file), want)
}

pub fn show_results<T: Display>(file: &str, part: usize, func: fn(&str) -> T) {
    let res = func(file);
    println!("\x1b[90mfile:\x1b[0m {file}\x1b[90m, part:\x1b[0m {part}\x1b[90m, res:\x1b[0m {res}");
}

pub fn measure(func: impl Fn()) {
    let start = Instant::now();
    func();
    let duration = start.elapsed();
    println!("\x1b[90mtime: {:?}\x1b[0m", duration)
}

pub fn count_values<T: Eq + Hash>(vec: Vec<T>) -> HashMap<T, isize> {
    let mut hash = HashMap::new();

    for r in vec {
        // `entry` gets the given key's corresponding entry in the map
        // for in-place manipulation
        *hash.entry(r).or_default() += 1;
    }

    return hash;
}

pub fn char_to_usize(c: char) -> usize {
    c.to_digit(10).unwrap() as usize
}

// fn to<T: FromStr + Default>(input: &str) -> T
// where <T as FromStr>::Err: Debug {
//     input
//         .parse::<T>()
//         .unwrap_or_default()
// }

pub fn to_int(input: &str) -> isize {
    input.parse::<isize>().unwrap_or_default()
}

pub fn debug_print<T: Debug>(label: &str, o: &T) {
    println!("{}: {:#?}", label, o);
}

pub fn get_hash_int<K: Eq>(hash: &HashMap<K, isize>, key: &K) -> isize
where
    K: Hash,
{
    return get_hash_value(&hash, key, 0);
}

pub fn get_hash_value<K: Eq + Hash, V: Copy>(hash: &HashMap<K, V>, key: &K, default: V) -> V {
    return hash.get(key).unwrap_or(&default).clone();
}

pub fn split_spaces_to_ints(line: &str) -> Vec<isize> {
    return line.split_whitespace().map(|s| to_int(s)).collect();
}

pub fn split_lines_to_vec(input: &str) -> Vec<String> {
    return input.lines().map(|line| line.to_string()).collect();
}

pub fn digit_count(n: usize) -> u32 {
    if n == 0 {
        return 1;
    }

    n.ilog10() + 1
}

pub fn signed_digit_count(n: isize) -> u32 {
    if n == 0 {
        return 1;
    }

    n.unsigned_abs().ilog10() + 1
}

pub fn split_digits(n: usize) -> (usize, usize) {
    let half_digits = (n.ilog10() + 1) / 2;
    let divisor = 10_usize.pow(half_digits);
    (n / divisor, n % divisor)
}
