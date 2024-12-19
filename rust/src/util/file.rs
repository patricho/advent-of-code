use std::fs::File;
use std::io::{BufRead, BufReader, Read};

pub fn read_file_lines(filename: &str) -> Vec<String> {
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    return reader.lines().map(|line| line.unwrap()).collect();
}

pub fn read_file_string(filename: &str) -> String {
    let file = File::open(filename).unwrap();
    let mut reader = BufReader::new(file);
    let mut sstr = String::new();
    reader.read_to_string(&mut sstr).unwrap();
    return sstr;
}
