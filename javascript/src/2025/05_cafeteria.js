import { readFileString, assert, prettyPrint, isPointInPolygon } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 3, () => part1('../inputs/2025/05-test.txt'))
        // assert('part 2 test', 14, () => part2('../inputs/2025/05-test.txt'))

        assert('part 1', 598, () => part1('../inputs/2025/05-input.txt'))
        // assert('part 2', 360341832208407, () => part2('../inputs/2025/05-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const fileparts = readFileString(filename).split('\n\n')

    const ranges = fileparts[0].split('\n').map((line) => line.split('-').map((n) => parseInt(n)))
    const ingredients = fileparts[1].split('\n').map((line) => parseInt(line))

    return ingredients.filter((i) => ranges.some((r) => i >= r[0] && i <= r[1])).length
}
