use crate::util::file::read_file_string;
use crate::util::misc::{assert_test, measure, show_results, to_int};
use std::cmp::Ordering;
use std::collections::HashMap;

static FILE_TEST: &str = "data/2024/05_test.txt";
static FILE_INPUT: &str = "data/2024/05_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 143);
    assert_test(FILE_TEST, part2, 123);

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> usize {
    let (rules, pages_lines) = parse_input(filename);

    // Filter out the lines of pages, where each individual page number is ok. A page is ok if
    // there isn't a matching rule for it, or the rule passes, which it does if the rightmost
    // page isn't found in the current line at all, or if it's found after the current page (ie
    // has a larger index). For each filtered line, find the middle index, convert that item to
    // a number, and sum them together

    pages_lines
        .iter()
        .filter(|pages| all_pages_are_ok(&rules, pages))
        .map(get_middle_value)
        .sum::<usize>()
}

fn part2(filename: &str) -> usize {
    let (rules, pages_lines) = parse_input(filename);

    let mut wrong_pages = pages_lines
        .into_iter()
        .filter(|pages| !all_pages_are_ok(&rules, pages))
        .collect::<Vec<_>>();

    wrong_pages
        .iter_mut()
        .map(|pages| {
            (*pages).sort_by(|next, current| {
                // The next value is checked against the current, and if there is a rule with the
                // next value as key, and its values contain the current that means the current one
                // should come _after_ the next, in other words they are in the wrong order, and
                // should be switched around

                match rules.get(next) {
                    Some(rule) if rule.contains(current) => Ordering::Less,
                    _ => Ordering::Equal,
                }
            });

            return pages;
        })
        .map(|pages| get_middle_value(&pages))
        .sum::<usize>()
}

/*
"The issue is that you're trying to return string slices (&str) that reference the input string,
which will be dropped when the function ends. To fix this, we need to convert the string slices
to owned Strings using to_string().
I've modified the return type to (Vec<(String, String)>, Vec<Vec<String>>) and added the
necessary conversions to owned strings throughout the function. This ensures that all data is
owned by the returned structures rather than borrowing from the local input variable that would
go out of scope."
*/

fn parse_input(filename: &str) -> (HashMap<String, Vec<String>>, Vec<Vec<String>>) {
    let input = read_file_string(filename);

    input
        .split_once("\n\n")
        .map(|(rules_str, pages_str)| {
            // Fold the rules to a hashmap, with the page number to check as key, and the value
            // as all the other pages (that need to come after) as a string vector
            let rules = rules_str
                .lines()
                .fold(HashMap::new(), |mut map, rule_ln| {
                    let (left, right) = rule_ln.split_once("|").unwrap();

                    map.entry(left.to_string())
                        .or_insert_with(Vec::new)
                        .push(right.to_string());
                    map
                });

            // Split the lines of page numbers to a vector of lines, where each line is split
            // into a vector of strings
            let pages_lines = pages_str
                .lines()
                .map(|rule_ln| {
                    rule_ln
                        .split(",")
                        .map(|s| s.to_string())
                        .collect::<Vec<_>>()
                })
                .collect::<Vec<_>>();

            (rules, pages_lines)
        })
        .unwrap()
}

fn rules_match_for_page(
    rules: &HashMap<String, Vec<String>>, pages: &Vec<String>, page: &String, idx_page: usize,
) -> bool {
    match rules.get(page) {
        Some(rules_for_page) => rules_for_page.iter().all(|rule_next| {
            // This rule is ok if the next page has a larger index than the current
            // one, or is not found at all
            match pages
                .iter()
                .position(|i| i == rule_next)
            {
                Some(idx_next) => idx_next > idx_page,
                None => true,
            }
        }),
        _ => true,
    }
}

fn get_middle_value(pages: &Vec<String>) -> usize {
    let mid_idx = pages.len() / 2;
    to_int(pages[mid_idx].as_str()) as usize
}

fn all_pages_are_ok(rules: &HashMap<String, Vec<String>>, pages: &Vec<String>) -> bool {
    pages
        .into_iter()
        .enumerate()
        .all(|(idx_page, page)| rules_match_for_page(&rules, pages, page, idx_page))
}
