use core::hash::Hash;
use std::collections::HashMap;
use std::time::Instant;

pub fn assert_test(file: &str, func: fn(&str) -> usize, expected: usize) {
    assert_eq!(func(file), expected)
}

pub fn show_results(file: &str, part: usize, func: fn(&str) -> usize) {
    let res = func(file);
    println!("file: {file}, part: {part}, res: {res}");
}

pub fn measure(func: impl Fn()) {
    let start = Instant::now();
    func();
    let duration = start.elapsed();
    println!("time: {:?}", duration)
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

// fn to<T: FromStr + Default>(input: &str) -> T
// where <T as FromStr>::Err: Debug {
//     input
//         .parse::<T>()
//         .unwrap_or_default()
// }

pub fn to_int(input: &str) -> isize {
    input.parse::<isize>().unwrap_or_default()
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
