use crate::util::misc::{assert_test, char_to_usize, measure, show_results};
use std::collections::HashSet; // For tracking unique reachable 9s
use std::collections::VecDeque; // For BFS queue

static FILE_TEST: &str = "../data/2024/10_test.txt";
static FILE_INPUT: &str = "../data/2024/10_input.txt";

pub fn main() {
    assert_test(FILE_TEST, part1, 36); // Updated expected value
    assert_test(FILE_TEST, part2, 0); // Keep part2 assertion as 0 for now

    measure(|| {
        show_results(FILE_INPUT, 1, part1);
        show_results(FILE_INPUT, 2, part2);
    });
}

fn part1(filename: &str) -> usize {
    let input_str = std::fs::read_to_string(filename)
        .unwrap_or_else(|e| panic!("Failed to read {}: {}", filename, e));
    let grid: Vec<Vec<usize>> = input_str
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| char_to_usize(c)) // Assuming char_to_usize handles numeric chars
                .collect()
        })
        .collect();

    if grid.is_empty() {
        return 0;
    }

    let rows = grid.len();
    let cols = grid[0].len();
    let mut trailheads = Vec::new();

    let mut total_queue_pushes = 0; // Initialize counter

    for r in 0..rows {
        for c in 0..cols {
            if grid[r][c] == 0 {
                trailheads.push((r, c));
            }
        }
    }

    let mut total_score_sum = 0;

    for &(start_r, start_c) in &trailheads {
        let mut reachable_nines_for_this_trailhead = HashSet::new();
        
        // Queue for BFS: (row, col, current_path_height)
        let mut queue = VecDeque::new();
        // Visited set for the current BFS to avoid cycles and redundant work for this specific trailhead search
        // Stores (row, col, height_reached_at_this_cell_on_current_path_from_trailhead)
        let mut visited_in_current_search = HashSet::new();

        if grid[start_r][start_c] == 0 { // Should always be true by definition of trailhead
            queue.push_back((start_r, start_c, 0));
            total_queue_pushes += 1; // Increment for initial push
            visited_in_current_search.insert((start_r, start_c, 0));
        }

        while let Some((r, c, current_h)) = queue.pop_front() {
            // If current_h is 8, we are looking for a 9.
            // If grid[r][c] is 9 and current_h is 9 (meaning we just stepped onto it from an 8)
            // then it's a valid end point.
            if current_h == 9 { // We are looking for a cell with value 9, having followed a path of height 0..8
                if grid[r][c] == 9 {
                     reachable_nines_for_this_trailhead.insert((r,c));
                }
                // Path ends at 9, no further steps from here for this path.
                continue; 
            }

            // Explore neighbors: up, down, left, right
            let dr = [-1, 1, 0, 0];
            let dc = [0, 0, -1, 1];

            for i in 0..4 {
                let nr_signed = r as isize + dr[i];
                let nc_signed = c as isize + dc[i];

                // Check bounds
                if nr_signed >= 0 && nr_signed < rows as isize && nc_signed >= 0 && nc_signed < cols as isize {
                    let nr = nr_signed as usize;
                    let nc = nc_signed as usize;

                    // Check if the neighbor's height is current_h + 1
                    if grid[nr][nc] == current_h + 1 {
                        // If we are about to step onto a 9, the next height is 9.
                        // Otherwise, the next height is grid[nr][nc]
                        let next_h = current_h + 1;
                        
                        if !visited_in_current_search.contains(&(nr, nc, next_h)) {
                            visited_in_current_search.insert((nr, nc, next_h));
                            queue.push_back((nr, nc, next_h));
                            total_queue_pushes += 1; // Increment for BFS push
                        }
                    }
                }
            }
        }
        let trailhead_score = reachable_nines_for_this_trailhead.len();
        if filename == FILE_TEST {
            println!("Trailhead at ({}, {}), score: {}", start_r, start_c, trailhead_score);
        }
        total_score_sum += trailhead_score;
    }

    println!("Total queue pushes for {}: {}", filename, total_queue_pushes);
    total_score_sum
}

fn part2(filename: &str) -> usize {
    let _input = std::fs::read_to_string(filename)
        .unwrap_or_else(|e| panic!("Failed to read {}: {}", filename, e));
    // TODO: Implement Day 10 Part 2 Logic
    0
}
