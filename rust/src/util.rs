use core::hash::Hash;
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader, Read};
use std::time::Instant;

pub fn assert_test(file: &str, func: fn(&str) -> usize, expected: usize) {
    assert_eq!(func(file), expected)
}

pub fn show_results(file: &str, part: usize, func: fn(&str) -> usize) {
    let res = func(file);
    println!("file: {file}, part: {part}, res: {res}");
}

pub fn read_file_lines(filename: &str) -> Vec<String> {
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    return reader
        .lines()
        .map(|line| line.unwrap())
        .collect();
}

pub fn read_file_string(filename: &str) -> String {
    let file = File::open(filename).unwrap();
    let mut reader = BufReader::new(file);
    let mut sstr = String::new();
    reader
        .read_to_string(&mut sstr)
        .unwrap();
    return sstr;
}

pub fn measure(func: impl Fn()) {
    let start = Instant::now();
    func();
    let duration = start.elapsed();
    println!("time: {:?}", duration)
}

pub fn count_values<T: Eq + Hash>(vec: Vec<T>) -> HashMap<T, i32> {
    let mut hash = HashMap::new();

    for r in vec {
        // `entry` gets the given key's corresponding entry in the map
        // for in-place manipulation
        *hash.entry(r).or_default() += 1;
    }

    return hash;
}

// fn to<T: FromStr + Default>(input: &str) -> T
// where <T as FromStr>::Err: Debug {
//     input
//         .parse::<T>()
//         .unwrap_or_default()
// }

pub fn to_int(input: &str) -> i32 {
    input.parse::<i32>().unwrap_or_default()
}

pub fn get_hash_int<K: Eq>(hash: &HashMap<K, i32>, key: &K) -> i32
where K: Hash {
    return get_hash_value(&hash, key, 0);
}

pub fn get_hash_value<K: Eq + Hash, V: Copy>(hash: &HashMap<K, V>, key: &K, default: V) -> V {
    return hash
        .get(key)
        .unwrap_or(&default)
        .clone();
}

pub fn split_spaces_to_ints(line: &str) -> Vec<i32> {
    return line
        .split_whitespace()
        .map(|s| to_int(s))
        .collect();
}
