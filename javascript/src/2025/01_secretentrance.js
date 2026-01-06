import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 3, () => part1('../inputs/2025/01-test.txt'))
        assert('part 2 test', 6, () => part2('../inputs/2025/01-test.txt'))

        assert('part 1', 1092, () => part1('../inputs/2025/01-input.txt'))
        assert('part 2', 6616, () => part2('../inputs/2025/01-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    let zeroes = 0
    let position = 50

    for (const step of getSteps(filename)) {
        position += step

        if (position % 100 === 0) {
            zeroes++
        }
    }

    return zeroes
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    return part2Loop(getSteps(filename))
}

/**
 * @param {number[]} steps
 * @returns {number}
 */
function part2Loop(steps) {
    let zeroes = 0
    let position = 50

    for (const step of steps) {
        const oldposition = position

        position += step

        const oldround = Math.floor(oldposition / 100)
        const posround = Math.floor(position / 100)

        let diff = Math.abs(posround - oldround)

        if (step < 0 && position % 100 === 0) {
            // Stepping backwards, landing on 0 - count that
            diff++
        } else if (step < 0 && oldposition % 100 === 0) {
            // Stepping backwards, starting at 0 - already counted that
            diff--
        }

        zeroes += diff
    }

    return zeroes
}

/**
 * @param {string} filename
 * @returns {number[]}
 */
function getSteps(filename) {
    const steps = []
    for (const line of readFileLines(filename)) {
        const step = parseInt(line.replace(/L/g, '-').replace(/R/g, ''))
        steps.push(step)
    }
    return steps
}
