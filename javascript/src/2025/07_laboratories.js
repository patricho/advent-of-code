import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 21, () => part1('../inputs/2025/07-test.txt'))
        assert('part 2 test', 40, () => part2('../inputs/2025/07-test.txt'))

        assert('part 1', 1658, () => part1('../inputs/2025/07-input.txt'))
        assert('part 2', 53916299384254, () => part2('../inputs/2025/07-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const { splits } = traverse(filename)
    return splits
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const { combinations } = traverse(filename)
    return combinations
}

/**
 * @param {string} filename
 * @returns {{ splits: number, combinations: number }}
 */
function traverse(filename) {
    const lines = readFileLines(filename)
    const grid = lines.map((line) => line.split(''))

    let splits = 0
    const beams = new Array(grid[0].length).fill(0)

    // Find starting beam
    for (let i = 0; i < grid[0].length; i++) {
        if (grid[0][i] === 'S') {
            beams[i] = 1
            break
        }
    }

    for (let y = 0; y < grid.length; y++) {
        for (let x = 0; x < grid[y].length; x++) {
            // We're only interested in when a beam reaches a splitter
            if (grid[y][x] !== '^' || beams[x] === 0) {
                continue
            }

            splits++

            // Split left
            if (x - 1 >= 0) {
                beams[x - 1] += beams[x]
            }

            // Split right
            if (x + 1 < grid[y].length) {
                beams[x + 1] += beams[x]
            }

            beams[x] = 0
        }
    }

    const combinations = beams.reduce((sum, b) => sum + b, 0)

    return { splits, combinations }
}
