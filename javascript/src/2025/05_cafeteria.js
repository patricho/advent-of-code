import { readFileString } from '../util/file.js'

export default main

/**
 * @param {string} label
 * @param {number} want
 * @param {number} got
 */
function assert(label, want, got) {
    const match = want === got
    const icon = match ? '✓' : '✗'
    const white = '\x1b[97m'
    const lgray = '\x1b[37m'
    const gray = '\x1b[90m'
    const green = '\x1b[32m'
    const red = '\x1b[31m'
    const color = match ? green : red
    const reset = '\x1b[0m'
    console.log(`${color}${icon} ${lgray}${label} ${gray}want ${white}${want}${gray} got ${color}${got}${reset}`)
}

function main() {
    assert('part 1 test', 4, part1('../inputs/2025/05-test.txt'))

    assert('part 1', 598, part1('../inputs/2025/05-input.txt'))
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
