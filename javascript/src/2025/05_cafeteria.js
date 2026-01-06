import { readFileString, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 3, () => part1('../inputs/2025/05-test.txt'))
        assert('part 2 test', 14, () => part2('../inputs/2025/05-test.txt'))

        assert('part 1', 598, () => part1('../inputs/2025/05-input.txt'))
        assert('part 2', 360341832208407, () => part2('../inputs/2025/05-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {{ ingredients: number[], ranges: number[][] }}
 */
function parseInput(filename) {
    const fileparts = readFileString(filename).split('\n\n')

    const ranges = fileparts[0].split('\n').map((line) => line.split('-').map((n) => parseInt(n)))
    const ingredients = fileparts[1]
        .split('\n')
        .filter((line) => line)
        .map((line) => parseInt(line))

    return { ingredients, ranges }
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const { ingredients, ranges } = parseInput(filename)

    return ingredients.filter((i) => ranges.some((r) => i >= r[0] && i <= r[1])).length
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const { ranges } = parseInput(filename)

    // Sort ranges by start value
    ranges.sort((a, b) => a[0] - b[0])

    let result = 0
    let maxCounted = -1

    for (const cr of ranges) {
        let start = cr[0]
        const end = cr[1]

        if (maxCounted >= end) {
            // Range is already completely covered
            continue
        }

        if (maxCounted >= start) {
            start = maxCounted + 1
        }

        const span = end - start + 1

        result += span
        maxCounted = end
    }

    return result
}
