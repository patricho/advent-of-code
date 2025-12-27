use crate::util::{
    file::read_file_string,
    misc::{assert_test, char_to_usize, measure, show_results},
};

static FILE_TEST: &str = "data/2024/09_test.txt";
static FILE_TEST_2: &str = "data/2024/09_test_2.txt";
static FILE_INPUT: &str = "data/2024/09_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 1928);
    assert_test(FILE_TEST, part2, 2858);
    assert_test(FILE_TEST_2, part2, 169);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });

    assert_test(FILE_INPUT, part1, 6446899523367);
    assert_test(FILE_INPUT, part2, 6478232739671);
}

fn part1(filename: &str) -> usize {
    let (mut files, mut blanks) = parse_input(filename);
    let mut blank_idx = 0;
    let mut file_idx = files.len() - 1;

    loop {
        let mut curr_blank = blanks[blank_idx];
        let mut curr_file = files[file_idx];

        // Is the current file smaller than the current blank we want to put it in?
        if curr_file.count <= curr_blank.count {
            // Move all the file contents
            curr_file.start = curr_blank.start;
            curr_blank.count -= curr_file.count;
            curr_blank.start += curr_file.count;

            files[file_idx] = curr_file;
            blanks[blank_idx] = curr_blank;

            if file_idx > 0 {
                file_idx -= 1;
            }
        } else if curr_blank.count > 0 {
            // Chop of a bit
            let mut new_file_chunk = curr_file.clone();
            new_file_chunk.count = curr_blank.count;
            new_file_chunk.start = curr_blank.start;
            curr_file.count -= new_file_chunk.count;
            curr_blank.count = 0;
            curr_blank.start += new_file_chunk.count;
            files[file_idx] = curr_file;
            blanks[blank_idx] = curr_blank;
            files.push(new_file_chunk);
            blank_idx += 1;
        } else {
            blank_idx += 1;
        }

        // If the file or blank checking has reached each end, or passed each other, we're done
        if file_idx <= 0 || blank_idx >= blanks.len() || file_idx <= blank_idx {
            break;
        }
    }

    files.sort_by(|a, b| a.start.cmp(&b.start));

    let mut char_idx = 0;

    let total_checksum = files
        .iter()
        .map(|file| {
            (0..file.count)
                .map(|_| {
                    let char_checksum = char_idx * file.id;
                    char_idx += 1;
                    return char_checksum;
                })
                .sum::<usize>()
        })
        .sum::<usize>();

    total_checksum
}

fn part2(filename: &str) -> usize {
    let (mut files, mut blanks) = parse_input(filename);

    for file_idx in (0..files.len()).rev() {
        let mut curr_file = files[file_idx];

        for blank_idx in 0..blanks.len() {
            let mut curr_blank = blanks[blank_idx];

            // Does the current file fit into the current blank?
            if curr_file.count <= curr_blank.count && curr_blank.start < curr_file.start {
                // Move all the file contents
                curr_file.start = curr_blank.start;
                curr_blank.count -= curr_file.count;
                curr_blank.start += curr_file.count;

                // Write back to the vecs
                files[file_idx] = curr_file;
                blanks[blank_idx] = curr_blank;

                break;
            }
        }
    }

    files.sort_by(|a, b| a.start.cmp(&b.start));

    // debug_print("after", &files);
    // print_files(&files, &blanks);

    let total_checksum = files
        .iter()
        .map(|file| {
            let start_idx = (*file).start;
            (0..file.count)
                .map(|n| {
                    let char_idx = start_idx + n;
                    let char_checksum = char_idx * file.id;
                    return char_checksum;
                })
                .sum::<usize>()
        })
        .sum::<usize>();

    total_checksum
}

fn parse_input(filename: &str) -> (Vec<File>, Vec<Blank>) {
    let input = read_file_string(filename);

    let mut files: Vec<File> = vec![];
    let mut blanks: Vec<Blank> = vec![];

    let mut in_file = true;
    let mut file_id = 0;
    let mut idx = 0;

    input.chars().for_each(|c| {
        let n = char_to_usize(c);

        if in_file {
            files.push(File {
                id: file_id,
                count: n,
                start: idx,
            });
            file_id += 1;
        } else {
            blanks.push(Blank {
                count: n,
                start: idx,
            });
        }

        idx += n;
        in_file = !in_file;
    });

    (files, blanks)
}

/* fn print_files(files: &Vec<File>, blanks: &Vec<Blank>) {
    // Combine files and blanks into a single list with position info
    enum Block {
        FileBlock(usize, usize, usize), // (id, count, start)
        BlankBlock(usize, usize),       // (count, start)
    }

    let mut blocks: Vec<(usize, Block)> = vec![];

    for f in files {
        blocks.push((f.start, Block::FileBlock(f.id, f.count, f.start)));
    }
    for b in blanks {
        blocks.push((b.start, Block::BlankBlock(b.count, b.start)));
    }

    // Sort by start position
    blocks.sort_by_key(|(start, _)| *start);

    // Print interleaved
    for (idx, block) in blocks {
        match block {
            Block::FileBlock(id, count, start) => {
                let rep = repeat(id.to_string()).take(count).collect::<String>();
                print!("{rep}");
                // println!("{idx} | {start} | file: #{id} {count} {rep}");
            }
            Block::BlankBlock(count, start) => {
                let rep = repeat(".").take(count).collect::<String>();
                print!("{rep}");
                // println!("{idx} | {start} | blnk: {count} {rep}");
            }
        }
    }
    println!();
} */

#[derive(Clone, Copy, Debug)]
pub struct File {
    pub id: usize,
    pub start: usize,
    pub count: usize,
}

#[derive(Clone, Copy, Debug)]
pub struct Blank {
    pub start: usize,
    pub count: usize,
}
